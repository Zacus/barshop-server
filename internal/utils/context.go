/*
 * @Author: zs
 * @Date: 2025-06-04 20:15:23
 * @LastEditors: zs
 * @LastEditTime: 2025-06-04 20:15:23
 * @FilePath: /barshop-server/internal/utils/context.go
 * @Description: 上下文工具函数
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package utils

import (
	"github.com/gin-gonic/gin"
)

// GetUserIDFromContext 从上下文中获取用户ID
func GetUserIDFromContext(c *gin.Context) uint {
	// 从上下文中获取用户ID
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}

	// 将接口类型转换为 uint
	if id, ok := userID.(uint); ok {
		return id
	}

	// 如果转换失败，返回0
	return 0
}

// GetUserRoleFromContext 从上下文中获取用户角色
func GetUserRoleFromContext(c *gin.Context) string {
	// 从上下文中获取用户角色
	role, exists := c.Get("user_role")
	if !exists {
		return ""
	}

	// 将接口类型转换为 string
	if roleStr, ok := role.(string); ok {
		return roleStr
	}

	// 如果转换失败，返回空字符串
	return ""
} 