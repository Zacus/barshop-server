/*
 * @Author: zs
 * @Date: 2025-06-07 16:22:28
 * @LastEditors: zs
 * @LastEditTime: 2025-06-07 16:36:27
 * @FilePath: /barshop-server/internal/router/router.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zacus/barshop-server/internal/config"
	"github.com/zacus/barshop-server/internal/middleware"
	"github.com/zacus/barshop-server/internal/services"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "github.com/zacus/barshop-server/docs"
)

type Router struct {
	engine   *gin.Engine
	config   *config.Config
	services *services.Container
	routes   []IRouterGroup
}

// NewRouter 创建一个新的路由管理器
func NewRouter(cfg *config.Config) *Router {
	// 设置gin模式
	gin.SetMode(cfg.Server.Mode)

	// 创建路由引擎
	engine := gin.Default()

	// 创建服务容器
	services := services.NewContainer()

	// 创建路由管理器
	router := &Router{
		engine:   engine,
		config:   cfg,
		services: services,
	}

	// 初始化路由组
	router.initRouteGroups()

	return router
}

// initRouteGroups 初始化所有路由组
func (r *Router) initRouteGroups() {
	// 创建中间件
	authMiddleware := middleware.AuthMiddleware(r.config.JWT.Secret)
	adminMiddleware := middleware.AdminOnly()

	// 初始化路由组
	r.routes = []IRouterGroup{
		// 认证路由
		NewAuthRoutes(r.services.User),
		
		// 用户路由
		NewUserRoutes(r.services.User, authMiddleware),
		
		// 服务路由（公开）
		NewPublicServiceRoutes(r.services.Service),
		
		// 服务路由（管理员）
		NewAdminServiceRoutes(r.services.Service, authMiddleware, adminMiddleware),
		
		// 预约路由
		NewAppointmentRoutes(r.services.Appointment, authMiddleware),
	}
}

// InitRoutes 初始化所有路由
func (r *Router) InitRoutes() {
	// 基础中间件
	r.engine.Use(gin.Logger())
	r.engine.Use(gin.Recovery())
	r.engine.Use(middleware.CORS())

	// Swagger文档路由
	r.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 健康检查路由
	r.engine.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "服务正常运行",
		})
	})

	// API路由组
	api := r.engine.Group("/api/v1")
	
	// 注册所有路由组
	for _, group := range r.routes {
		group.Register(api)
	}
}

// Run 启动HTTP服务器
func (r *Router) Run() error {
	return r.engine.Run(":" + r.config.Server.Port)
} 