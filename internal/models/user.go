/*
 * @Author: zs
 * @Date: 2025-05-30 11:58:23
 * @LastEditors: zs
 * @LastEditTime: 2025-06-04 17:06:47
 * @FilePath: /barshop-server/internal/models/user.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username    string    `gorm:"uniqueIndex;size:50" json:"username"`
	Password    string    `gorm:"size:100" json:"-"`
	Name        string    `gorm:"size:50" json:"name"`
	Phone       string    `gorm:"size:20" json:"phone"`
	Email       string    `gorm:"size:100" json:"email"`
	Role        string    `gorm:"size:20" json:"role"`
	LastLoginAt time.Time `json:"last_login_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
} 