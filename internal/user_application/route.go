package user_application

import (
	"company_iam/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/user-applications")
	{
		routeGroup.POST("/", ctrl.Create)
		routeGroup.DELETE("/:id", ctrl.Delete)
		routeGroup.GET("/user/:id", ctrl.GetByUserID)
		routeGroup.GET("/application/:id", ctrl.GetByApplicationID)

	}
}
