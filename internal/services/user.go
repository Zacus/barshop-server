package services

import (
	"context"
	"errors"
	"github.com/zacus/barshop-server/internal/models"
	"github.com/zacus/barshop-server/internal/repository/repointerface"
	"github.com/zacus/barshop-server/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repointerface.UserRepository
}

func NewUserService(repo repointerface.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

// Register 用户注册
func (s *UserService) Register(ctx context.Context,user *models.User) error {

	// 检查用户名是否已存在
	existingUser, err := s.repo.FindByUsername(ctx, user.Username)
	if err == nil && existingUser != nil {
		return errors.New("用户名已存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// 创建新用户
	return s.repo.Create(ctx, user)
}

// Login 用户登录
func (s *UserService) Login(ctx context.Context,username, password string) (string, error) {

	// 查找用户
	user, err := s.repo.FindByUsername(ctx, username)
	if err != nil {
		return "", errors.New("用户不存在")
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("密码错误")
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// GetUserByID 根据ID获取用户信息
func (s *UserService) GetUserByID(ctx context.Context,id uint) (*models.User, error) {

	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	user.Password = "" // 清空密码，不返回给客户端
	return user, nil
}

// UpdateProfile 更新用户信息
func (s *UserService) UpdateProfile(ctx context.Context,id uint, req *models.UpdateUserRequest) error {
	
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// 更新用户信息
	user.Name = req.Name
	user.Phone = req.Phone
	user.Email = req.Email

	return s.repo.Update(ctx, user)
}

// ListBarbers 获取理发师列表
func (s *UserService) ListBarbers(ctx context.Context) ([]*models.User, error) {
	barbers, err := s.repo.ListBarbers(ctx)
	if err != nil {
		return nil, err
	}

	// 清空所有理发师的密码
	for _, barber := range barbers {
		barber.Password = ""
	}

	return barbers, nil
}

// ChangePassword 修改密码
func (s *UserService) ChangePassword(ctx context.Context,id uint, oldPassword, newPassword string) error {

	// 获取用户信息
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	// 验证旧密码
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(oldPassword)); err != nil {
		return errors.New("原密码错误")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 更新密码
	user.Password = string(hashedPassword)
	return s.repo.Update(ctx, user)
}