package role_permission

import (
	"company_iam/internal/rbac"
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config, rbacService *rbac.Service) {
	routeGroup := r.Group("/api/role-permissions")
	routeGroup.Use(middlewares.Authenticate(cfg))
	{
		routeGroup.POST("/",middlewares.AuthorizePermission(rbacService, "iam.role-permission.create"), ctrl.CreateRolePermission)
		routeGroup.DELETE(
			"/role/:roleID/permission/:permissionID",
			middlewares.AuthorizePermission(rbacService, "iam.role-permission.delete"),
			ctrl.DeleteRolePermission,
		)
		routeGroup.GET("/role/:roleID",middlewares.AuthorizePermission(rbacService, "iam.role-permission.read"), ctrl.FindByRoleID)
		routeGroup.GET("/permission/:permissionID",middlewares.AuthorizePermission(rbacService, "iam.role-permission.read"), ctrl.FindByPermissionID)
	}
}
