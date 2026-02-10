package auth

import (
	"company_iam/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func (ctrl *Controller) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	token, userRes, err := ctrl.service.Login(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Login failed: "+err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
		"user":  userRes,
	})
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}