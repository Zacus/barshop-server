<!--
 * @Author: zs
 * @Date: 2025-06-04 19:06:12
 * @LastEditors: zs
 * @LastEditTime: 2025-06-04 20:00:36
 * @FilePath: /barshop-server/README.md
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
-->
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
git clone https://github.com/zacus/barshop-server.git
cd barshop-server
```