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
### 4. Nginx 反向代理配置（Apple Silicon 芯片）

前端通过 `http://localhost/uc/...` 请求后端时，默认会访问 `80` 端口。你可以通过配置 Nginx 反向代理，将请求转发给运行在 `8888` 端口的 `ucenter-api` 网关服务。

#### 🍏 Nginx 常用管理命令 (Apple Silicon):

- **启动 Nginx**: `brew services start nginx`
- **停止 Nginx**: `brew services stop nginx`
- **重启 Nginx**: `brew services restart nginx`
- **重载配置 (不重启)**: `nginx -s reload`

#### 📂 配置文件路径 (Apple Silicon):
`/opt/homebrew/etc/nginx/nginx.conf`

#### ⚙️ 反向代理配置示例:
打开配置文件，在 `http` 块的 `server`（通常监听 `80` 端口）中添加以下 `location` 规则，将前端的 `/uc/` 转发至后端：

```nginx
server {
    listen       80;
    server_name  localhost;

    # 反向代理用户中心 API 网关 (ucenter-api)
    location /uc/ {
        proxy_pass http://127.0.0.1:8888;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    }
}
```

---

## 📝 提交与更新日志 (Change Log)

> [!NOTE]
> 本项目使用符合 [Conventional Commits](https://www.conventionalcommits.org/zh-hans/v1.0.0/) 规范 the 提交历史，并在下方记录详细的更新日志。

| 日期 | 提交类型 | 提交内容描述 | 提交人 |
| :--- | :--- | :--- | :--- |
| 2026-06-04 | `feat` | 实现发送验证码与手机号注册业务流程，引入公共工具库并完成数据库持久化与联调 | hddm |
| 2026-06-04 | `docs` | 整理 Apple Silicon 芯片的 Nginx 常用管理命令与反向代理配置到 README 中 | hddm |
| 2026-06-04 | `fix` | 修复 `ucenter-api` 注册逻辑的 `copier.Copy` 错误、Captcha 空指针异常，解决 `ucenter` 中 `member.go` 的编译类型对比报错，并延长 zRPC 客户端超时与 GORM 慢 SQL 判定阈值 | hddm |
| 2026-06-03 | `feat` | 实现发送短信验证码功能并接入 Redis 缓存进行验证码有效期控制 | hddm |
| 2026-06-03 | `docs` | 创建并完善根目录 `README.md` 项目说明文件，增加提交日志记录 | hddm |
| 2026-06-03 | `fix` | 修复 `ucenter-api` 中连接 RPC 服务时的 Etcd Key 拼写错误 (`ucenter.rpc` -> `uclient.rpc`) | hddm |
| 2026-06-03 | `feat` | 初始化 Go 工作区（go.work）及项目基础骨架，实现用户中心服务骨架与手机注册接口，配置 Etcd 注册发现 | hddm |
