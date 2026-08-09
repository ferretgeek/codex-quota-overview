# 运维指南 / Operations

## 架构与数据 / Architecture and data

React 静态前端与 Go API 由同一个回环 HTTP 服务提供。后端读取明确选择的认证目录，导入副本写入仓库运行目录的 `imports/`，扫描快照写入 `results/`，界面设置和最近结果摘要保存在浏览器本地存储。三处都可能包含敏感账号信息。

The React bundle and Go API share one loopback HTTP origin. Imported copies live in `imports/`, persisted scan snapshots in `results/`, and UI preferences/recent summaries in browser storage. Treat all three as sensitive.

默认扫描根目录是应用内已忽略的 `workspace/`，避免开发克隆或便携目录把无关的同级项目误当成凭据目录。只有经过明确检查后才通过 `-workspace-root` 指向其他目录。浏览器元数据与启动诊断会把当前用户主目录前缀替换为 `~` 或 `%USERPROFILE%`；实际文件操作仍使用真实路径。

The default scan root is the ignored app-local `workspace/` directory. This prevents a development clone or portable copy from treating unrelated sibling projects as credential folders. Use `-workspace-root` only for an explicitly reviewed directory. Browser metadata and startup diagnostics replace the current user-home prefix with `~` or `%USERPROFILE%`; internal filesystem access still uses the real path.

## 本机运行 / Local operation

普通用户使用 Release 压缩包和仓库根目录的安装、启动、停止脚本。源码模式：

```powershell
cd web
npm ci
npm run build
cd ..\backend
go test ./...
go run .\cmd\server -addr 127.0.0.1:8787
```

默认只允许回环监听；这也是安全门禁，不要通过改源码或代理把 API 裸露到其他主机。

Packaged users should use the root install/start/stop scripts. Source users build the frontend, test the backend, and keep the server on a loopback address.

## Linux 服务器 + SSH 隧道 / Linux server with an SSH tunnel

```bash
npm ci --prefix web
npm run build --prefix web
cd backend
go test ./...
go run ./cmd/server \
  -addr 127.0.0.1:8787 \
  -open-browser=false \
  -workspace-root /srv/codex-quota/auth
```

客户端建立隧道后仍访问本机 URL：

```bash
ssh -N -L 8787:127.0.0.1:8787 operator@example.com
```

服务器账号应只读所需认证目录，并独占可写的 `imports/` 与 `results/`。不要用公网反向代理代替 SSH 隧道；产品没有多用户授权、会话或 CSRF Token 层。

Keep the service loopback-only and use SSH port forwarding. The service account needs read access only to the selected auth root and private write access to `imports/` and `results/`. Do not substitute a public reverse proxy: this product has no multi-user authorization or session layer.

## 升级、备份与恢复 / Upgrade, backup, and restore

1. 停止服务，确认没有扫描任务运行。
2. 加密备份需要保留的 `imports/`、`results/`；浏览器设置可在浏览器站点数据中单独清除，通常无需迁移。
3. 在新源码或新解压目录运行 Go 测试、`go vet ./...`、前端 lint 与构建，再启动新版本。
4. 需要恢复时，先安装相同或更新版本，再把备份复制回相同的运行根目录并限制为当前服务账号可读。
5. 回滚使用旧程序副本和对应数据快照；不要让两个版本同时写同一 `imports/` 或 `results/`。

Stop the service before copying encrypted backups of `imports/` and `results/`. Validate a new version before switching. Restore into the same private runtime root, and never let two versions write the same data directories concurrently.

## 健康检查与排错 / Health checks and troubleshooting

- `GET http://127.0.0.1:8787/api/health` 应返回成功；响应应包含 `X-Content-Type-Options: nosniff`，且不应包含通配 `Access-Control-Allow-Origin`。
- 启动拒绝地址：只使用 `127.0.0.1:8787`、`localhost:8787` 或 `[::1]:8787`。
- 页面提示前端缺失：在 `web/` 运行 `npm ci` 与 `npm run build`，再重启后端。
- 导入失败：只选择有效 JSON；单文件上限 4 MiB，单次请求上限 128 MiB，前端会按文件数分批。
- 扫描失败：确认服务账号能读取目标目录、认证内容仍有效且官方端点可达；公开求助前清除所有账号字段和路径。

The health endpoint should succeed and return hardened headers without wildcard CORS. Address errors indicate a non-loopback bind. Rebuild a missing frontend, respect the 4 MiB per-file and 128 MiB per-request import limits, and redact all diagnostics before sharing them.

## 卸载 / Uninstall

先运行停止脚本或结束服务进程，再删除程序目录。若不再需要数据，单独删除 `imports/`、`results/` 和浏览器对 `127.0.0.1:8787` 的站点数据；这些删除不可从项目恢复，先确认加密备份有效。服务器还应移除对应的 systemd 单元或启动配置和 SSH 转发配置。

Stop the process first. Remove the program, then separately remove `imports/`, `results/`, and the browser site data only after verifying any required encrypted backup. Remove any server unit and SSH-forwarding configuration as well.
