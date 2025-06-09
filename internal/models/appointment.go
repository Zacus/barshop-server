/*
 * @Author: zs
 * @Date: 2025-06-04 16:53:07
 * @LastEditors: zs
 * @LastEditTime: 2025-06-08 15:01:54
 * @FilePath: /barshop-server/internal/models/appointment.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package models

import (
	"gorm.io/gorm"
	"time"
)

type Appointment struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt   time.Time `gorm:"index" json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	CustomerID  uint      `gorm:"index;not null" json:"customer_id"`
	Customer    User      `gorm:"foreignKey:CustomerID" json:"customer"`
	BarberID    uint      `gorm:"uniqueIndex:idx_barber_time;not null" json:"barber_id"`
	Barber      User      `gorm:"foreignKey:BarberID" json:"barber"`
	ServiceID   uint      `gorm:"index;not null" json:"service_id"`
	Service     Service   `json:"service"`
	StartTime   time.Time `gorm:"uniqueIndex:idx_barber_time;not null" json:"start_time"`
	EndTime     time.Time `gorm:"index;not null" json:"end_time"`
	Status      string    `gorm:"size:20;index;not null;default:'pending'" json:"status"` // pending, confirmed, completed, cancelled
	Note        string    `gorm:"type:text" json:"note"`
	PaymentStatus string  `gorm:"size:20;index;not null;default:'unpaid'" json:"payment_status"` // unpaid, paid, refunded
	Price        float64  `gorm:"type:decimal(10,2);not null" json:"price"`
}

type AppointmentRequest struct {
	BarberID  uint      `json:"barber_id" binding:"required"`
	ServiceID uint      `json:"service_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	Note      string    `json:"note"`
} 