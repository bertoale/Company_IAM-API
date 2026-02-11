package role_permission

import (
	"company_iam/internal/permission"
)

func ToRolePermissionResponse(rp *RolePermission) *RolePermissionResponse {
	return &RolePermissionResponse{
		RoleID:       rp.RoleID,
		PermissionID: rp.PermissionID,
	}
}

func ToRolePermissionWithPermissionResponse(roleID uint,rolePermissions []RolePermission) *RolePermissionWithPermissionResponse {
	permissions := make([]permission.PermissionResponse, 0, len(rolePermissions))

	for _, r := range rolePermissions {
		permissions = append(permissions, permission.PermissionResponse{
			ID:   r.Permission.ID,
			Code: r.Permission.Code,
		})
	}
	
	return &RolePermissionWithPermissionResponse{
		RoleID:     roleID,
		Permission: permissions,
	}
}
