package services

import (
	"errors"
	"github.com/zacus/barshop-server/internal/database"
	"github.com/zacus/barshop-server/internal/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

// Register 用户注册
func (s *UserService) Register(req *models.RegisterRequest) (*models.User, error) {
	// 检查用户名是否已存在
	var existingUser models.User
	if err := database.DB.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return nil, errors.New("用户名已存在")
	} else if err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建新用户
	user := &models.User{
		Username: req.Username,
		Password: string(hashedPassword),
		Name:     req.Name,
		Phone:    req.Phone,
		Email:    req.Email,
		Role:     "customer", // 默认角色为顾客
	}

	if err := database.DB.Create(user).Error; err != nil {
		return nil, err
	}

	user.Password = "" // 清空密码，不返回给客户端
	return user, nil
}

// Login 用户登录
func (s *UserService) Login(req *models.LoginRequest) (*models.User, error) {
	var user models.User
	if err := database.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("密码错误")
	}

	user.Password = "" // 清空密码，不返回给客户端
	return &user, nil
}

// GetUserByID 根据ID获取用户信息
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, err
	}

	user.Password = "" // 清空密码，不返回给客户端
	return &user, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(id uint, req *models.UpdateUserRequest) (*models.User, error) {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return nil, err
	}

	// 更新用户信息
	updates := map[string]interface{}{
		"name":  req.Name,
		"phone": req.Phone,
		"email": req.Email,
	}

	if err := database.DB.Model(&user).Updates(updates).Error; err != nil {
		return nil, err
	}

	user.Password = "" // 清空密码，不返回给客户端
	return &user, nil
}

// ListBarbers 获取理发师列表
func (s *UserService) ListBarbers() ([]models.User, error) {
	var barbers []models.User
	if err := database.DB.Where("role = ?", "barber").Find(&barbers).Error; err != nil {
		return nil, err
	}

	// 清空所有理发师的密码
	for i := range barbers {
		barbers[i].Password = ""
	}

	return barbers, nil
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(id uint, oldPassword, newPassword string) error {
	var user models.User
	if err := database.DB.First(&user, id).Error; err != nil {
		return err
	}

	// 验证旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("旧密码错误")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	return database.DB.Model(&user).Update("password", string(hashedPassword)).Error
} 