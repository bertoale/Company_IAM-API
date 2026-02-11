package role_permission

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
func ParsePermissionID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (ctrl *Controller) CreateRolePermission(c *gin.Context) {
	var req RolePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request")
		return
	}
	rolePermission, err := ctrl.service.CreateRolePermission(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create role permission")
		return
	}
	response.Success(c, http.StatusCreated, "Role permission created successfully", rolePermission)
}

func (ctrl *Controller) DeleteRolePermission(c *gin.Context) {
	roleID, err := ParseRoleID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role ID")
		return
	}
	permissionID, err := ParsePermissionID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid permission ID")
		return
	}
	err = ctrl.service.DeleteRolePermission(roleID, permissionID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete role permission")
		return
	}
	response.Success(c, http.StatusOK, "Role permission deleted successfully", nil)
}

func (ctrl *Controller) FindByRoleID(c *gin.Context) {
	roleID, err := ParseRoleID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role ID")
		return
	}
	rolePermissions, err := ctrl.service.FindByRoleID(roleID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get role permissions by role ID")
		return
	}
	response.Success(c, http.StatusOK, "Role permissions retrieved successfully", rolePermissions)
}

func (ctrl *Controller) FindByPermissionID(c *gin.Context) {
	permissionID, err := ParsePermissionID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid permission ID")
		return
	}
	rolePermissions, err := ctrl.service.FindByPermissionID(permissionID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get role permissions by permission ID")
		return
	}
	response.Success(c, http.StatusOK, "Role permissions retrieved successfully", rolePermissions)
}

func (ctrl *Controller) FindByRoleIDWithPermission(c *gin.Context) {
	roleID, err := ParseRoleID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role ID")
		return
	}
	rolePermissions, err := ctrl.service.FindByRoleIDWithPermission(roleID)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to get role permissions with permission by role ID")
		return
	}
	response.Success(c, http.StatusOK, "Role permissions with permissions retrieved successfully", rolePermissions)
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}
