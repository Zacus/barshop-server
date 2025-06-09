/*
 * @Author: zs
 * @Date: 2025-06-07 16:35:50
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 16:49:10
 * @FilePath: /barshop-server/internal/router/appointment_routes.go
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

type AppointmentRoutes struct {
	BaseRouterGroup
	appointmentService *services.AppointmentService
}

func NewAppointmentRoutes(appointmentService *services.AppointmentService, authMiddleware gin.HandlerFunc) *AppointmentRoutes {
	return &AppointmentRoutes{
		BaseRouterGroup:     NewBaseRouterGroup("/appointments", authMiddleware),
		appointmentService:  appointmentService,
	}
}

func (r *AppointmentRoutes) Register(group *gin.RouterGroup) {
	handler := handlers.NewAppointmentHandler(r.appointmentService)
	
	router := group.Group(r.Prefix, r.Middlewares...)
	{
		router.POST("/", handler.Create)
		router.GET("/", handler.List)
		router.PUT("/:id/status", handler.UpdateStatus)
	}
} 