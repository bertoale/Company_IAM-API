package user

import (
	"company_iam/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service Service
}

func ParseUserID(c *gin.Context) (uint, error) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func (ctrl *Controller) CreateUser(c *gin.Context) {
	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	res, err := ctrl.service.CreateUser(&req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to create user: "+err.Error())
		return
	}
	response.Success(c, http.StatusCreated, "User created successfully", res)
}

func (ctrl *Controller) GetUserByID(c *gin.Context) {
	id, err := ParseUserID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	res, err := ctrl.service.GetUserByID(id)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve user: "+err.Error())
		return
	}
	response.Success(c, http.StatusOK, "User retrieved successfully", res)
}

func (ctrl *Controller) GetAllUsers(c *gin.Context) {
	res, err := ctrl.service.GetAllUsers()
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to retrieve users: "+err.Error())
		return
	}
	response.Success(c, http.StatusOK, "Users retrieved successfully", res)
}
func (ctrl *Controller) UpdateUser(c *gin.Context) {
	id, err := ParseUserID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	var req UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request payload")
		return
	}
	res, err := ctrl.service.UpdateUser(id, &req)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to update user: "+err.Error())
		return
	}
	response.Success(c, http.StatusOK, "User updated successfully", res)

}

func (ctrl *Controller) DeleteUser(c *gin.Context) {
	id, err := ParseUserID(c)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID")
		return
	}
	if err := ctrl.service.DeleteUser(id); err != nil {
		response.Error(c, http.StatusInternalServerError, "Failed to delete user: "+err.Error())
		return
	}
	response.Success(c, http.StatusOK, "User deleted successfully", nil)
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}