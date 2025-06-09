/*
 * @Author: zs
 * @Date: 2025-06-08 16:40:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-08 16:40:37
 * @FilePath: /barshop-server/internal/database/repository/postgres/base.go
 * @Description: PostgreSQL基础仓储实现
 */
package postgres

import (
	"context"
	"gorm.io/gorm"
)

// BaseRepository PostgreSQL基础仓储实现
type BaseRepository struct {
	db *gorm.DB
}

// NewBaseRepository 创建基础仓储实例
func NewBaseRepository(db *gorm.DB) *BaseRepository {
	return &BaseRepository{
		db: db,
	}
}

// Create 创建记录
func (r *BaseRepository) Create(ctx context.Context, model interface{}) error {
	return r.db.WithContext(ctx).Create(model).Error
}

// FindByID 根据ID查找记录
func (r *BaseRepository) FindByID(ctx context.Context, id uint, model interface{}) error {
	return r.db.WithContext(ctx).First(model, id).Error
}

// Update 更新记录
func (r *BaseRepository) Update(ctx context.Context, model interface{}) error {
	return r.db.WithContext(ctx).Save(model).Error
}

// Delete 删除记录
func (r *BaseRepository) Delete(ctx context.Context, model interface{}) error {
	return r.db.WithContext(ctx).Delete(model).Error
}

// GetDB 获取数据库连接
func (r *BaseRepository) GetDB() *gorm.DB {
	return r.db
} 