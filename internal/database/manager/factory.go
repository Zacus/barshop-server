/*
 * @Author: zs
 * @Date: 2025-06-08 16:30:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-08 16:30:37
 * @FilePath: /barshop-server/internal/database/manager/factory.go
 * @Description: 数据库管理器工厂
 */
package manager

import (
	"fmt"
)

const (
	// DBTypePostgres PostgreSQL数据库类型
	DBTypePostgres = "postgres"
)

// NewDBManager 创建数据库管理器实例
func NewDBManager(dbType string) (DBManager, error) {
	switch dbType {
	case DBTypePostgres:
		return NewPostgresManager(), nil
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}
} 