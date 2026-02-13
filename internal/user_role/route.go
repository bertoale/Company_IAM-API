package user_role

import (
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/user-roles")
	routeGroup.Use(middlewares.Authenticate(cfg))
	{
		routeGroup.POST("/", ctrl.Create)
		routeGroup.DELETE(
			"/user/:userID/role/:roleID",
			middlewares.AuthorizePermission("iam.user-role.delete"),
			ctrl.Delete,
		)
		routeGroup.GET("/user/:id",
		middlewares.AuthorizePermission("iam.user-role.read"),
		ctrl.GetByUserID)
		routeGroup.GET("/role/:id",
		middlewares.AuthorizePermission("iam.user-role.read"),
		ctrl.GetByRoleID)
	}
}
