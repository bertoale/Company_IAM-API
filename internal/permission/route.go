package permission

import (
	"company_iam/internal/rbac"
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	ctrl *Controller,
	cfg *config.Config,
	rbacService *rbac.Service,
) {
	permissionGroup := r.Group("/api/permissions")
	permissionGroup.Use(middlewares.Authenticate(cfg))
	{
		permissionGroup.GET("/",
			middlewares.AuthorizePermission(rbacService, "iam.permission.read"),
			ctrl.GetAllPermissions,
		)


	}
}