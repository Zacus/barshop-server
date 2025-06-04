package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zacus/barshop-server/internal/models"
	"github.com/zacus/barshop-server/internal/services"
	"net/http"
	"strconv"
	"time"
)

type AppointmentHandler struct {
	service *services.AppointmentService
}

func NewAppointmentHandler(service *services.AppointmentService) *AppointmentHandler {
	return &AppointmentHandler{service: service}
}

// Create 创建预约
// @Summary 创建新预约
// @Description 创建一个新的理发预约
// @Tags 预约管理
// @Accept json
// @Produce json
// @Param appointment body models.AppointmentRequest true "预约信息"
// @Security Bearer
// @Success 200 {object} models.Response{data=models.Appointment}
// @Failure 400,409 {object} models.Response
// @Router /appointments [post]
func (h *AppointmentHandler) Create(c *gin.Context) {
	var req models.AppointmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	customerID := c.GetUint("user_id")
	appointment, err := h.service.CreateAppointment(customerID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(appointment))
}

// List 获取预约列表
// @Summary 获取预约列表
// @Description 获取用户相关的预约列表
// @Tags 预约管理
// @Produce json
// @Param status query string false "预约状态"
// @Param start_date query string false "开始日期 (YYYY-MM-DD)"
// @Param end_date query string false "结束日期 (YYYY-MM-DD)"
// @Security Bearer
// @Success 200 {object} models.Response{data=[]models.Appointment}
// @Router /appointments [get]
func (h *AppointmentHandler) List(c *gin.Context) {
	userID := c.GetUint("user_id")
	role := c.GetString("user_role")
	status := c.Query("status")

	startDate := time.Now().Truncate(24 * time.Hour)
	endDate := startDate.AddDate(0, 1, 0)

	if startDateStr := c.Query("start_date"); startDateStr != "" {
		if date, err := time.Parse("2006-01-02", startDateStr); err == nil {
			startDate = date
		}
	}

	if endDateStr := c.Query("end_date"); endDateStr != "" {
		if date, err := time.Parse("2006-01-02", endDateStr); err == nil {
			endDate = date.Add(24 * time.Hour)
		}
	}

	appointments, err := h.service.GetAppointments(userID, role, status, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(appointments))
}

// UpdateStatus 更新预约状态
// @Summary 更新预约状态
// @Description 更新指定预约的状态
// @Tags 预约管理
// @Accept json
// @Produce json
// @Param id path int true "预约ID"
// @Param status body string true "新状态"
// @Security Bearer
// @Success 200 {object} models.Response{data=models.Appointment}
// @Failure 400,404 {object} models.Response
// @Router /appointments/{id}/status [put]
func (h *AppointmentHandler) UpdateStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, "Invalid appointment ID"))
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	appointment, err := h.service.UpdateAppointmentStatus(uint(id), req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(appointment))
} 