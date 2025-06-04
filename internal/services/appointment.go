package services

import (
	"errors"
	"github.com/zacus/barshop-server/internal/database"
	"github.com/zacus/barshop-server/internal/models"
	"time"
)

type AppointmentService struct{}

func NewAppointmentService() *AppointmentService {
	return &AppointmentService{}
}

// CreateAppointment 创建预约
func (s *AppointmentService) CreateAppointment(customerID uint, req *models.AppointmentRequest) (*models.Appointment, error) {
	// 获取服务信息
	service := &models.Service{}
	if err := database.DB.First(service, req.ServiceID).Error; err != nil {
		return nil, err
	}

	// 计算结束时间
	endTime := req.StartTime.Add(time.Duration(service.Duration) * time.Minute)

	// 检查理发师在该时间段是否有其他预约
	var conflictCount int64
	err := database.DB.Model(&models.Appointment{}).
		Where("barber_id = ? AND status != 'cancelled' AND "+
			"((start_time BETWEEN ? AND ?) OR "+
			"(end_time BETWEEN ? AND ?) OR "+
			"(start_time <= ? AND end_time >= ?))",
			req.BarberID,
			req.StartTime, endTime,
			req.StartTime, endTime,
			req.StartTime, endTime).
		Count(&conflictCount).Error
	if err != nil {
		return nil, err
	}

	if conflictCount > 0 {
		return nil, errors.New("理发师在该时间段已有预约")
	}

	// 创建预约
	appointment := &models.Appointment{
		CustomerID: customerID,
		BarberID:   req.BarberID,
		ServiceID:  req.ServiceID,
		StartTime:  req.StartTime,
		EndTime:    endTime,
		Status:     "pending",
		Note:       req.Note,
	}

	if err := database.DB.Create(appointment).Error; err != nil {
		return nil, err
	}

	// 加载关联数据
	if err := database.DB.Preload("Customer").Preload("Barber").Preload("Service").First(appointment, appointment.ID).Error; err != nil {
		return nil, err
	}

	return appointment, nil
}

// GetAppointments 获取预约列表
func (s *AppointmentService) GetAppointments(userID uint, role string, status string, startDate, endDate time.Time) ([]models.Appointment, error) {
	var appointments []models.Appointment
	query := database.DB.Preload("Customer").Preload("Barber").Preload("Service")

	// 根据角色筛选
	switch role {
	case "customer":
		query = query.Where("customer_id = ?", userID)
	case "barber":
		query = query.Where("barber_id = ?", userID)
	}

	// 根据状态筛选
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// 根据日期范围筛选
	query = query.Where("start_time BETWEEN ? AND ?", startDate, endDate)

	// 按时间排序
	query = query.Order("start_time desc")

	if err := query.Find(&appointments).Error; err != nil {
		return nil, err
	}

	return appointments, nil
}

// UpdateAppointmentStatus 更新预约状态
func (s *AppointmentService) UpdateAppointmentStatus(id uint, status string) (*models.Appointment, error) {
	appointment := &models.Appointment{}
	if err := database.DB.First(appointment, id).Error; err != nil {
		return nil, err
	}

	appointment.Status = status
	if err := database.DB.Save(appointment).Error; err != nil {
		return nil, err
	}

	// 重新加载关联数据
	if err := database.DB.Preload("Customer").Preload("Barber").Preload("Service").First(appointment, id).Error; err != nil {
		return nil, err
	}

	return appointment, nil
} 