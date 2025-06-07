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
		userService:     userService,
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