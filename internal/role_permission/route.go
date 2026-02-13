package role_permission

import (
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/role-permissions")
	routeGroup.Use(middlewares.Authenticate(cfg)) // Add authentication middleware here
	{
		routeGroup.POST("/",middlewares.AuthorizePermission("iam.role-permission.create"), ctrl.CreateRolePermission)
		routeGroup.DELETE(
			"/role/:roleID/permission/:permissionID",
			middlewares.AuthorizePermission("iam.role-permission.delete"),
			ctrl.DeleteRolePermission,
		)
		routeGroup.GET("/role/:id",middlewares.AuthorizePermission("iam.role-permission.read"), ctrl.FindByRoleID)
		routeGroup.GET("/permission/:id",middlewares.AuthorizePermission("iam.role-permission.read"), ctrl.FindByPermissionID)

	}
}
