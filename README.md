[![Codex 额度总览](./docs/images/social-preview.png)](https://github.com/ferretgeek/codex-quota-overview/releases/latest)

# Codex 额度总览 · 多账号批量

中文 · [English](./README_EN.md)

[![CI](https://github.com/ferretgeek/codex-quota-overview/actions/workflows/ci.yml/badge.svg)](https://github.com/ferretgeek/codex-quota-overview/actions/workflows/ci.yml)
[![最新版本](https://img.shields.io/github/v/release/ferretgeek/codex-quota-overview?display_name=tag&label=%E7%89%88%E6%9C%AC)](https://github.com/ferretgeek/codex-quota-overview/releases/latest)
[![开源许可](https://img.shields.io/github/license/ferretgeek/codex-quota-overview?label=%E8%AE%B8%E5%8F%AF)](LICENSE)

> 一次导入几十上百个 Codex 登录文件，一次问清每个账号还剩多少。

## 为什么会需要它

如果你手上只有一个 Codex 账号，[额度悬浮窗](https://github.com/ferretgeek/codex-quota-widget)就够了。

但如果你手上有几十个 `auth.json`——散落在不同文件夹里，不知道哪些还有额度、哪些已经空了——一个个登录去看是不可能的事。

这个工具做的就是这件事：把文件夹丢进去，它递归找出所有认证文件，并发查询，然后给你一张表：总额度、当前剩余、已经用掉的部分、每个账号的窗口详情。想留档就导出 CSV。

**全程在你自己电脑上跑，只监听 `127.0.0.1`，没有登录层也不需要联网以外的任何东西。**

## 界面

| 浅色 | 深色 |
|---|---|
| ![浅色主题总览](./docs/images/demo-01-light.png) | ![深色主题总览](./docs/images/demo-02-dark.png) |

> 截图里的额度和汇率数字都是为展示做的脱敏示例，实际结果以程序当次计算为准。

## 三步开始

**普通用户请直接下载现成的**：[Releases 里的 Windows 免安装包](https://github.com/ferretgeek/codex-quota-overview/releases/latest)，解压后按包内说明启动，不用装开发环境。

1. 双击 `一键安装环境.bat`
2. 双击 `一键启动服务.bat`
3. 浏览器打开 `http://127.0.0.1:8787`，点"手动选择文件夹"，再点"立即扫描"

用完双击 `一键停止服务.bat`。完整图文步骤见 `操作说明.txt`。

> 浏览器一次通常只能选一个文件夹。有多个就连续选几次，它们会先进待导入队列，确认无误再一起扫。
>
> 页面**不会**自动扫描，只有你点按钮才开始。

<details>
<summary>开发模式（要改代码的话）</summary>

<br />

需要 Windows 10/11、Go 1.25+、Node.js 18+ 和 npm。

后端：

```powershell
cd backend
go run .\cmd\server -open-browser=false
```

前端：

```powershell
cd web
npm install
npm run dev
```

</details>

## 它能做什么

- **递归导入** — 在浏览器里手动选文件夹，可以连续选多个再一次性导入，自动递归找出里面所有 `JSON`。
- **并发查询** — 按 CPU 线程数自动算出推荐并发，不用你猜。
- **撑得住号池** — 结果由服务端分页返回，几百个账号也不会把页面卡死。
- **结果落盘** — 刷新页面不会自动重扫，扫过的结果还在。
- **导出与清理** — 导出 CSV、清空统计、清空导入目录，各自独立。

## 技术上值得一提的地方

**扫描根目录是收紧过的。** 默认只扫应用内已被忽略的 `workspace/` 目录，不会去枚举仓库或安装目录的同级文件夹。要读其他目录必须显式传 `-workspace-root <PATH>`——这是刻意的，避免"随手一扫把整个用户目录翻了"。

**只监听回环，而且没有公网登录层。** 需要在 Linux 服务器上跑，就仍然绑 `127.0.0.1`，通过 SSH 隧道访问；**不要改成 `0.0.0.0`**。这不是没做，是故意不做——一个能读本地凭据文件的服务不该有对外入口。

**导出接口不触发扫描。** `GET /api/export.csv` 只导出已完成的结果，也不接受 `force` 参数，避免有人靠一个 GET 请求把号池刷一遍。

**Go 后端 + React 前端。** 分页、扫描任务、落盘都在 `backend/internal/app` 里；任务是异步的，前端轮询 `GET /api/job?id=...` 拿进度。

<details>
<summary>接口一览</summary>

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
GET  /api/export.csv          只导出已完成结果，不触发扫描，不接受 force
```

</details>

## 目录结构

```text
backend/                  Go 后端
├─ cmd/server/            服务入口
└─ internal/app/          路由、扫描、分页、落盘
web/                      React + Vite 前端
docs/images/              演示图片
一键安装环境.bat
一键启动服务.bat
一键停止服务.bat
操作说明.txt
AI接手指南.md              给编码助手看的接手说明
```

## 它不做什么

- 不帮你登录、不创建账号、不修改任何认证文件（只读）。
- 不绕过、不放宽任何额度限制。
- 不提供公网访问入口。
- 不自动扫描——必须你点。

## 隐私

本仓库不包含任何真实凭证、号池、扫描结果或运行日志。**请不要把真实账号文件、导入目录或结果目录提交进 Git。**

## 开发校验

```powershell
cd backend; go test ./...; go vet ./...
cd ..\web; npm install; npm run build
```

## 更多文档

[运维指南](./docs/OPERATIONS.md) · [版本变更](./CHANGELOG.md) · [参与开发](./CONTRIBUTING.md) · [安全策略](./SECURITY.md) · [获取支持](./SUPPORT.md) · [行为准则](./CODE_OF_CONDUCT.md)

安全问题请按[安全策略](./SECURITY.md)私密报告，不要在公开 Issue 里贴真实账号、Token 或路径。

## 许可与声明

MIT License，见 [LICENSE](LICENSE)。

这是独立的社区项目，与 OpenAI 没有隶属、授权或背书关系，也不绕过任何额度限制。相关商标归其权利人所有。
