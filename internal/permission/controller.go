package permission

import (
	"company_iam/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func (ctrl *Controller) GetAllPermissions(c *gin.Context) {
	res, err := ctrl.service.GetAllPermissions()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve permissions")
		return
	}
	response.Success(c, http.StatusOK, "Permissions retrieved successfully", res)
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}