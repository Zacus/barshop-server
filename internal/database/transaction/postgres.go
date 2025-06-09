/*
 * @Author: zs
 * @Date: 2025-06-08 16:30:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-08 16:30:37
 * @FilePath: /barshop-server/internal/database/transaction/postgres.go
 * @Description: PostgreSQL 事务实现
 */
package transaction

import (
	"gorm.io/gorm"
)

// PostgresTransaction PostgreSQL事务实现
type PostgresTransaction struct {
	tx *gorm.DB
}

// NewPostgresTransaction 创建PostgreSQL事务实例
func NewPostgresTransaction(tx *gorm.DB) Transaction {
	return &PostgresTransaction{
		tx: tx,
	}
}

// Commit 提交事务
func (t *PostgresTransaction) Commit() error {
	return t.tx.Commit().Error
}

// Rollback 回滚事务
func (t *PostgresTransaction) Rollback() error {
	return t.tx.Rollback().Error
}

// GetTx 获取事务实例
func (t *PostgresTransaction) GetTx() interface{} {
	return t.tx
} 