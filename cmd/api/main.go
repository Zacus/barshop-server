package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zacus/barshop-server/internal/config"
	"github.com/zacus/barshop-server/internal/database"
	"github.com/zacus/barshop-server/internal/handlers"
	"github.com/zacus/barshop-server/internal/logger"
	"github.com/zacus/barshop-server/internal/middleware"
	"github.com/zacus/barshop-server/internal/services"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "github.com/zacus/barshop-server/docs"
	"os"
	"os/signal"
	"syscall"
)

// @title           理发店管理系统 API
// @version         1.0
// @description     这是一个理发店管理系统的API文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description 请在此输入 Bearer {token} 格式的JWT令牌

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// 初始化日志
	if err := logger.InitLogger(cfg.Log.Level, cfg.Log.IsDev); err != nil {
		fmt.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// 初始化数据库
	if err := database.InitDB(cfg); err != nil {
		logger.Fatal("Failed to initialize database", "error", err)
	}

	// 设置gin模式
	gin.SetMode(cfg.Server.Mode)

	// 创建路由
	r := gin.Default()

	// 基础中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS()) // 添加CORS中间件

	// Swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 创建服务实例
	userService := services.NewUserService()
	serviceService := services.NewServiceService()
	appointmentService := services.NewAppointmentService()

	// 创建处理器实例
	userHandler := handlers.NewUserHandler(userService)
	serviceHandler := handlers.NewServiceHandler(serviceService)
	appointmentHandler := handlers.NewAppointmentHandler(appointmentService)

	// API路由组
	api := r.Group("/api/v1")
	{
		// 认证路由（无需登录）
		auth := api.Group("/auth")
		{
			auth.POST("/register", userHandler.Register)
			auth.POST("/login", userHandler.Login)
		}

		// 公开路由
		public := api.Group("/")
		{
			public.GET("/health", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"status": "ok",
					"message": "服务正常运行",
				})
			})

			// 服务相关路由
			public.GET("/services", serviceHandler.List)
		}

		// 需要认证的路由
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWT.Secret))
		{
			// 用户相关路由
			users := protected.Group("/users")
			{
				users.GET("/profile", userHandler.GetProfile)
				users.PUT("/profile", userHandler.UpdateProfile)
				users.PUT("/password", userHandler.ChangePassword)
				users.GET("/barbers", userHandler.ListBarbers)
			}

			// 预约相关路由
			appointments := protected.Group("/appointments")
			{
				appointments.POST("/", appointmentHandler.Create)
				appointments.GET("/", appointmentHandler.List)
				appointments.PUT("/:id/status", appointmentHandler.UpdateStatus)
			}

			// 服务管理路由（仅管理员）
			services := protected.Group("/admin/services")
			services.Use(middleware.AdminOnly())
			{
				services.POST("/", serviceHandler.Create)
				services.PUT("/:id", serviceHandler.Update)
				services.DELETE("/:id", serviceHandler.Delete)
				services.PUT("/:id/toggle", serviceHandler.ToggleStatus)
			}
		}
	}

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		addr := fmt.Sprintf(":%s", cfg.Server.Port)
		logger.Info("Server starting", "address", addr)
		if err := r.Run(addr); err != nil {
			logger.Fatal("Failed to start server", "error", err)
		}
	}()

	<-quit
	logger.Info("Shutting down server...")
} 