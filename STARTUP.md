# 项目启动指南

> 仅保留本地/非容器化启动方式，数据库使用 MySQL，文件存储使用阿里云 OSS。

## 启动前准备
- 操作系统：Windows / macOS / Linux
- 基础环境：
  - Node.js ≥ 18（Vite 5 需此版本；建议安装 pnpm 或 npm，任选其一）
  - Go ≥ 1.18
  - MySQL 5.7+（已运行，确保账号/密码/数据库名可用）
  - Redis（若在配置中 `use-redis: true`，请先启动 Redis 服务）
- 前端依赖：在首次启动前需安装 web 依赖
  - 进入 `web` 目录后运行 `npm install`（或 `pnpm install`）
- 后端依赖：在 `server` 目录运行 `go mod tidy`（首次或依赖变更时）
- 配置文件：
  - 编辑 `server/config.yaml`，确认以下项：
    - `system.db-type: mysql`
    - `mysql` 段填好 `path`/`port`/`db-name`/`username`/`password`
    - `redis` 段：如启用 Redis，填写 `addr`/`db`/`password`
    - `system.oss-type: aliyun-oss`，并在 `aliyun-oss` 段填好密钥与 bucket 信息

## 启动命令

### 后端（server）
```bash
cd server
# 如首次或依赖有改动
go mod tidy
# 生成代码（可选，如需）
go generate
# 构建
go build -o server main.go   # Windows 可改为 server.exe
# 运行
./server                      # Windows: server.exe
```

### 前端（web）
```bash
cd web
# 安装依赖（首次/依赖更新时）
npm install
# 启动开发服务器
npm run serve
```

## 常见检查项
- 数据库连接失败：确认 `server/config.yaml` 中 MySQL 地址、端口、用户名、密码、库名正确且可访问。
- Redis 相关报错：若不用 Redis，将 `system.use-redis` 设为 `false`，并重启后端。
- OSS 访问失败：检查 `aliyun-oss` 配置（endpoint、access-key-id/secret、bucket-name、bucket-url）。
- 端口占用：后端默认 8888，前端默认 `VITE_CLI_PORT`（见 `web/.env.*`）。如冲突请修改配置并重启。

