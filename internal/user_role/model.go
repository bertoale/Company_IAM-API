package user_role

import (
	"company_iam/internal/role"
	"company_iam/internal/user"
)

type UserRole struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID uint `gorm:"not null"`
	RoleID uint `gorm:"not null"`
	//relations
	User user.User `gorm:"foreignKey:UserID"`
	Role role.Role `gorm:"foreignKey:RoleID"`
}

type UserRoleRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	RoleID uint `json:"role_id" binding:"required"`
}

type UserRoleResponse struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}