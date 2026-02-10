package user_role

import (
	"company_iam/pkg/response"
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
	userRoleID, err := ParseID(c)
	if err != nil {
		response.Error(c, 400, "Invalid User-Role ID")
		return
	}
	if err := ctrl.service.Delete(userRoleID); err != nil {
		response.Error(c, 500, "Failed to delete User-Role")
		return
	}
	response.Success(c, 200, "User-Role deleted successfully", nil)
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}