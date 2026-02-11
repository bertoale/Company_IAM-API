package user_role

import "company_iam/internal/role"

func ToUserRoleResponse(userRole *UserRole) *UserRoleResponse {
	return &UserRoleResponse{
		UserID: userRole.UserID,
		RoleID: userRole.RoleID,
	}
}

func ToUserRolesResponse(userID uint, userRoles []UserRole) *UserRoleWithRoleResponse {
	roles := make([]role.RoleResponse, 0, len(userRoles))

	for _, ur := range userRoles {
		roles = append(roles, role.RoleResponse{
			ID:   ur.Role.ID,
			Name: ur.Role.Name,
		})
	}

	return &UserRoleWithRoleResponse{
		UserID: userID,
		Role:   roles,
	}
}
