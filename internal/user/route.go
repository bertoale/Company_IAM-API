package user

import (
	"company_iam/internal/rbac"
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config, rbacService *rbac.Service) {	routeGroup := r.Group("/api/users")
	routeGroup.Use(middlewares.Authenticate(cfg))
	{
		routeGroup.GET("/me", ctrl.GetCurrentUser) // ⚠️ harus di atas /:id
		routeGroup.POST("/", middlewares.AuthorizePermission(rbacService, "iam.user.create"), ctrl.CreateUser)
		routeGroup.GET("/", middlewares.AuthorizePermission(rbacService, "iam.user.read"), ctrl.GetAllUsers)
		routeGroup.GET("/:id", middlewares.AuthorizePermission(rbacService, "iam.user.read"), ctrl.GetUserByID)
		routeGroup.PUT("/:id", middlewares.AuthorizePermission(rbacService, "iam.user.update"), ctrl.UpdateUser)
		routeGroup.DELETE("/:id", middlewares.AuthorizePermission(rbacService, "iam.user.delete"), ctrl.DeleteUser)
	}
}
