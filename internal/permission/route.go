package permission

import (
	"company_iam/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	permissionGroup := r.Group("/api/permissions")
	{
		permissionGroup.GET("/", ctrl.GetAllPermissions)

	}
}
