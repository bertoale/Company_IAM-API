package auth

import (
	"company_iam/pkg/config"
	"company_iam/pkg/response"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	config  config.Config
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
	c.SetCookie(
		"token",
		token,
		int((7 * 24 * time.Hour).Seconds()),
		"/",
		"",
		ctrl.config.NodeEnv == "production",
		true,
	)
	response.Success(c, http.StatusOK, "Login successful", gin.H{
		"token": token,
		"user":  userRes,
	})
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}
