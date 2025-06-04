package models

import (
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name        string  `gorm:"size:100" json:"name"`
	Description string  `gorm:"type:text" json:"description"`
	Duration    int     `json:"duration"` // 服务时长（分钟）
	Price       float64 `json:"price"`
	IsActive    bool    `gorm:"default:true" json:"is_active"`
}

type ServiceRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Duration    int     `json:"duration" binding:"required,min=1"`
	Price       float64 `json:"price" binding:"required,min=0"`
} 