/*
 * @Author: zs
 * @Date: 2025-06-08 16:40:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 01:39:37
 * @FilePath: /barshop-server/internal/repository/postgres/user.go
 * @Description: 用户仓储PostgreSQL实现
 */
package postgres

import (
	"context"
	"gorm.io/gorm"
	"github.com/zacus/barshop-server/internal/models"
	"github.com/zacus/barshop-server/internal/repository/repointerface"
)

// UserRepository PostgreSQL用户仓储实现
type UserRepository struct {
	*BaseRepository
}

// NewUserRepository 创建用户仓储实例
func NewUserRepository(db *gorm.DB) repointerface.UserRepository {
	return &UserRepository{
		BaseRepository: NewBaseRepository(db),
	}
}

// FindByID 根据ID查找用户
func (r *UserRepository) FindByID(ctx context.Context, id uint) (*models.User, error) {
	var user models.User
	if err := r.BaseRepository.FindByID(ctx, id, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername 根据用户名查找用户
func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	if err := r.GetDB().WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByPhone 根据手机号查找用户
func (r *UserRepository) FindByPhone(ctx context.Context, phone string) (*models.User, error) {
	var user models.User
	if err := r.GetDB().WithContext(ctx).Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail 根据邮箱查找用户
func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	if err := r.GetDB().WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// ListBarbers 获取所有理发师列表
func (r *UserRepository) ListBarbers(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	if err := r.GetDB().WithContext(ctx).Where("role = ?", "barber").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Create 创建用户
func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.BaseRepository.Create(ctx, user)
}

// Update 更新用户
func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	return r.BaseRepository.Update(ctx, user)
}

// Delete 删除用户
func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.BaseRepository.Delete(ctx, &models.User{ID: id})
} 