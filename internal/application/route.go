package application

import (
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	applicationGroup := r.Group("/api/applications")
	applicationGroup.Use(middlewares.Authenticate(cfg))
	{
		applicationGroup.GET("/",middlewares.AuthorizePermission("iam.application.read"), ctrl.GetAllApplications)
	}
}
