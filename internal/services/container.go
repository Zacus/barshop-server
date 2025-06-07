package services

// Container 服务容器，用于管理所有服务实例
type Container struct {
	User        *UserService
	Service     *ServiceService
	Appointment *AppointmentService
}

// NewContainer 创建新的服务容器
func NewContainer() *Container {
	return &Container{
		User:        NewUserService(),
		Service:     NewServiceService(),
		Appointment: NewAppointmentService(),
	}
} 