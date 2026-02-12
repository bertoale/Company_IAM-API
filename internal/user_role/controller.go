package user_role

import (
	"company_iam/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func ParseUserID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
func ParseRoleID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (ctrl *Controller) Create(c *gin.Context) {
	var req UserRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request data")
		return
	}
	res, err := ctrl.service.Create(&req)
	if err != nil {
		response.Error(c, 500, "Failed to create User-Role: "+err.Error())
		return
	}
	response.Success(c, 201, "User-Role created successfully", res)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	userID, err := ParseUserID(c)
	if err != nil {
		response.Error(c, 400, "Invalid userID parameter")
		return
	}

	roleID, err := ParseRoleID(c)
	if err != nil {
		response.Error(c, 400, "Invalid roleID parameter")
		return
	}

	err = ctrl.service.Delete(userID, roleID)
	if err != nil {
		response.Error(c, 500, "Failed to delete User-Role: "+err.Error())
		return
	}

	response.Success(c, 200, "User-Role deleted successfully", nil)
}

func (ctrl *Controller) GetByUserID(c *gin.Context) {
	id, err := ParseUserID(c)
	if err != nil {
		response.Error(c, 400, "Invalid ID parameter")
		return
	}
	res, err := ctrl.service.GetByUserID(id)
	if err != nil {
		response.Error(c, 500, "Failed to get User-Roles by User ID: "+err.Error())
		return
	}
	response.Success(c, 200, "User-Roles retrieved successfully", res)
}

func (ctrl *Controller) GetByRoleID(c *gin.Context) {
	id, err := ParseRoleID(c)
	if err != nil {
		response.Error(c, 400, "Invalid ID parameter")
		return
	}
	res, err := ctrl.service.GetByRoleID(id)
	if err != nil {
		response.Error(c, 500, "Failed to get User-Roles by Role ID: "+err.Error())
		return
	}
	response.Success(c, 200, "User-Roles retrieved successfully", res)
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}
