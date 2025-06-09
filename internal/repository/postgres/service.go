/*
 * @Author: zs
 * @Date: 2025-06-08 16:50:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 16:32:35
 * @FilePath: /barshop-server/internal/repository/postgres/service.go
 * @Description: 服务仓储PostgreSQL实现
 */
package postgres

import (
	"context"
	"gorm.io/gorm"
	"github.com/zacus/barshop-server/internal/models"
	"github.com/zacus/barshop-server/internal/repository/repointerface"
)

// ServiceRepository PostgreSQL服务仓储实现
type ServiceRepository struct {
	*BaseRepository
}

// NewServiceRepository 创建服务仓储实例
func NewServiceRepository(db *gorm.DB) repointerface.ServiceRepository {
	return &ServiceRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// List 获取所有服务列表
func (r *ServiceRepository) List(ctx context.Context) ([]*models.Service, error) {
    var services []*models.Service
    if err := r.GetDB().WithContext(ctx).Find(&services).Error; err != nil {
        return nil, err
    }
    return services, nil
}

// FindByID 根据ID查找服务
func (r *ServiceRepository) FindByID(ctx context.Context, id uint) (*models.Service, error) {
	var service models.Service
	if err := r.BaseRepository.FindByID(ctx, id, &service); err != nil {
		return nil, err
	}
	return &service, nil
}

// ListActive 获取所有激活的服务
func (r *ServiceRepository) ListActive(ctx context.Context) ([]*models.Service, error) {
	var services []*models.Service
	if err := r.GetDB().WithContext(ctx).Where("is_active = ?", true).Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

// ListByCategory 获取指定分类的服务
func (r *ServiceRepository) ListByCategory(ctx context.Context, categoryID uint) ([]*models.Service, error) {
	var services []*models.Service
	if err := r.GetDB().WithContext(ctx).Where("category_id = ? AND is_active = ?", categoryID, true).Find(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

// FindByName 根据名称查找服务
func (r *ServiceRepository) FindByName(ctx context.Context, name string) (*models.Service, error) {
	var service models.Service
	if err := r.GetDB().WithContext(ctx).Where("name = ?", name).First(&service).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

// Create 创建服务
func (r *ServiceRepository) Create(ctx context.Context, service *models.Service) error {
	return r.BaseRepository.Create(ctx, service)
}

// Update 更新服务
func (r *ServiceRepository) Update(ctx context.Context, service *models.Service) error {
	return r.BaseRepository.Update(ctx, service)
}

// Delete 删除服务
func (r *ServiceRepository) Delete(ctx context.Context, id uint) error {
	return r.BaseRepository.Delete(ctx, &models.Service{ID: id})
} 