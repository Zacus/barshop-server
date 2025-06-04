package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zacus/barshop-server/internal/database"
	"github.com/zacus/barshop-server/internal/models"
	"net/http"
)

// AdminOnly 仅允许管理员访问的中间件
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		
		var user models.User
		if err := database.DB.First(&user, userID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse(http.StatusUnauthorized, "未找到用户"))
			c.Abort()
			return
		}

		if user.Role != "admin" {
			c.JSON(http.StatusForbidden, models.ErrorResponse(http.StatusForbidden, "需要管理员权限"))
			c.Abort()
			return
		}

		c.Next()
	}
} 