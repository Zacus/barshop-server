/*
 * @Author: zs
 * @Date: 2025-06-08 16:50:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 01:38:41
 * @FilePath: /barshop-server/internal/repository/postgres/appointment.go
 * @Description: 预约仓储PostgreSQL实现
 */
package postgres

import (
	"context"
	"time"
	"gorm.io/gorm"
	"github.com/zacus/barshop-server/internal/models"
	"github.com/zacus/barshop-server/internal/repository/repointerface"
)

// AppointmentRepository PostgreSQL预约仓储实现
type AppointmentRepository struct {
	*BaseRepository
}

// NewAppointmentRepository 创建预约仓储实例
func NewAppointmentRepository(db *gorm.DB) repointerface.AppointmentRepository {
	return &AppointmentRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// FindByID 根据ID查找预约
func (r *AppointmentRepository) FindByID(ctx context.Context, id uint) (*models.Appointment, error) {
	var appointment models.Appointment
	if err := r.GetDB().WithContext(ctx).
		Preload("Customer").
		Preload("Barber").
		Preload("Service").
		First(&appointment, id).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

// ListByCustomer 获取客户的预约列表
func (r *AppointmentRepository) ListByCustomer(ctx context.Context, customerID uint) ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	if err := r.GetDB().WithContext(ctx).
		Preload("Customer").
		Preload("Barber").
		Preload("Service").
		Where("customer_id = ?", customerID).
		Order("start_time desc").
		Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

// ListByBarber 获取理发师的预约列表
func (r *AppointmentRepository) ListByBarber(ctx context.Context, barberID uint) ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	if err := r.GetDB().WithContext(ctx).
		Preload("Customer").
		Preload("Barber").
		Preload("Service").
		Where("barber_id = ?", barberID).
		Order("start_time desc").
		Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

// ListByDateRange 获取指定日期范围的预约列表
func (r *AppointmentRepository) ListByDateRange(ctx context.Context, barberID uint, start, end time.Time) ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	if err := r.GetDB().WithContext(ctx).
		Preload("Customer").
		Preload("Barber").
		Preload("Service").
		Where("barber_id = ? AND start_time BETWEEN ? AND ?", barberID, start, end).
		Order("start_time").
		Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

// ListByStatus 获取指定状态的预约列表
func (r *AppointmentRepository) ListByStatus(ctx context.Context, status string) ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	if err := r.GetDB().WithContext(ctx).
		Preload("Customer").
		Preload("Barber").
		Preload("Service").
		Where("status = ?", status).
		Order("start_time desc").
		Find(&appointments).Error; err != nil {
		return nil, err
	}
	return appointments, nil
}

// CheckTimeConflict 检查时间冲突
func (r *AppointmentRepository) CheckTimeConflict(ctx context.Context, barberID uint, start, end time.Time) (bool, error) {
	var count int64
	err := r.GetDB().WithContext(ctx).Model(&models.Appointment{}).
		Where("barber_id = ? AND status != ? AND ((start_time BETWEEN ? AND ?) OR (end_time BETWEEN ? AND ?))",
			barberID, "cancelled", start, end, start, end).
		Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// Create 创建预约
func (r *AppointmentRepository) Create(ctx context.Context, appointment *models.Appointment) error {
	return r.BaseRepository.Create(ctx, appointment)
}

// Update 更新预约
func (r *AppointmentRepository) Update(ctx context.Context, appointment *models.Appointment) error {
	return r.BaseRepository.Update(ctx, appointment)
}

// Delete 删除预约
func (r *AppointmentRepository) Delete(ctx context.Context, id uint) error {
	return r.BaseRepository.Delete(ctx, &models.Appointment{ID: id})
} 