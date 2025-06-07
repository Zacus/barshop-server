/*
 * @Author: zs
 * @Date: 2025-06-07 16:34:43
 * @LastEditors: zs
 * @LastEditTime: 2025-06-07 16:38:54
 * @FilePath: /barshop-server/internal/router/service_routes.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zacus/barshop-server/internal/handlers"
	"github.com/zacus/barshop-server/internal/services"
)

type ServiceRoutes struct {
	BaseRouterGroup
	serviceService *services.ServiceService
}

// NewPublicServiceRoutes 创建公开的服务路由
func NewPublicServiceRoutes(serviceService *services.ServiceService) *ServiceRoutes {
	return &ServiceRoutes{
		BaseRouterGroup: NewBaseRouterGroup("/services"),
		serviceService:  serviceService,
	}
}

// NewAdminServiceRoutes 创建管理员的服务路由
func NewAdminServiceRoutes(serviceService *services.ServiceService, authMiddleware, adminMiddleware gin.HandlerFunc) *ServiceRoutes {
	return &ServiceRoutes{
		BaseRouterGroup: NewBaseRouterGroup("/admin/services", authMiddleware, adminMiddleware),
		serviceService:  serviceService,
	}
}

func (r *ServiceRoutes) Register(group *gin.RouterGroup) {
	handler := handlers.NewServiceHandler(r.serviceService)
	
	router := group.Group(r.Prefix, r.Middlewares...)
	
	// 根据前缀判断是否为管理员路由
	if r.Prefix == "/admin/services" {
		router.POST("/", handler.Create)
		router.PUT("/:id", handler.Update)
		router.DELETE("/:id", handler.Delete)
		router.PUT("/:id/toggle", handler.ToggleStatus)
	} else {
		router.GET("/", handler.List)
	}
}