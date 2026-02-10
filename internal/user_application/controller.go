package user_application

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
	var req UserApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, 400, "Invalid request data")
		return
	}
	if err := ctrl.service.Create(&req); err != nil {
		response.Error(c, 500, "Failed to create User-Application: ")
		return
	}
	response.Success(c, 201, "User-Application created successfully", nil)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	userApplicationID, err := ParseID(c)
	if err != nil {
		response.Error(c, 400, "Invalid User-Application ID")
		return
	}
	if err := ctrl.service.Delete(userApplicationID); err != nil {
		response.Error(c, 500, "Failed to delete User-Application")
		return
	}
	response.Success(c, 200, "User-Application deleted successfully", nil)
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}