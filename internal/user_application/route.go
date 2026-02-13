package user_application

import (
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/user-applications")
	routeGroup.Use(middlewares.Authenticate(cfg))
	{
		routeGroup.POST("/",middlewares.AuthorizePermission("iam.user-application.create"), ctrl.Create)
		routeGroup.DELETE("/:id",middlewares.AuthorizePermission("iam.user-application.delete"), ctrl.Delete)
		routeGroup.GET("/user/:id",middlewares.AuthorizePermission("iam.user-application.read"), ctrl.GetByUserID)
		routeGroup.GET("/application/:id",middlewares.AuthorizePermission("iam.user-application.read"), ctrl.GetByApplicationID)

	}
}
