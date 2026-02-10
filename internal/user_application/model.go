package user_application

import (
	"company_iam/internal/role"
	"company_iam/internal/user"
)

type UserApplication struct {
	ID            uint `gorm:"primaryKey;autoIncrement"`
	UserID        uint `gorm:"not null"`
	ApplicationID uint `gorm:"not null"`
	//relations
	User        user.User `gorm:"foreignKey:UserID"`
	Application role.Role `gorm:"foreignKey:ApplicationID"`
}

type UserApplicationRequest struct {
	UserID        uint `json:"user_id" binding:"required"`
	ApplicationID uint `json:"application_id" binding:"required"`
}

type UserApplicationResponse struct {
	ID            uint `json:"id"`
	UserID        uint `json:"user_id"`
	ApplicationID uint `json:"application_id"`
}
