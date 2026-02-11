package role

import (
	"company_iam/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func ParseRoleID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (ctrl *Controller) CreateRole(c *gin.Context) {
	var req RoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	res, err := ctrl.service.CreateRole(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create role: ")
		return
	}
	response.Success(c, http.StatusCreated, "Role created successfully", res)
}

func (ctrl *Controller) GetRoleByID(c *gin.Context) {
	id, err := ParseRoleID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role ID")
		return
	}
	res, err := ctrl.service.GetRoleByID(id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve role")
		return
	}
	response.Success(c, http.StatusOK, "Role retrieved successfully", res)
}

func (ctrl *Controller) GetAllRoles(c *gin.Context) {
	res, err := ctrl.service.GetAllRoles()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve roles")
		return
	}
	response.Success(c, http.StatusOK, "Roles retrieved successfully", res)
}

func (ctrl *Controller) UpdateRole(c *gin.Context) {
	id, err := ParseRoleID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role ID")
		return
	}
	var req UpdateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	res, err := ctrl.service.UpdateRole(id, &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update role")
		return
	}
	response.Success(c, http.StatusOK, "Role updated successfully", res)
}

func (ctrl *Controller) DeleteRole(c *gin.Context) {
	id, err := ParseRoleID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role ID")
		return
	}
	if err := ctrl.service.DeleteRole(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete role")
		return
	}
	response.Success(c, http.StatusOK, "Role deleted successfully", nil)
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}
