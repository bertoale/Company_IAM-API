package user_role

import (
	"company_iam/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/user-roles")
	{
		routeGroup.POST("/", ctrl.Create)
		routeGroup.DELETE("/:id", ctrl.Delete)

	}
}
