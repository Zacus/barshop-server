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
		userService:     userService,
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