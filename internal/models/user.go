/*
 * @Author: zs
 * @Date: 2025-05-30 11:58:23
 * @LastEditors: zs
 * @LastEditTime: 2025-06-04 20:12:18
 * @FilePath: /barshop-server/internal/models/user.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package models

import (
	"gorm.io/gorm"
	"time"
	"strings"
	"github.com/zacus/barshop-server/internal/validator"
)

type User struct {
	gorm.Model
	Username    string    `gorm:"uniqueIndex;size:50" json:"username" validate:"required,username"`
	Password    string    `gorm:"size:100" json:"-" validate:"required,password_strength"`
	Name        string    `gorm:"size:50" json:"name" validate:"required,min=2,max=50"`
	Phone       string    `gorm:"size:20" json:"phone" validate:"required,phone_cn"`
	Email       string    `gorm:"size:100" json:"email" validate:"required,email"`
	Role        string    `gorm:"size:20" json:"role" validate:"required,oneof=customer barber admin"` 
	LastLoginAt time.Time `json:"last_login_at"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required,min=8"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,username"`
	Password string `json:"password" binding:"required,password_strength"`
	Name     string `json:"name" binding:"required,min=2,max=50"`
	Phone    string `json:"phone" binding:"required,phone_cn"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateUserRequest struct {
	Name  string `json:"name" binding:"required,min=2,max=50"`
	Phone string `json:"phone" binding:"required,phone_cn"`
	Email string `json:"email" binding:"required,email"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,password_strength"`
}

// Validate 验证用户数据
func (u *User) Validate() error {
	return validator.ValidateStruct(u)
}

// Clean 清理用户数据
func (u *User) Clean() {
	u.Username = validator.CleanString(u.Username)
	u.Name = validator.CleanString(u.Name)
	u.Phone = validator.CleanPhone(u.Phone)
	u.Email = validator.CleanEmail(u.Email)
	u.Role = strings.ToLower(strings.TrimSpace(u.Role))
}

// Clean 清理注册请求数据
func (r *RegisterRequest) Clean() {
	r.Username = validator.CleanString(r.Username)
	r.Name = validator.CleanString(r.Name)
	r.Phone = validator.CleanPhone(r.Phone)
	r.Email = validator.CleanEmail(r.Email)
}

// Clean 清理更新请求数据
func (r *UpdateUserRequest) Clean() {
	r.Name = validator.CleanString(r.Name)
	r.Phone = validator.CleanPhone(r.Phone)
	r.Email = validator.CleanEmail(r.Email)
}