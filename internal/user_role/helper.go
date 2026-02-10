package user_role

func ToUserRoleResponse(userRole *UserRole) *UserRoleResponse {
	return &UserRoleResponse{
		ID:     userRole.ID,
		UserID: userRole.UserID,
		RoleID: userRole.RoleID,
	}
}