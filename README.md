<!--
 * @Author: zs
 * @Date: 2025-06-04 19:06:12
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 17:50:25
 * @FilePath: /barshop-server/README.md
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
-->
# 巴巴理发店管理系统

## 项目简介

巴巴理发店管理系统是一个现代化的理发店管理解决方案，帮助理发店提升运营效率和客户服务质量。

### 主要功能

- 预约管理
- 服务项目管理
- 会员管理
- 收银结算
- 数据统计分析

## 技术栈

- 后端：Go
- 数据库：PostgreSQL
- 日志：Zap
- 配置：YAML

## 系统要求

- Go 1.20 或以上
- PostgreSQL 14 或以上
- Python 3.8 或以上（用于构建和部署脚本）

## 快速开始

### 1. 安装依赖

```bash
# 安装 Python 依赖
cd scripts
pip install -r requirements.txt

# 安装 Go 依赖
go mod download
```

### 2. 配置

复制示例配置文件并修改：

```bash
cp config.yaml.example config.yaml
```

主要配置项：
```yaml
server:
  port: "8080"
  mode: "debug"

database:
  driver: "postgres"  # 可选: postgres, mysql
  host: "localhost"
  port: 5432
  user: "postgres"
  password: "your_password"
  dbname: "barshop"
  sslmode: "disable"  # PostgreSQL专用
  charset: "utf8mb4"  # MySQL专用
  auto_migrate: false  # 默认不进行自动迁移
  options:
    max_idle_conns: 10
    max_open_conns: 100
    conn_max_lifetime: 60  # 分钟
    conn_max_idle_time: 10 # 分钟
    debug: true

redis:
  host: "localhost"
  port: 6379
  db: 0
  password: ""
  pool_size: 100

jwt:
  secret: "your-secret-key"
  expire: "24h"

log:
  level: "debug"
  is_dev: true
```

### 3. 构建和运行

#### 本地开发环境

```bash
# 构建
./scripts/build.sh local

# 运行
./bin/barshop-server
```

#### 生产环境部署

1. 创建发布包：
```bash
./scripts/build.sh release
```

2. 解压并运行：
```bash
tar xzf barshop-server-*.tar.gz
cd barshop-server-*

# 修改配置
vim config.yaml

# 运行
./bin/barshop-server
```

#### Docker 环境

```bash
# 构建镜像
./scripts/build.sh docker

# 运行容器
docker run -d \
  -p 8080:8080 \
  -v $(pwd)/config.yaml:/app/config.yaml \
  -v $(pwd)/logs:/app/logs \
  barshop-server:latest
```

## 项目结构

```
.
├── cmd/api/          # 应用入口
├── configs/          # 配置文件
├── docs/            # API文档
├── internal/        # 内部代码
│   ├── cache/      # 缓存实现
│   ├── config/     # 配置管理
│   ├── database/   # 数据库管理
│   │   ├── manager/   # 数据库管理器
│   │   └── transaction/ # 事务管理
│   ├── handlers/   # HTTP处理器
│   ├── middleware/ # 中间件
│   ├── models/     # 数据模型
│   ├── repository/ # 数据访问层
│   │   ├── postgres/     # PostgreSQL实现
│   │   └── repointerface/ # 仓储接口
│   ├── router/     # 路由管理
│   ├── services/   # 业务逻辑
│   └── utils/      # 工具函数
└── scripts/        # 部署脚本
```

## 脚本说明

### build.sh

构建脚本，支持以下命令：
```bash
./scripts/build.sh clean    # 清理构建文件
./scripts/build.sh local    # 本地构建
./scripts/build.sh docker   # 构建 Docker 镜像
./scripts/build.sh release  # 构建发布包
```

### deploy.sh

部署脚本，支持以下命令：
```bash
export DEPLOY_HOST="your-server"  # 设置部署服务器
export DEPLOY_USER="deploy"       # 设置部署用户
export APP_PORT="8080"           # 设置应用端口

./scripts/deploy.sh setup    # 设置部署环境
./scripts/deploy.sh deploy   # 部署应用
./scripts/deploy.sh docker   # 部署 Docker 容器
```

## 发布包结构

```
barshop-server-v1.0.0/
├── bin/
│   └── barshop-server    # 可执行文件
├── config.yaml           # 配置文件
├── scripts/              # 管理脚本
└── logs/                # 日志目录
```

## 常见问题

### 1. 配置文件找不到

确保：
- 配置文件 `config.yaml` 与可执行文件在同一目录
- 或者在程序运行目录下有 `config.yaml`

### 2. 日志目录权限

确保：
- `logs` 目录存在
- 运行程序的用户有写入权限

### 3. 数据库连接

检查：
- 数据库配置是否正确
- 数据库服务是否运行
- 网络连接是否正常

## 许可证

本项目采用 MIT 许可证