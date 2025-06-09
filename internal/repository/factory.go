/*
 * @Author: zs
 * @Date: 2025-06-08 16:40:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 01:33:28
 * @FilePath: /barshop-server/internal/repository/factory.go
 * @Description: 仓储工厂
 */
package repository

import (
	"gorm.io/gorm"
	"github.com/zacus/barshop-server/internal/repository/repointerface"
	"github.com/zacus/barshop-server/internal/repository/postgres"
)

// Factory 仓储工厂接口
type Factory interface {
	// UserRepository 获取用户仓储
	UserRepository() repointerface.UserRepository
	
	// ServiceRepository 获取服务仓储
	ServiceRepository() repointerface.ServiceRepository
	
	// AppointmentRepository 获取预约仓储
	AppointmentRepository() repointerface.AppointmentRepository
}

// PostgresFactory PostgreSQL仓储工厂实现
type PostgresFactory struct {
	db *gorm.DB
}

// NewPostgresFactory 创建PostgreSQL仓储工厂
func NewPostgresFactory(db *gorm.DB) Factory {
	return &PostgresFactory{
		db: db,
	}
}

// UserRepository 获取用户仓储
func (f *PostgresFactory) UserRepository() repointerface.UserRepository {
	return postgres.NewUserRepository(f.db)
}

// ServiceRepository 获取服务仓储
func (f *PostgresFactory) ServiceRepository() repointerface.ServiceRepository {
	return postgres.NewServiceRepository(f.db)
}

// AppointmentRepository 获取预约仓储
func (f *PostgresFactory) AppointmentRepository() repointerface.AppointmentRepository {
	return postgres.NewAppointmentRepository(f.db)
} 