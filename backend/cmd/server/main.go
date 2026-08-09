package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"codex-overview-backend/internal/app"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("resolve working directory failed: %v", err)
	}
	appDir := resolveAppDir(workingDir)
	workspaceRoot := filepath.Join(appDir, "workspace")
	staticDir := filepath.Join(appDir, "web", "dist")

	addrFlag := flag.String("addr", "127.0.0.1:8787", "http listen address")
	workspaceFlag := flag.String("workspace-root", workspaceRoot, "workspace root containing auth directories")
	staticFlag := flag.String("static-dir", staticDir, "frontend dist directory")
	openBrowserFlag := flag.Bool("open-browser", true, "open browser after server starts")
	cacheTTLFlag := flag.Duration("cache-ttl", 20*time.Second, "snapshot cache ttl")
	flag.Parse()

	resolvedWorkspace := strings.TrimSpace(*workspaceFlag)
	if resolvedWorkspace == "" {
		resolvedWorkspace = workspaceRoot
	}
	if err = os.MkdirAll(resolvedWorkspace, 0o700); err != nil {
		log.Fatalf("create workspace root failed: %v", err)
	}

	server := app.NewServer(app.ServerConfig{
		AppRoot:       appDir,
		WorkspaceRoot: resolvedWorkspace,
		StaticDir:     strings.TrimSpace(*staticFlag),
		CacheTTL:      *cacheTTLFlag,
		AppName:       "Codex 额度总览",
		DefaultPrice:  7.5,
	})

	address := strings.TrimSpace(*addrFlag)
	if address == "" {
		address = "127.0.0.1:8787"
	}
	if !isLoopbackListenAddress(address) {
		log.Fatalf("refusing non-loopback listen address %q: this credential dashboard has no remote authentication layer; use an SSH tunnel", address)
	}

	httpServer := &http.Server{
		Addr:              address,
		Handler:           server.Handler(),
		ReadHeaderTimeout: 10 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}

	url := fmt.Sprintf("http://%s", address)
	log.Printf("%s 启动中，逻辑 CPU=%d，工作目录=%s", server.Config().AppName, runtime.NumCPU(), app.DisplayPath(server.Config().WorkspaceRoot))
	log.Printf("打开地址：%s", url)

	if *openBrowserFlag {
		go func() {
			time.Sleep(900 * time.Millisecond)
			_ = app.OpenBrowser(url)
		}()
	}

	if err = httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server exited: %v", err)
	}
}

func isLoopbackListenAddress(address string) bool {
	host, _, err := net.SplitHostPort(strings.TrimSpace(address))
	if err != nil {
		return false
	}
	host = strings.TrimSpace(strings.Trim(host, "[]"))
	if strings.EqualFold(host, "localhost") {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

func resolveAppDir(workingDir string) string {
	candidates := []string{workingDir, filepath.Dir(workingDir)}
	if executable, err := os.Executable(); err == nil {
		executableDir := filepath.Dir(executable)
		candidates = append(candidates, filepath.Dir(executableDir), executableDir)
	}

	seen := make(map[string]struct{}, len(candidates))
	for _, candidate := range candidates {
		candidate = filepath.Clean(candidate)
		if _, exists := seen[candidate]; exists {
			continue
		}
		seen[candidate] = struct{}{}
		if info, err := os.Stat(filepath.Join(candidate, "web")); err == nil && info.IsDir() {
			return candidate
		}
	}
	return filepath.Dir(workingDir)
}
