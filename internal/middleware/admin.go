/*
 * @Author: zs
 * @Date: 2025-06-08 00:30:27
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 01:52:48
 * @FilePath: /barshop-server/internal/middleware/admin.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zacus/barshop-server/internal/models"
	"github.com/zacus/barshop-server/internal/repository/repointerface"
	"net/http"
)

// AdminOnly 仅允许管理员访问的中间件
func AdminOnly(userRepo repointerface.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetUint("user_id")
		

		user,err := userRepo.FindByID(context.Background(), userID)
		if err != nil {
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