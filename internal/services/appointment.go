/*
 * @Author: zs
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 16:11:40
 * @FilePath: /barshop-server/internal/services/appointment.go
 * @Description: 
 * 
 * Copyright (c) 2025 by zs, All Rights Reserved. 
 */
package services

import (
	"context"
	"errors"
	"fmt"
	"time"
	"github.com/zacus/barshop-server/internal/models"
	"github.com/zacus/barshop-server/internal/repository/repointerface"
)

type AppointmentService struct {
	repo            repointerface.AppointmentRepository
	serviceRepo     repointerface.ServiceRepository
}

func NewAppointmentService(
	repo repointerface.AppointmentRepository,
	serviceRepo repointerface.ServiceRepository,
) *AppointmentService {
	return &AppointmentService{
		repo:        repo,
		serviceRepo: serviceRepo,
	}
}

// CreateAppointment 创建预约
func (s *AppointmentService) CreateAppointment(ctx context.Context,customerID uint, req *models.AppointmentRequest) (*models.Appointment, error) {

	// 获取服务信息以获取持续时间
	service, err := s.serviceRepo.FindByID(ctx, req.ServiceID)
	if err != nil {
		return nil, fmt.Errorf("获取服务信息失败: %w", err)
	}

	// 计算结束时间
	endTime := req.StartTime.Add(time.Duration(service.Duration) * time.Minute)

	// 检查时间冲突
	hasConflict, err := s.repo.CheckTimeConflict(ctx, req.BarberID, req.StartTime, endTime)
	if err != nil {
		return nil, err
	}
	if hasConflict {
		return nil, errors.New("理发师在该时间段已有预约")
	}

	// 创建预约
	appointment := &models.Appointment{
		CustomerID:     customerID,
		BarberID:      req.BarberID,
		ServiceID:     req.ServiceID,
		StartTime:     req.StartTime,
		EndTime:       endTime,
		Status:        "pending",
		Note:          req.Note,
		Price:         service.Price,
		PaymentStatus: "unpaid",
	}

	if err := s.repo.Create(ctx, appointment); err != nil {
		return nil, err
	}

	// 获取完整预约信息
	return s.repo.FindByID(ctx, appointment.ID)
}

// GetAppointments 获取预约列表
func (s *AppointmentService) GetAppointments(ctx context.Context,userID uint, role string, status string, startDate, endDate time.Time) ([]*models.Appointment, error) {
	var appointments []*models.Appointment
	var err error

	// 根据角色获取不同的预约列表
	switch role {
	case "customer":
		appointments, err = s.repo.ListByCustomer(ctx, userID)
	case "barber":
		if status != "" {
			appointments, err = s.repo.ListByStatus(ctx, status)
		} else {
			appointments, err = s.repo.ListByBarber(ctx, userID)
		}
	default:
		// 管理员可以看到按日期范围的所有预约
		appointments, err = s.repo.ListByDateRange(ctx, 0, startDate, endDate)
	}

	if err != nil {
		return nil, err
	}

	return appointments, nil
}

// UpdateAppointmentStatus 更新预约状态
func (s *AppointmentService) UpdateAppointmentStatus(ctx context.Context,id uint, status string) (*models.Appointment, error) {

	// 先获取预约信息
	appointment, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 更新状态
	appointment.Status = status
	if err := s.repo.Update(ctx, appointment); err != nil {
		return nil, err
	}

	return s.repo.FindByID(ctx, id)
}