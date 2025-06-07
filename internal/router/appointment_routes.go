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