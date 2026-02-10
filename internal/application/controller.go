package application

import (
	"company_iam/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func (ctrl *Controller) GetAllApplications(c *gin.Context) {
	res, err := ctrl.service.GetAllApplications()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve applications")
		return
	}
	response.Success(c, http.StatusOK, "Applications retrieved successfully", res)
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}