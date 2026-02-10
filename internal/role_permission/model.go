package role_permission

import (
	"company_iam/internal/permission"
	"company_iam/internal/role"
)

type RolePermission struct {
	ID           uint `gorm:"primaryKey;autoIncrement"`
	RoleID       uint `gorm:"not null"`
	PermissionID uint `gorm:"not null"`
	//relations
	Role       role.Role             `gorm:"foreignKey:RoleID"`
	Permission permission.Permission `gorm:"foreignKey:PermissionID"`
}

type RolePermissionResponse struct {
	ID           uint `json:"id"`
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
}

type RolePermissionRequest struct {
	RoleID       uint `json:"role_id" binding:"required"`
	PermissionID uint `json:"permission_id" binding:"required"`
}
