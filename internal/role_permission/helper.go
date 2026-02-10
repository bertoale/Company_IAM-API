package role_permission

func ToRolePermissionResponse(rp *RolePermission) *RolePermissionResponse {
	return &RolePermissionResponse{
		ID:           rp.ID,
		RoleID:       rp.RoleID,
		PermissionID: rp.PermissionID,
	}
}
