/*
 * @Author: zs
 * @Date: 2025-06-07 16:34:18
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 16:49:35
 * @FilePath: /barshop-server/internal/router/auth_routes.go
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

type AuthRoutes struct {
	BaseRouterGroup
	userService *services.UserService
}

func NewAuthRoutes(userService *services.UserService) *AuthRoutes {
	return &AuthRoutes{
		BaseRouterGroup: NewBaseRouterGroup("/auth"),
		userService: userService,
	}
}

func (r *AuthRoutes) Register(group *gin.RouterGroup) {
	handler := handlers.NewUserHandler(r.userService)
	
	router := group.Group(r.Prefix)
	{
		router.POST("/register", handler.Register)
		router.POST("/login", handler.Login)
	}
} 