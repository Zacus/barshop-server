/*
 * @Author: zs
 * @Date: 2025-06-04 16:52:48
 * @LastEditors: zs
 * @LastEditTime: 2025-06-08 00:46:47
 * @FilePath: /barshop-server/internal/models/service.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package models

import (
	"time"
	"gorm.io/gorm"
)

type Service struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt   time.Time `gorm:"index" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Name        string  `gorm:"size:100" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Duration    int     `gorm:"not null;default:30" json:"duration"` // 服务时长（分钟）
	Price       float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	IsActive    bool    `gorm:"default:true;index" json:"is_active"`
	CategoryID  uint    `gorm:"index" json:"category_id"`
}

type ServiceRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Duration    int     `json:"duration" binding:"required,min=1"`
	Price       float64 `json:"price" binding:"required,min=0"`
} 