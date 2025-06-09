/*
 * @Author: zs
 * @Date: 2025-06-08 16:20:37
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 00:59:11
 * @FilePath: /barshop-server/internal/repository/interface/repository.go
 * @Description: 基础仓储接口定义
 */
package repointerface

import (
	"context"
)

// Repository 基础仓储接口
type Repository interface {
	// Create 创建记录
	Create(ctx context.Context, model interface{}) error
	
	// FindByID 根据ID查找记录
	FindByID(ctx context.Context, id uint, model interface{}) error
	
	// Update 更新记录
	Update(ctx context.Context, model interface{}) error
	
	// Delete 删除记录
	Delete(ctx context.Context, model interface{}) error
}