package user_application

import (
	"company_iam/internal/application"
	"company_iam/internal/user"
)

type UserApplication struct {
	UserID        uint `gorm:"primaryKey"`
	ApplicationID uint `gorm:"primaryKey"`
	//relations
	User        user.User              `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Application application.Application `gorm:"foreignKey:ApplicationID;constraint:OnDelete:CASCADE"`
}

type UserApplicationRequest struct {
	UserID        uint `json:"user_id" binding:"required"`
	ApplicationID uint `json:"application_id" binding:"required"`
}

type UserApplicationResponse struct {
	UserID        uint `json:"user_id"`
	ApplicationID uint `json:"application_id"`
}

type SimpleUserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Username string `json:"username"`
}

type SimpleApplicationResponse struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type ApplicationWithUsersResponse struct {
	ID    uint                `json:"id"`
	Code  string              `json:"code"`
	Name  string              `json:"name"`
	Users []SimpleUserResponse `json:"users"`
}

type UserWithApplicationsResponse struct {
	ID        uint                      `json:"id"`
	Email     string                    `json:"email"`
	Username  string                    `json:"username"`
	Applications []SimpleApplicationResponse `json:"applications"`
}