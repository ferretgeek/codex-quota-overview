# Codex Quota Overview v1.0.2

这是一次面向普通用户使用体验与长期维护的更新。

## 本版改进

- 产品与仓库名称统一为 **Codex Quota Overview / Codex 额度总览**。
- 修复账户列表和窗口详情在筛选条件变化时的分页状态处理。
- 清理已经停用的自动刷新属性，保持“只由用户手动触发扫描”的产品约束。
- 更新前端锁定依赖，并把 lint 纳入持续集成。
- 增加清晰的下载入口、版本记录、贡献指南、支持说明与 Issue 模板。

## Windows 免安装版

下载 `codex-quota-overview-windows-portable-v1.0.2.zip` 后：

1. 解压到一个独立文件夹。
2. 双击 `一键启动服务.bat`。
3. 浏览器会自动打开 `http://127.0.0.1:8787`。
4. 使用结束后双击 `一键停止服务.bat`。

压缩包已经包含前端构建产物与 Windows 后端程序，不需要安装 Go、Node.js 或 npm。

---

This maintenance release unifies the product name, fixes pagination state transitions, refreshes locked frontend dependencies, enables lint in CI, and improves project documentation. The portable Windows archive includes the prebuilt frontend and backend; no development toolchain is required.
