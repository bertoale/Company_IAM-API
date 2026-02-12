package role_permission

func ToRolePermissionResponse(rp *RolePermission) *RolePermissionResponse {
	return &RolePermissionResponse{
		RoleID:       rp.RoleID,
		PermissionID: rp.PermissionID,
	}
}

func ToPermissionWithRolesResponse(permissionID uint, rolePermissions []RolePermission) *PermissionWithRolesResponse {
	var roles []SimpleRoleResponse
	var permissionCode string
	var permissionDescription string
	if len(rolePermissions) > 0 {
		permissionCode = rolePermissions[0].Permission.Code
		permissionDescription = rolePermissions[0].Permission.Description
	}
	for _, rp := range rolePermissions {
		if rp.Role.ID != 0 {
			roles = append(roles, SimpleRoleResponse{
				ID: rp.Role.ID,
				Name: rp.Role.Name,
				Description: rp.Role.Description,
			})
		}
	}
	return &PermissionWithRolesResponse{
		ID:          permissionID,
		Code:        permissionCode,
		Description: permissionDescription,
		Roles:       roles,
	}
}

func ToRoleWithPermissionsResponse(roleID uint, rolePermissions []RolePermission) *RoleWithPermissionsResponse {
	var permissions []SimplePermissionResponse
	var roleName string
	var roleDescription string
	if len(rolePermissions) > 0 {
		roleName = rolePermissions[0].Role.Name
		roleDescription = rolePermissions[0].Role.Description
	}
	for _, rp := range rolePermissions {
		if rp.Permission.ID != 0 {
			permissions = append(permissions, SimplePermissionResponse{
				ID: rp.Permission.ID,
				Code: rp.Permission.Code,
				Description: rp.Permission.Description,
			})
		}
	}
	return &RoleWithPermissionsResponse{
		ID:          roleID,
		Name:        roleName,
		Description: roleDescription,
		Permissions:  permissions,
	}
}
