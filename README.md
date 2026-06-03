# MS-Coin 微服务项目

MS-Coin 是基于 [go-zero](https://github.com/zeromicro/go-zero) 框架开发的高并发、易扩展的微服务系统。项目使用 Go 多模块工作区（Go Workspaces）进行代码组织。

---

## 📂 项目结构

```text
ms-coin/
├── ucenter-api/       # 用户中心 API 网关服务 (HTTP)
├── ucenter/           # 用户中心 RPC 微服务 (gRPC)
├── grpc-common/       # gRPC 协议文件（Proto）与生成的客户端、服务端代码
├── mscoin-common/     # 项目通用公共库
└── go.work            # Go 工作区配置文件
```

---

## 🛠️ 技术栈与依赖

- **开发语言**: Go
- **微服务框架**: [go-zero](https://github.com/zeromicro/go-zero)
- **服务注册与发现**: Etcd (127.0.0.1:2379)
- **通信协议**: HTTP / gRPC

---

## ⚙️ 快速启动

### 1. 启动 Etcd
服务发现依赖于 Etcd。在本地启动 Etcd 容器：
```bash
docker run -d --name etcd -p 2379:2379 -p 2380:2380 appscode/etcd:3.5.0
```

### 2. 启动用户中心 RPC 服务
```bash
cd ucenter
go run main.go -f etc/conf.yaml
```

### 3. 启动用户中心 API 网关服务
```bash
cd ucenter-api
go run main.go -f etc/conf.yaml
```

---

## 📝 提交与更新日志 (Change Log)

> [!NOTE]
> 本项目使用符合 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/v1.0.0/) 规范的提交历史，并在下方记录详细的更新日志。

| 日期 | 提交类型 | 提交内容描述 | 提交人 |
| :--- | :--- | :--- | :--- |
| 2026-06-03 | `feat` | 初始化 Go 工作区（go.work）及项目基础骨架，实现用户中心服务骨架与手机注册接口，配置 Etcd 注册发现 | Antigravity / User |
| 2026-06-03 | `fix` | 修复 `ucenter-api` 中连接 RPC 服务时的 Etcd Key 拼写错误 (`ucenter.rpc` -> `uclient.rpc`) | Antigravity |
| 2026-06-03 | `docs` | 创建并完善根目录 `README.md` 项目说明文件，增加提交日志记录 | Antigravity |
| 2026-06-03 | `feat` | 实现发送短信验证码功能并接入 Redis 缓存进行验证码有效期控制 | Antigravity / User |
