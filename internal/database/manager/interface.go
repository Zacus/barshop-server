/*
 * @Author: zs
 * @Date: 2025-06-08 16:20:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-08 23:20:49
 * @FilePath: /barshop-server/internal/database/manager/interface.go
 * @Description: 数据库管理器接口定义
 */
package manager

import (
	"context"
	"github.com/zacus/barshop-server/internal/config"
	"github.com/zacus/barshop-server/internal/repository"
	"github.com/zacus/barshop-server/internal/database/transaction"
)

// DBManager 数据库管理器接口
type DBManager interface {
	// Initialize 初始化数据库连接
	Initialize(cfg *config.Config) error
	
	// Close 关闭数据库连接
	Close() error
	
	// Health 检查数据库健康状态
	Health() error
	
	// GetDB 获取数据库连接
	GetDB() (interface{}, error)

	// BeginTx 开启事务
	BeginTx(ctx context.Context) (transaction.Transaction, error)
	
	// Transaction 事务执行
	Transaction(ctx context.Context, fn func(transaction.Transaction) error) error

	// GetRepository 获取仓储工厂
	GetRepository() repository.Factory
} 