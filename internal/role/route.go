package role

import (
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/roles")
	routeGroup.Use(middlewares.Authenticate(cfg))
	{
		routeGroup.POST("/",middlewares.AuthorizePermission("iam.role.create"), ctrl.CreateRole)
		routeGroup.GET("/",middlewares.AuthorizePermission("iam.role.read"), ctrl.GetAllRoles)
		routeGroup.GET("/:id",middlewares.AuthorizePermission("iam.role.read"), ctrl.GetRoleByID)
		routeGroup.PUT("/:id",middlewares.AuthorizePermission("iam.role.update"), ctrl.UpdateRole)
		routeGroup.DELETE("/:id",middlewares.AuthorizePermission("iam.role.delete"), ctrl.DeleteRole)
	}
}
