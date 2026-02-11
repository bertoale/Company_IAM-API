package role_permission

import (
	"company_iam/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/role-permissions")
	{
		routeGroup.POST("/", ctrl.CreateRolePermission)
		routeGroup.DELETE(
			"/roles/:roleID/permissions/:permissionID",

			ctrl.DeleteRolePermission,
		)
		routeGroup.GET("/roles/:id", ctrl.FindByRoleID)
		routeGroup.GET("/permissions/:id", ctrl.FindByPermissionID)
		routeGroup.GET("/roles/:id/permissions", ctrl.FindByRoleIDWithPermission)

	}
}
