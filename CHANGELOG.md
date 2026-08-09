# 更新记录 / Changelog

本项目遵循 [Semantic Versioning](https://semver.org/)。

## [Unreleased]

- 强制后端只监听回环地址，移除通配 CORS，并拒绝非本机跨域来源。
- 为导入和 JSON API 增加请求总量、单文件大小和安全响应头门禁。
- 补齐标准 Web 图标与 Manifest，修复主题首帧闪烁和键盘操作，移除无功能占位按钮。
- 新增双语安全策略与本地/服务器运维说明。
- 默认扫描根目录改为应用内 `workspace/`，不再隐式枚举仓库或安装目录的同级文件夹；其他根目录必须显式传入。
- UI 元数据、导入响应与启动诊断会遮蔽当前用户主目录前缀，避免截图和日志暴露本机用户名与绝对用户路径。

## [1.0.2] - 2026-08-07

- 产品与仓库名称统一为 Codex Quota Overview / Codex 额度总览。
- 强化普通用户下载入口，补充贡献指南、支持说明、行为准则与结构化模板。
- 修复免安装包在缺少源码目录时无法可靠启动的问题。
- CI 新增前端 lint，并修复分页状态与无效旧属性引起的 lint 问题。
- 更新前端锁定依赖。

## [1.0.1] - 2026-03-08

- 修复 GitHub 社交分享封面中的中文字体渲染。
- 补充仓库社交分享封面资源。

## [1.0.0] - 2026-03-08

- 首个公开版本。
- 支持批量导入认证 JSON、并发扫描、服务端分页与结果持久化。
- 提供额度总览、账户列表、窗口详情、CSV 导出与明暗主题。
- 提供 Windows 一键安装、启动和停止脚本。

[Unreleased]: https://github.com/ferretgeek/codex-quota-overview/compare/v1.0.2...HEAD
[1.0.2]: https://github.com/ferretgeek/codex-quota-overview/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/ferretgeek/codex-quota-overview/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/ferretgeek/codex-quota-overview/releases/tag/v1.0.0
