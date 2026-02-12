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

type RolePermissionRequest struct {
	RoleID       uint `json:"role_id" binding:"required"`
	PermissionID uint `json:"permission_id" binding:"required"`
}

type RolePermissionResponse struct {
	RoleID       uint `json:"role_id"`
	PermissionID uint `json:"permission_id"`
}

type SimpleRoleResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type SimplePermissionResponse struct {
	ID   uint   `json:"id"`
	Code string `json:"code"`
	Description string `json:"description"`
}

type PermissionWithRolesResponse struct {
	ID          uint                        `json:"id"`
	Code        string                      `json:"code"`
	Description string                      `json:"description"`
	Roles       []SimpleRoleResponse         `json:"roles"`
}

type RoleWithPermissionsResponse struct {
	ID          uint                        `json:"id"`
	Name        string                      `json:"name"`
	Description string                      `json:"description"`
	Permissions  []SimplePermissionResponse  `json:"permissions"`
}