# 安全策略 / Security policy

## 支持范围 / Supported versions

安全修复只面向最新发布版本与当前默认分支。报告前请先在不覆盖真实数据的副本中确认问题仍可复现。

Security fixes target the latest release and the current default branch. Reproduce issues on a copy that does not overwrite real data.

## 私密报告 / Private reporting

请使用 GitHub 仓库的 **Security → Report a vulnerability**。不要在公开 Issue、截图或日志中上传认证 JSON、Token、邮箱、账号 ID、扫描结果、本机路径或导入目录。只提供从零生成的合成样本。

Use GitHub **Security → Report a vulnerability**. Never place auth JSON, tokens, email addresses, account IDs, scan results, local paths, or imported directories in public issues, screenshots, or logs. Use synthetic samples only.

## 安全边界 / Security boundary

- 后端强制绑定 `localhost`、`127.0.0.0/8` 或 `::1`，没有远程认证层，不应直接暴露到局域网或公网。
- 浏览器 API 拒绝非本机跨域来源；凭证导入有请求和单文件大小上限，导入文件权限按私有数据处理。
- HTTP Host 必须是 localhost 或回环 IP 字面量，以阻断 DNS rebinding；CSV 导出只读取已完成扫描结果并中和公式前缀。
- 服务会读取认证 JSON 并向官方配额端点发出请求；导入副本、持久化结果和浏览器本地存储都可能包含敏感账号信息。
- 本项目不加密这些运行数据，也不能保护已控制主机、浏览器进程、备份或用户账号的攻击者。磁盘、备份和操作系统账号必须由部署者保护。

The backend is loopback-only and has no remote authentication layer. It requires a literal loopback Host, rejects non-local browser origins, and limits credential uploads. CSV export reads completed results only and neutralizes spreadsheet formula prefixes. Imported files, persisted results, and browser storage can contain sensitive account data and are not encrypted by this project. Host, browser, filesystem, and backup protection remain the operator's responsibility.
