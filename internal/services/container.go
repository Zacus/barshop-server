/*
 * @Author: zs
 * @Date: 2025-06-07 16:36:23
 * @LastEditors: zs
 * @LastEditTime: 2025-06-09 14:34:00
 * @FilePath: /barshop-server/internal/services/container.go
 * @Description: 服务容器管理
 */
package services

import (
	"github.com/zacus/barshop-server/internal/repository/repointerface"
)

// ContainerConfig 服务容器配置
type ContainerConfig struct {
	UserRepo        repointerface.UserRepository
	ServiceRepo     repointerface.ServiceRepository
	AppointmentRepo repointerface.AppointmentRepository
}

// Container 服务容器，用于管理所有服务实例
type Container struct {
	User        *UserService
	Service     *ServiceService
	Appointment *AppointmentService
}

// NewContainer 创建新的服务容器
func NewContainer(config ContainerConfig) *Container {
	return &Container{
		User:        NewUserService(config.UserRepo),
		Service:     NewServiceService(config.ServiceRepo),
		Appointment: NewAppointmentService(config.AppointmentRepo, config.ServiceRepo),
	}
}