package user

import (
	"company_iam/pkg/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine, ctrl *Controller, cfg *config.Config) {
	routeGroup := r.Group("/api/users")
	{
		routeGroup.POST("/", ctrl.CreateUser)
		routeGroup.GET("/", ctrl.GetAllUsers)
		routeGroup.GET("/:id", ctrl.GetUserByID)
		routeGroup.PUT("/:id", ctrl.UpdateUser)
		routeGroup.DELETE("/:id", ctrl.DeleteUser)

	}
}
