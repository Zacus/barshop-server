/*
 * @Author: zs
 * @Date: 2025-06-08 16:40:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 00:59:25
 * @FilePath: /barshop-server/internal/repository/interface/user.go
 * @Description: 用户仓储接口定义
 */
package repointerface

import (
	"context"
	"github.com/zacus/barshop-server/internal/models"
)

// UserRepository 用户仓储接口
type UserRepository interface {
	// 基础CRUD
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error

	// 业务查询
	FindByUsername(ctx context.Context, username string) (*models.User, error)
	FindByPhone(ctx context.Context, phone string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	ListBarbers(ctx context.Context) ([]*models.User, error)
}