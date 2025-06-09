/*
 * @Author: zs
 * @Date: 2025-06-08 16:20:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-08 16:20:37
 * @FilePath: /barshop-server/internal/database/errors.go
 * @Description: 数据库错误定义
 */
package database

import (
	"errors"
)

var (
	// ErrConnectionFailed 数据库连接失败
	ErrConnectionFailed = errors.New("database connection failed")
	
	// ErrTransactionFailed 事务执行失败
	ErrTransactionFailed = errors.New("transaction failed")
	
	// ErrRecordNotFound 记录未找到
	ErrRecordNotFound = errors.New("record not found")
)

// DBError 数据库错误
type DBError struct {
	Err     error
	Message string
	Code    string
}

func (e *DBError) Error() string {
	return e.Message
}

// NewDBError 创建数据库错误
func NewDBError(err error, message string, code string) *DBError {
	return &DBError{
		Err:     err,
		Message: message,
		Code:    code,
	}
} 