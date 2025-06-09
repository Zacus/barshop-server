/*
 * @Author: zs
 * @Date: 2025-06-08 16:30:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 17:35:32
 * @FilePath: /barshop-server/internal/database/manager/postgres.go
 * @Description: PostgreSQL 数据库管理器实现
 */
package manager

import (
	"context"
	"fmt"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/zacus/barshop-server/internal/config"
	"github.com/zacus/barshop-server/internal/repository"
	"github.com/zacus/barshop-server/internal/database/transaction"
	"github.com/zacus/barshop-server/internal/models"
)

// PostgresManager PostgreSQL管理器实现
type PostgresManager struct {
	db     *gorm.DB
	config *config.Config
	repoFactory repository.Factory
}

// NewPostgresManager 创建PostgreSQL管理器实例
func NewPostgresManager() DBManager {
	return &PostgresManager{}
}

// Initialize 初始化数据库连接
func (m *PostgresManager) Initialize(cfg *config.Config) error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(cfg.Database.Options.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.Database.Options.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.Database.Options.ConnMaxLifetime) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(cfg.Database.Options.ConnMaxIdleTime) * time.Minute)

	models := []interface{}{
		&models.User{},
		&models.Service{},
		&models.Appointment{},
	}
	
	if cfg.Database.AutoMigrate {
		if err := db.AutoMigrate(
			models...,
			); err != nil {
			return fmt.Errorf("failed to migrate database: %v", err)
		}
	}

	m.db = db
	m.config = cfg
	m.repoFactory = repository.NewPostgresFactory(db)
	return nil
}

// Close 关闭数据库连接
func (m *PostgresManager) Close() error {
	if m.db != nil {
		sqlDB, err := m.db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}

// Health 检查数据库健康状态
func (m *PostgresManager) Health() error {
	if m.db == nil {
		return fmt.Errorf("database connection not initialized")
	}
	
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	
	return sqlDB.Ping()
}

// GetDB 获取数据库连接
func (m *PostgresManager) GetDB() (interface{}, error) {
	if m.db == nil {
		return nil, fmt.Errorf("database connection not initialized")
	}
	return m.db, nil
}

// BeginTx 开启事务
func (m *PostgresManager) BeginTx(ctx context.Context) (transaction.Transaction, error) {
	tx := m.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return transaction.NewPostgresTransaction(tx), nil
}

// Transaction 事务执行
func (m *PostgresManager) Transaction(ctx context.Context, fn func(transaction.Transaction) error) error {
	tx, err := m.BeginTx(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

// GetRepository 获取仓储工厂
func (m *PostgresManager) GetRepository() repository.Factory {
	return m.repoFactory
} 