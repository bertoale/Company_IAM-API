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
	res, err := ctrl.service.Create(&req)
	if err != nil {
		response.Error(c, 500, "Failed to create User-Application: "+err.Error())
		return
	}
	response.Success(c, 201, "User-Application created successfully", res)
}

func (ctrl *Controller) Delete(c *gin.Context) {
	userIDParam := c.Param("userID")
	userID, err := strconv.ParseUint(userIDParam, 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid userID parameter")
		return
	}

	appIDParam := c.Param("applicationID")
	applicationID, err := strconv.ParseUint(appIDParam, 10, 32)
	if err != nil {
		response.Error(c, 400, "Invalid applicationID parameter")
		return
	}

	err = ctrl.service.Delete(uint(userID), uint(applicationID))
	if err != nil {
		response.Error(c, 500, "Failed to delete User-Application: "+err.Error())
		return
	}
	response.Success(c, 200, "User-Application deleted successfully", nil)
}

func (ctrl *Controller) GetByUserID(c *gin.Context) {
	userID, err := ParseID(c)
	if err != nil {
		response.Error(c, 400, "Invalid userID parameter")
		return
	}
	
	res, err := ctrl.service.GetByUserID(userID)	
	if err != nil {
		response.Error(c, 500, "Failed to get User-Applications by User ID: "+err.Error())
		return
	}
	response.Success(c, 200, "User-Applications retrieved successfully", res)
}

func (ctrl *Controller) GetByApplicationID(c *gin.Context) {
	applicationID, err := ParseID(c)
	if err != nil {
		response.Error(c, 400, "Invalid applicationID parameter")
		return
	}
	
	res, err := ctrl.service.GetByApplicationID(applicationID)	
	if err != nil {
		response.Error(c, 500, "Failed to get User-Applications by Application ID: "+err.Error())
		return
	}
	response.Success(c, 200, "User-Applications retrieved successfully", res)
}


func NewController(service Service) *Controller {
	return &Controller{service: service}
}
