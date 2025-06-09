package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zacus/barshop-server/internal/models"
	"github.com/zacus/barshop-server/internal/services"
	"net/http"
	"strconv"
)

type ServiceHandler struct {
	service *services.ServiceService
}

func NewServiceHandler(service *services.ServiceService) *ServiceHandler {
	return &ServiceHandler{
		service: service,
	}
}

// Create 创建服务
// @Summary 创建新服务
// @Description 创建一个新的理发服务
// @Tags 服务管理
// @Accept json
// @Produce json
// @Param service body models.ServiceRequest true "服务信息"
// @Security Bearer
// @Success 200 {object} models.Response{data=models.Service}
// @Failure 400 {object} models.Response
// @Router /admin/services [post]
func (h *ServiceHandler) Create(c *gin.Context) {
	var req models.ServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	service, err := h.service.CreateService(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(service))
}

// List 获取服务列表
// @Summary 获取服务列表
// @Description 获取所有可用的理发服务列表
// @Tags 服务管理
// @Produce json
// @Param active query bool false "是否只显示启用的服务"
// @Success 200 {object} models.Response{data=[]models.Service}
// @Router /services [get]
func (h *ServiceHandler) List(c *gin.Context) {
	activeOnly := c.Query("active") == "true"
	services, err := h.service.GetServices(c.Request.Context(), activeOnly)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(services))
}

// Update 更新服务
// @Summary 更新服务信息
// @Description 更新指定服务的信息
// @Tags 服务管理
// @Accept json
// @Produce json
// @Param id path int true "服务ID"
// @Param service body models.ServiceRequest true "服务信息"
// @Security Bearer
// @Success 200 {object} models.Response{data=models.Service}
// @Failure 400,404 {object} models.Response
// @Router /admin/services/{id} [put]
func (h *ServiceHandler) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, "Invalid service ID"))
		return
	}

	var req models.ServiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	service, err := h.service.UpdateService(c.Request.Context(), uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(service))
}

// Delete 删除服务
// @Summary 删除服务
// @Description 删除指定的服务
// @Tags 服务管理
// @Produce json
// @Param id path int true "服务ID"
// @Security Bearer
// @Success 200 {object} models.Response
// @Failure 400,404 {object} models.Response
// @Router /admin/services/{id} [delete]
func (h *ServiceHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, "Invalid service ID"))
		return
	}

	if err := h.service.DeleteService(c.Request.Context(), uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(nil))
}

// ToggleStatus 切换服务状态
// @Summary 切换服务状态
// @Description 启用或禁用指定的服务
// @Tags 服务管理
// @Produce json
// @Param id path int true "服务ID"
// @Security Bearer
// @Success 200 {object} models.Response{data=models.Service}
// @Failure 400,404 {object} models.Response
// @Router /admin/services/{id}/toggle [put]
func (h *ServiceHandler) ToggleStatus(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse(http.StatusBadRequest, "Invalid service ID"))
		return
	}

	service, err := h.service.ToggleServiceStatus(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse(http.StatusInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse(service))
} 