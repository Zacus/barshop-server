/*
 * @Author: zs
 * @Date: 2025-06-07 16:34:29
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 16:50:13
 * @FilePath: /barshop-server/internal/router/user_routes.go
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

type UserRoutes struct {
	BaseRouterGroup
	userService *services.UserService
}

func NewUserRoutes(userService *services.UserService, authMiddleware gin.HandlerFunc) *UserRoutes {
	return &UserRoutes{
		BaseRouterGroup: NewBaseRouterGroup("/users", authMiddleware),
		userService: userService,
	}
}

func (r *UserRoutes) Register(group *gin.RouterGroup) {
	handler := handlers.NewUserHandler(r.userService)
	
	router := group.Group(r.Prefix, r.Middlewares...)
	{
		router.GET("/profile", handler.GetProfile)
		router.PUT("/profile", handler.UpdateProfile)
		router.PUT("/password", handler.ChangePassword)
		router.GET("/barbers", handler.ListBarbers)
	}
} 