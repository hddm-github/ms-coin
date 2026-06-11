---
name: git-commit-guidelines
description: "强制要求并规范在本项目（MS-Coin）中使用中文编写 Conventional Commits 格式的 Git 提交消息，并在每次提交前更新 README.md 中的 Change Log。"
---

# MS-Coin Git Commit & Changelog Guidelines

本技能（Skill）规定了在开发本项目时，AI 助手和开发者应遵循的 Git 提交及 Changelog 更新规范。

## When to Use (适用场景)
- **只在准备执行 `git commit` 提交代码时**，才需要触发和遵循本技能更新 `README.md` 中的 Change Log。
- **平时的代码修改、调试或普通答疑开发过程中**，请勿修改根目录下 `README.md` 中的 Change Log。

## Guidelines (规范详情)

### 1. Git 提交信息格式 (Commit Message Format)
提交信息必须全部采用**中文**书写，并且必须采用列表（List）格式详细说明修改项。

格式模板：
```text
<type>(<scope>): <首行简要总结说明>

- <详细修改项 1>
- <详细修改项 2>
- <详细修改项 3>
```

#### 常用分类类型 (Conventional Types)：
- `feat`: 新增特性/功能。**如果本次提交同时包含功能开发 and bug 修复，必须优先归类为 `feat`**。
- `fix`: 修复 Bug（包含编译错误、空指针异常、网络/超时优化等）。
- `docs`: 文档改动（如配置说明、Nginx 代理、README 描述等）。
- `refactor`: 重构代码。
- `chore`: 杂务/构建/依赖/配置调整。

### 2. Changelog 更新流程
1. **优先更新**: 在提交代码（`git commit`）前，必须先更新根目录下 [README.md](README.md) 中的 `Change Log`（提交与更新日志）表格。
2. **统一提交人**: 更新日志中的**提交人**一栏统一写为 `hddm`。
3. **内容同步**: 日志中的描述与类型必须与 Git Commit 消息保持高度一致。
4. **严格限制**: **仅在需要进行 Git 提交时才修改 Changelog**。若仅仅是修改代码、调试或日常对话，绝对不要修改或更新 `README.md` 中的日志。

---

## Examples (应用示例)

### 示例 1: 新增功能提交
```text
feat(ucenter): 实现发送验证码与手机号注册功能

- 实现发送短信验证码并接入 Redis 缓存控制有效期
- 实现手机号注册完整业务流程（包含查重、密码 bcrypt 哈希加密、数据落库）
- 引入 mscoin-common 公共库，支持基础数据库连接、事务处理及各种通用工具包（JWT、真实IP等）
- 建立 ucenter 数据库操作层（DAO/Repo）和 Member 领域模型，完成数据持久化
```

### 示例 2: Bug 修复与配置调优提交
```text
fix(ucenter): 修复注册接口 copier 报错、人机校验空指针，并优化超时判定

- 修复 ucenter-api 注册接口中 copier.Copy 的目标指针为 nil 导致的运行时错误
- 增加对前端传参 Captcha 为 nil 时的防空指针安全处理，自动初始化空结构体
- 修复 ucenter 模块中 member.go 对 sql.NullInt64 结构体与整型常量进行非法对比的编译报错
- 在 ucenter-api 配置中将 RPC 客户端调用超时时间延长至 10 秒
- 调整 GORM 慢 SQL 判定阈值至 2 秒，调整 Redis 与 SQL 慢调用阈值至 3 秒，避免开发日志刷屏
```
