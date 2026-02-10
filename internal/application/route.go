package application

import (
	"company_iam/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	applicationGroup := r.Group("/api/applications")
	{
		applicationGroup.GET("/", ctrl.GetAllApplications)

	}
}
