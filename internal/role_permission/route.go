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
			"/role/:roleID/permission/:permissionID",

			ctrl.DeleteRolePermission,
		)
		routeGroup.GET("/role/:id", ctrl.FindByRoleID)
		routeGroup.GET("/permission/:id", ctrl.FindByPermissionID)

	}
}
