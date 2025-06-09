/*
 * @Author: zs
 * @Date: 2025-06-08 16:50:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 15:43:25
 * @FilePath: /barshop-server/internal/repository/repointerface/service.go
 * @Description: 服务仓储接口定义
 */
package repointerface

import (
	"context"

	"github.com/zacus/barshop-server/internal/models"
)

// ServiceRepository 服务仓储接口
type ServiceRepository interface {
	// 基础CRUD
	Create(ctx context.Context, service *models.Service) error
	FindByID(ctx context.Context, id uint) (*models.Service, error)
	Update(ctx context.Context, service *models.Service) error
	Delete(ctx context.Context, id uint) error

	// 业务查询
	List(ctx context.Context) ([]*models.Service, error)
	ListActive(ctx context.Context) ([]*models.Service, error)
	ListByCategory(ctx context.Context, categoryID uint) ([]*models.Service, error)
	FindByName(ctx context.Context, name string) (*models.Service, error)
}