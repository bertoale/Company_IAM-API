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
	token, refreshToken, userRes, err := ctrl.service.Login(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Login failed: "+err.Error())
		return
	}
	
	// Set access token cookie (short expiry)
	c.SetCookie(
		"token",
		token,
		int((9 * time.Hour).Seconds()),
		"/",
		"",
		ctrl.config.NodeEnv == "production",
		true,
	)
	
	// Set refresh token cookie (long expiry)
	c.SetCookie(
		"refresh_token",
		refreshToken,
		int((7 * 24 * time.Hour).Seconds()),
		"/",
		"",
		ctrl.config.NodeEnv == "production",
		true,
	)
	
	response.Success(c, http.StatusOK, "Login successful", gin.H{
		"token":         token,
		"refresh_token": refreshToken,
		"user":          userRes,
	})
}

func (ctrl *Controller) RefreshToken(c *gin.Context) {
	// Read refresh_token from HttpOnly cookie first; fall back to JSON body
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		// Try JSON body as fallback
		var req RefreshTokenRequest
		if bindErr := c.ShouldBindJSON(&req); bindErr != nil || req.RefreshToken == "" {
			response.Error(c, http.StatusUnauthorized, "Refresh token not found")
			return
		}
		refreshToken = req.RefreshToken
	}

	newToken, newRefreshToken, err := ctrl.service.RefreshToken(refreshToken)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, "Invalid or expired refresh token: "+err.Error())
		return
	}

	// Set new access token cookie
	c.SetCookie(
		"token",
		newToken,
		int((8 * time.Hour).Seconds()),
		"/",
		"",
		ctrl.config.NodeEnv == "production",
		true,
	)
	
	// Set new refresh token cookie
	c.SetCookie(
		"refresh_token",
		newRefreshToken,
		int((7 * 24 * time.Hour).Seconds()),
		"/",
		"",
		ctrl.config.NodeEnv == "production",
		true,
	)

	response.Success(c, http.StatusOK, "Token refreshed successfully", gin.H{
		"token":         newToken,
		"refresh_token": newRefreshToken,
	})
}

func NewController(service Service, cfg *config.Config) *Controller {
	return &Controller{
		service: service,
		config:  *cfg,
	}
}
