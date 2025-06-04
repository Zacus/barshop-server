package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/barshop-server/internal/models"
	"github.com/yourusername/barshop-server/internal/services"
	"github.com/yourusername/barshop-server/internal/utils"
	"net/http"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Register 用户注册
// @Summary 用户注册
// @Description 注册新用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.RegisterRequest true "用户注册信息"
// @Success 200 {object} models.Response{data=models.User}
// @Failure 400,500 {object} models.Response
// @Router /auth/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user, err := h.service.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(user))
}

// Login 用户登录
// @Summary 用户登录
// @Description 用户登录并获取令牌
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.LoginRequest true "登录信息"
// @Success 200 {object} models.Response{data=map[string]interface{}}
// @Failure 400,401 {object} models.Response
// @Router /auth/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user, err := h.service.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.ErrorResponse(http.StatusUnauthorized, err.Error()))
		return
	}

	// 生成JWT令牌
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, "生成令牌失败"))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(gin.H{
		"token": token,
		"user":  user,
	}))
}

// GetProfile 获取用户信息
// @Summary 获取用户信息
// @Description 获取当前登录用户的信息
// @Tags 用户管理
// @Produce json
// @Security Bearer
// @Success 200 {object} models.Response{data=models.User}
// @Failure 401,404 {object} models.Response
// @Router /users/profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	user, err := h.service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse(http.StatusNotFound, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(user))
}

// UpdateProfile 更新用户信息
// @Summary 更新用户信息
// @Description 更新当前登录用户的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body models.UpdateUserRequest true "用户信息"
// @Security Bearer
// @Success 200 {object} models.Response{data=models.User}
// @Failure 400,401,404 {object} models.Response
// @Router /users/profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	user, err := h.service.UpdateUser(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(user))
}

// ListBarbers 获取理发师列表
// @Summary 获取理发师列表
// @Description 获取所有理发师的列表
// @Tags 用户管理
// @Produce json
// @Security Bearer
// @Success 200 {object} models.Response{data=[]models.User}
// @Router /users/barbers [get]
func (h *UserHandler) ListBarbers(c *gin.Context) {
	barbers, err := h.service.ListBarbers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(barbers))
}

// ChangePassword 修改密码
// @Summary 修改密码
// @Description 修改当前登录用户的密码
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param passwords body models.ChangePasswordRequest true "密码信息"
// @Security Bearer
// @Success 200 {object} models.Response
// @Failure 400,401,500 {object} models.Response
// @Router /users/password [put]
func (h *UserHandler) ChangePassword(c *gin.Context) {
	userID := c.GetUint("user_id")
	var req models.ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	if err := h.service.ChangePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(nil))
} 