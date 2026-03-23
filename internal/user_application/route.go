package user_application

import (
	"company_iam/internal/rbac"
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config, rbacService *rbac.Service) {
	routeGroup := r.Group("/api/user-applications")
	routeGroup.Use(middlewares.Authenticate(cfg))
	{
		routeGroup.POST("/", middlewares.AuthorizePermission(rbacService, "iam.user-application.create"), ctrl.Create)
		routeGroup.DELETE("/user/:userID/application/:applicationID", middlewares.AuthorizePermission(rbacService, "iam.user-application.delete"), ctrl.Delete)
		routeGroup.GET("/user/:id", middlewares.AuthorizePermission(rbacService, "iam.user-application.read"), ctrl.GetByUserID)
		routeGroup.GET("/application/:id", middlewares.AuthorizePermission(rbacService, "iam.user-application.read"), ctrl.GetByApplicationID)
	}
}
