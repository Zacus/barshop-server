/*
 * @Author: zs
 * @Date: 2025-06-08 00:30:27
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 16:48:16
 * @FilePath: /barshop-server/internal/services/service.go
 * @Description: 服务管理
 */
 package services

 import (
	 "context"
	 "github.com/zacus/barshop-server/internal/models"
	 "github.com/zacus/barshop-server/internal/repository/repointerface"
 )
 
 // ServiceService 服务管理服务
 type ServiceService struct {
	 repo repointerface.ServiceRepository
 }
 
 // NewServiceService 创建服务管理服务实例
 func NewServiceService(repo repointerface.ServiceRepository) *ServiceService {
	 return &ServiceService{repo: repo}
 }
 
 // CreateService 创建新服务
 func (s *ServiceService) CreateService(ctx context.Context, req *models.ServiceRequest) (*models.Service, error) {
	 service := &models.Service{
		 Name:        req.Name,
		 Description: req.Description,
		 Duration:    req.Duration,
		 Price:       req.Price,
		 IsActive:    true,
	 }
 
	 if err := s.repo.Create(ctx, service); err != nil {
		 return nil, err
	 }
 
	 return service, nil
 }
 
 // GetServices 获取所有服务列表
 func (s *ServiceService) GetServices(ctx context.Context, isActive bool) ([]*models.Service, error) {
	if isActive {
		return s.repo.ListActive(ctx)
	}
	return s.repo.List(ctx)
 }
 
 // GetServiceByID 根据ID获取服务
 func (s *ServiceService) GetServiceByID(ctx context.Context, id uint)(*models.Service, error) {
	 return s.repo.FindByID(ctx, id)
 }
 
 // UpdateService 更新服务信息
 func (s *ServiceService) UpdateService(ctx context.Context, id uint, req *models.ServiceRequest) (*models.Service, error) {
	 service := &models.Service{
		 ID:          id,
		 Name:        req.Name,
		 Description: req.Description,
		 Duration:    req.Duration,
		 Price:       req.Price,
	 }
	 return service,s.repo.Update(ctx, service)
 }
 
 // DeleteService 删除服务
 func (s *ServiceService) DeleteService(ctx context.Context, id uint) error {
	 return s.repo.Delete(ctx, id)
 }

 // ToggleServiceStatus 切换服务状态（启用/禁用）
func (s *ServiceService) ToggleServiceStatus(ctx context.Context, id uint) (*models.Service, error) {
    // 获取当前服务
    service, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, err
    }

    // 切换状态
    service.IsActive = !service.IsActive

    // 保存更新
    if err := s.repo.Update(ctx, service); err != nil {
        return nil, err
    }

    return service, nil
}
 
//  // UpdateServiceStatus 更新服务状态
//  func (s *ServiceService) UpdateServiceStatus(ctx context.Context, id uint, isActive bool) error {
// 	 return s.repo.UpdateStatus(ctx, id, isActive)
//  }