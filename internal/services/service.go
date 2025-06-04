package services

import (
	"github.com/zacus/barshop-server/internal/database"
	"github.com/zacus/barshop-server/internal/models"
)

type ServiceService struct{}

func NewServiceService() *ServiceService {
	return &ServiceService{}
}

// CreateService 创建新服务
func (s *ServiceService) CreateService(req *models.ServiceRequest) (*models.Service, error) {
	service := &models.Service{
		Name:        req.Name,
		Description: req.Description,
		Duration:    req.Duration,
		Price:       req.Price,
		IsActive:    true,
	}

	if err := database.DB.Create(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

// GetServices 获取所有服务列表
func (s *ServiceService) GetServices(onlyActive bool) ([]models.Service, error) {
	var services []models.Service
	query := database.DB

	if onlyActive {
		query = query.Where("is_active = ?", true)
	}

	if err := query.Find(&services).Error; err != nil {
		return nil, err
	}

	return services, nil
}

// UpdateService 更新服务信息
func (s *ServiceService) UpdateService(id uint, req *models.ServiceRequest) (*models.Service, error) {
	service := &models.Service{}
	if err := database.DB.First(service, id).Error; err != nil {
		return nil, err
	}

	service.Name = req.Name
	service.Description = req.Description
	service.Duration = req.Duration
	service.Price = req.Price

	if err := database.DB.Save(service).Error; err != nil {
		return nil, err
	}

	return service, nil
}

// DeleteService 删除服务（软删除）
func (s *ServiceService) DeleteService(id uint) error {
	return database.DB.Delete(&models.Service{}, id).Error
}

// ToggleServiceStatus 切换服务状态（启用/禁用）
func (s *ServiceService) ToggleServiceStatus(id uint) (*models.Service, error) {
	service := &models.Service{}
	if err := database.DB.First(service, id).Error; err != nil {
		return nil, err
	}

	service.IsActive = !service.IsActive
	if err := database.DB.Save(service).Error; err != nil {
		return nil, err
	}

	return service, nil
} 