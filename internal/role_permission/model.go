package role_permission

import (
	"company_iam/internal/permission"
	"company_iam/internal/role"
)

type RolePermission struct {
	RoleID       uint `gorm:"primaryKey"`
	PermissionID uint `gorm:"primaryKey"`

	Role       role.Role             `gorm:"foreignKey:RoleID;constraint:OnDelete:CASCADE"`
	Permission permission.Permission `gorm:"foreignKey:PermissionID;constraint:OnDelete:CASCADE"`
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

type RolePermissionWithPermissionResponse struct {
	RoleID     uint                 `json:"role_id"`
	Permission []permission.PermissionResponse `json:"permission"`
}