package models

import (
	"gorm.io/gorm"
	"time"
)

type Appointment struct {
	gorm.Model
	CustomerID uint      `json:"customer_id"`
	Customer   User      `gorm:"foreignKey:CustomerID" json:"customer"`
	BarberID   uint      `json:"barber_id"`
	Barber     User      `gorm:"foreignKey:BarberID" json:"barber"`
	ServiceID  uint      `json:"service_id"`
	Service    Service   `json:"service"`
	StartTime  time.Time `json:"start_time"`
	EndTime    time.Time `json:"end_time"`
	Status     string    `gorm:"size:20;default:'pending'" json:"status"` // pending, confirmed, completed, cancelled
	Note       string    `gorm:"type:text" json:"note"`
}

type AppointmentRequest struct {
	BarberID  uint      `json:"barber_id" binding:"required"`
	ServiceID uint      `json:"service_id" binding:"required"`
	StartTime time.Time `json:"start_time" binding:"required"`
	Note      string    `json:"note"`
} 