# 贡献指南 / Contributing

感谢你愿意改进 Codex Quota Overview。为了让问题更容易复现、改动更容易审核，请遵循下面的流程。

## 开始之前

- 功能建议或使用问题请先搜索现有 Issues。
- 较大的功能改动请先创建 Feature request，说明使用场景和预期行为。
- 不要提交真实认证文件、扫描结果、日志、运行目录或其他个人数据。

## 本地开发

需要 Go 1.25+、Node.js 18+ 和 npm。

```powershell
cd backend
go test ./...
go vet ./...

cd ..\web
npm ci
npm run lint
npm run build
```

启动开发环境的方法见 [README](./README.md)。

## 提交改动

1. 从最新的 `main` 创建分支。
2. 保持改动聚焦，不混入无关重构或格式化。
3. 为后端行为变化补充或更新测试。
4. 修改界面时同时检查桌面宽屏、窄屏、亮色和深色主题。
5. 提交 Pull Request，并完整填写模板中的验证结果与界面截图。

## 设计与行为约束

- 扫描只能由用户手动触发，刷新页面不得自动开始扫描。
- 大型账户结果必须保持服务端分页。
- 导入文件必须保留安全的相对目录结构。
- 结果快照必须可按 `resultId` 持久化和恢复。
- 面向普通用户的主要界面文案使用自然、清晰的简体中文。

---

Contributions are welcome. Please keep changes focused, run the backend and frontend checks above, avoid committing real account data, and open an issue before starting a large behavioral change.
