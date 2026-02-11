package user_role

import (
	"company_iam/internal/role"
	"company_iam/internal/user"
)

type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`

	User user.User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Role role.Role `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
}

type UserRoleRequest struct {
	UserID uint `json:"user_id" binding:"required"`
	RoleID uint `json:"role_id" binding:"required"`
}

type UserRoleResponse struct {
	UserID uint `json:"user_id"`
	RoleID uint `json:"role_id"`
}

type UserRoleWithRoleResponse struct {
	UserID uint                `json:"user_id"`
	Role   []role.RoleResponse `json:"role"`
}
