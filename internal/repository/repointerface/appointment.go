/*
 * @Author: zs
 * @Date: 2025-06-08 16:50:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 00:59:09
 * @FilePath: /barshop-server/internal/repository/interface/appointment.go
 * @Description: 预约仓储接口定义
 */
package repointerface

import (
	"context"
	"time"

	"github.com/zacus/barshop-server/internal/models"
)

// AppointmentRepository 预约仓储接口
type AppointmentRepository interface {
	// 基础CRUD
	Create(ctx context.Context, appointment *models.Appointment) error
	FindByID(ctx context.Context, id uint) (*models.Appointment, error)
	Update(ctx context.Context, appointment *models.Appointment) error
	Delete(ctx context.Context, id uint) error

	// 业务查询
	ListByCustomer(ctx context.Context, customerID uint) ([]*models.Appointment, error)
	ListByBarber(ctx context.Context, barberID uint) ([]*models.Appointment, error)
	ListByDateRange(ctx context.Context, barberID uint, start, end time.Time) ([]*models.Appointment, error)
	ListByStatus(ctx context.Context, status string) ([]*models.Appointment, error)
	CheckTimeConflict(ctx context.Context, barberID uint, start, end time.Time) (bool, error)
}