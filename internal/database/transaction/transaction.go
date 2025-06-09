/*
 * @Author: zs
 * @Date: 2025-06-08 16:20:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-08 16:20:37
 * @FilePath: /barshop-server/internal/database/transaction/transaction.go
 * @Description: 事务管理接口定义
 */
package transaction

// Transaction 事务接口
type Transaction interface {
	// Commit 提交事务
	Commit() error
	
	// Rollback 回滚事务
	Rollback() error
	
	// GetTx 获取事务实例
	GetTx() interface{}
} 