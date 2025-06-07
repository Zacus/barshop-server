/*
 * @Author: zs
 * @Date: 2025-06-04 19:31:16
 * @LastEditors: zs
 * @LastEditTime: 2025-06-07 16:22:51
 * @FilePath: /barshop-server/cmd/api/main.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package main

import (
	"fmt"
	"github.com/zacus/barshop-server/internal/cache"
	"github.com/zacus/barshop-server/internal/config"
	"github.com/zacus/barshop-server/internal/database"
	"github.com/zacus/barshop-server/internal/logger"
	"github.com/zacus/barshop-server/internal/router"
	"os"
	"os/signal"
	"syscall"
)

// @title           理发店管理系统 API
// @version         1.0
// @description     这是一个理发店管理系统的API文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description 请在此输入 Bearer {token} 格式的JWT令牌

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// 初始化日志
	if err := logger.InitLogger(cfg.Log.Level, cfg.Log.IsDev); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// 初始化数据库
	if err := database.InitDB(cfg); err != nil {
		logger.Fatal("Failed to initialize database", "error", err)
	}

	// 初始化Redis缓存
	if err := cache.InitRedis(cfg); err != nil {
		logger.Fatal("Failed to initialize Redis", "error", err)
	}
	defer cache.Close()

	// 创建路由管理器
	r := router.NewRouter(cfg)
	
	// 初始化路由
	r.InitRoutes()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Server starting", "port", cfg.Server.Port)
		if err := r.Run(); err != nil {
			logger.Fatal("Failed to start server", "error", err)
		}
	}()

	<-quit
	logger.Info("Shutting down server...")
} 