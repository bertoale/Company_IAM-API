package role

import (
	"company_iam/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/roles")
	{
		routeGroup.POST("/", ctrl.CreateRole)
		routeGroup.GET("/", ctrl.GetAllRoles)
		routeGroup.GET("/:id", ctrl.GetRoleByID)
		routeGroup.PUT("/:id", ctrl.UpdateRole)
		routeGroup.DELETE("/:id", ctrl.DeleteRole)
	}
}
