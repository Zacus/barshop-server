# Barshop Server

理发店管理系统服务端

## 功能特性

- 用户管理（顾客、理发师）
- 服务管理
- 预约管理
- JWT认证
- 日志系统
- Swagger API文档

## 技术栈

- Go 1.24
- Gin Web框架
- GORM
- PostgreSQL
- Zap日志
- Swagger文档

## 项目结构

```
barshop-server/
├── cmd/                # 应用程序入口
│   └── api/           # API服务器
├── internal/          # 内部包
│   ├── config/       # 配置
│   ├── database/     # 数据库
│   ├── handlers/     # HTTP处理器
│   ├── logger/       # 日志
│   ├── middleware/   # 中间件
│   ├── models/       # 数据模型
│   └── services/     # 业务逻辑
├── docs/             # 文档
└── logs/             # 日志文件
```

## 快速开始

1. 克隆项目

```bash
git clone https://github.com/yourusername/barshop-server.git
cd barshop-server
```

2. 安装依赖

```bash
go mod download
```

3. 配置数据库

创建`config.yaml`文件：

```yaml
server:
  port: "8080"
  mode: "debug"

database:
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "your_password"
  dbname: "barshop"
  sslmode: "disable"

redis:
  host: "localhost"
  port: 6379
  db: 0

jwt:
  secret: "your_jwt_secret"
  expire: "24h"

log:
  level: "debug"
  is_dev: true
```

4. 运行服务

```bash
go run cmd/api/main.go
```

## API文档

启动服务后访问：`http://localhost:8080/swagger/index.html`

## 开发

### 生成Swagger文档

```bash
swag init -g cmd/api/main.go
```

### 运行测试

```bash
go test ./...
```

## 许可证

MIT License 