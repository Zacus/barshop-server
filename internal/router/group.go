package router

import (
	"github.com/gin-gonic/gin"
)

// IRouterGroup 定义路由组的接口
type IRouterGroup interface {
	Register(group *gin.RouterGroup)
}

// BaseRouterGroup 提供基础路由组实现
type BaseRouterGroup struct {
	Prefix      string
	Middlewares []gin.HandlerFunc
}

// NewBaseRouterGroup 创建基础路由组
func NewBaseRouterGroup(prefix string, middlewares ...gin.HandlerFunc) BaseRouterGroup {
	return BaseRouterGroup{
		Prefix:      prefix,
		Middlewares: middlewares,
	}
} 