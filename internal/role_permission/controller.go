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

func ParseID(c *gin.Context) (uint, error) {
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
		response.Error(c, http.StatusBadRequest, "Invalid request data")
		return
	}
	res, err := ctrl.service.CreateRolePermission(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "Role-Permission created successfully", res)
}

func (ctrl *Controller) DeleteRolePermission(c *gin.Context) {
	rolePermissionID, err := ParseID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid role-permission ID")
		return
	}
	if err := ctrl.service.DeleteRolePermission(rolePermissionID); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete Role-Permission")
		return
	}
	response.Success(c, http.StatusOK, "Role-Permission deleted successfully", nil)
}


func NewController(service Service) *Controller {
	return &Controller{service: service}
}