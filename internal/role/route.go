package role

import (
	"company_iam/internal/rbac"
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config, rbacService *rbac.Service) {
	routeGroup := r.Group("/api/roles")
	routeGroup.Use(middlewares.Authenticate(cfg))
	{
		routeGroup.POST("/",middlewares.AuthorizePermission(rbacService, "iam.role.create"), ctrl.CreateRole)
		routeGroup.GET("/",middlewares.AuthorizePermission(rbacService, "iam.role.read"), ctrl.GetAllRoles)
		routeGroup.GET("/:id",middlewares.AuthorizePermission(rbacService, "iam.role.read"), ctrl.GetRoleByID)
		routeGroup.PUT("/:id",middlewares.AuthorizePermission(rbacService, "iam.role.update"), ctrl.UpdateRole)
		routeGroup.DELETE("/:id",middlewares.AuthorizePermission(rbacService, "iam.role.delete"), ctrl.DeleteRole)
	}
}
