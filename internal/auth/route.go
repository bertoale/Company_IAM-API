package auth

import (
	"company_iam/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	authGroup := r.Group("/api/auth")
	{
		authGroup.POST("/login", ctrl.Login)
		authGroup.POST("/refresh-token", ctrl.RefreshToken)
	}
}
