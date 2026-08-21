[![Codex quota overview](./docs/images/social-preview.png)](https://github.com/ferretgeek/codex-quota-overview/releases/latest)

# Codex quota overview · many accounts at once

[中文](./README.md) · English

[![CI](https://github.com/ferretgeek/codex-quota-overview/actions/workflows/ci.yml/badge.svg)](https://github.com/ferretgeek/codex-quota-overview/actions/workflows/ci.yml)
[![Latest release](https://img.shields.io/github/v/release/ferretgeek/codex-quota-overview?display_name=tag)](https://github.com/ferretgeek/codex-quota-overview/releases/latest)
[![License](https://img.shields.io/github/license/ferretgeek/codex-quota-overview)](LICENSE)

> Import dozens of Codex auth files at once and find out, in one pass, how much quota each account has left.

## Why this exists

If you have one Codex account, the [quota widget](https://github.com/ferretgeek/codex-quota-widget) is all you need.

If you have dozens of `auth.json` files scattered across folders, with no idea which still have quota and which are spent, signing into each one is not a plan.

This tool does that job. Point it at folders, it recursively finds every auth file, queries them concurrently, and gives you a table: total quota, remaining, consumed, and per-account window details. Export CSV if you want a record.

**Everything runs on your own machine, bound to `127.0.0.1`, with no login layer.**

## Interface

| Light | Dark |
|---|---|
| ![Light theme overview](./docs/images/demo-01-light.png) | ![Dark theme overview](./docs/images/demo-02-dark.png) |

> The quota and exchange-rate figures in these screenshots are redacted display samples. Real results come from whatever the program computes at run time.

## Getting started

**If you just want to use it**, download the [prebuilt Windows package from Releases](https://github.com/ferretgeek/codex-quota-overview/releases/latest), unzip, and follow the included instructions. No development environment needed.

1. Run `一键安装环境.bat` (install environment)
2. Run `一键启动服务.bat` (start service)
3. Open `http://127.0.0.1:8787`, click "select folder," then "scan now"

When you're done, run `一键停止服务.bat` (stop service). Step-by-step instructions are in `操作说明.txt`.

> Browsers usually only allow picking one folder at a time. Pick several in a row — they queue up for import so you can confirm the list before scanning.
>
> The page **never** scans on its own. Nothing happens until you click.

<details>
<summary>Development mode</summary>

<br />

Requires Windows 10/11, Go 1.25+, Node.js 18+, and npm.

Backend:

```powershell
cd backend
go run .\cmd\server -open-browser=false
```

Frontend:

```powershell
cd web
npm install
npm run dev
```

</details>

## What it does

- **Recursive import** — pick folders in the browser, several in a row, and every `JSON` inside is found recursively.
- **Concurrent queries** — recommended concurrency is derived from your CPU thread count, so you don't have to guess.
- **Handles large pools** — results are paged server-side, so hundreds of accounts don't freeze the page.
- **Results persist** — refreshing the page doesn't trigger a rescan; previous results are still there.
- **Export and cleanup** — CSV export, clear statistics, and clear the import directory are three separate actions.

## Worth noting technically

**The scan root is deliberately narrow.** By default only the application's own git-ignored `workspace/` directory is scanned — it never enumerates siblings of the repository or install directory. Reading any other root requires passing `-workspace-root <PATH>` explicitly. That restriction is intentional: nobody should be able to "just scan" their entire home directory by accident.

**Loopback only, and no public login layer.** To run it on a Linux server, still bind `127.0.0.1` and reach it through an SSH tunnel. **Do not change it to `0.0.0.0`.** This isn't a missing feature — a service that reads local credential files shouldn't have a public entry point.

**The export endpoint never triggers a scan.** `GET /api/export.csv` returns completed results only and rejects a `force` parameter, so nobody can re-scan an account pool with a single GET.

**Go backend, React frontend.** Paging, scan jobs, and persistence live in `backend/internal/app`. Jobs are asynchronous; the frontend polls `GET /api/job?id=...` for progress.

<details>
<summary>API surface</summary>

<br />

```text
GET  /api/health
GET  /api/meta
POST /api/import-folder
POST /api/scan-job
POST /api/refresh-job
GET  /api/job?id=...
GET  /api/accounts?resultId=...
POST /api/clear-imported-files
POST /api/clear-stats
GET  /api/export.csv          completed results only; never scans; rejects force
```

</details>

## Layout

```text
backend/                  Go backend
├─ cmd/server/            entry point
└─ internal/app/          routing, scanning, paging, persistence
web/                      React + Vite frontend
docs/images/              demo screenshots
一键安装环境.bat            install environment
一键启动服务.bat            start service
一键停止服务.bat            stop service
操作说明.txt                usage walkthrough
AI接手指南.md               handover notes written for coding agents
```

## What it doesn't do

- It doesn't sign in for you, create accounts, or modify any auth file (read-only).
- It doesn't bypass or relax any usage limit.
- It exposes no public entry point.
- It doesn't scan automatically — you click.

## Privacy

This repository contains no real credentials, account pools, scan results, or runtime logs. **Never commit real auth files, import directories, or result directories.**

## Development checks

```powershell
cd backend; go test ./...; go vet ./...
cd ..\web; npm install; npm run build
```

## More documentation

[Operations](./docs/OPERATIONS.md) · [Changelog](./CHANGELOG.md) · [Contributing](./CONTRIBUTING.md) · [Security policy](./SECURITY.md) · [Support](./SUPPORT.md) · [Code of conduct](./CODE_OF_CONDUCT.md)

Report security issues privately per the [security policy](./SECURITY.md). Never paste real accounts, tokens, or paths into a public issue.

## License and disclaimer

MIT License — see [LICENSE](LICENSE).

This is an independent community project with no affiliation with, authorization from, or endorsement by OpenAI, and it does not bypass any usage limit. All trademarks belong to their respective owners.
