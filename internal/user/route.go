package user

import (
	"company_iam/pkg/config"
	"company_iam/pkg/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/users")
	routeGroup.Use(middlewares.Authenticate(cfg))
	{
		routeGroup.POST("/", middlewares.AuthorizePermission("iam.user.create"), ctrl.CreateUser)
		routeGroup.GET("/", middlewares.AuthorizePermission("iam.user.read"), ctrl.GetAllUsers)
		routeGroup.GET("/:id", middlewares.AuthorizePermission("iam.user.read"), ctrl.GetUserByID)
		routeGroup.PUT("/:id", middlewares.AuthorizePermission("iam.user.update"), ctrl.UpdateUser)
		routeGroup.DELETE("/:id",middlewares.AuthorizePermission("iam.user.delete"), ctrl.DeleteUser)
	}
}
