package user_role

func ToUserRoleResponse(userRole *UserRole) *UserRoleResponse {
	return &UserRoleResponse{
		UserID: userRole.UserID,
		RoleID: userRole.RoleID,
	}
}

func ToRoleWithUsersResponse(roleID uint, userRoles []UserRole) *RoleWithUsersResponse {
	var users []SimpleUserResponse
	var roleName string
	if len(userRoles) > 0 {
		roleName = userRoles[0].Role.Name
	}
	for _, ur := range userRoles {
		if ur.User.ID != 0 {
			users = append(users, SimpleUserResponse{
				ID:    ur.User.ID,
				Email: ur.User.Email,
				Username: ur.User.Username,
			})
		}
	}
	return &RoleWithUsersResponse{
		ID:    roleID,
		Name:  roleName,
		Users: users,
	}
}

func ToUserWithRolesResponse(userID uint, userRoles []UserRole) *UserWithRolesResponse {
	var roles []SimpleRoleResponse
	var userEmail string
	var userName string
	if len(userRoles) > 0 {
		userEmail = userRoles[0].User.Email
		userName = userRoles[0].User.Username
	}
	for _, ur := range userRoles {
		if ur.Role.ID != 0 {
			roles = append(roles, SimpleRoleResponse{
				ID:   ur.Role.ID,
				Name: ur.Role.Name,
			})
		}
	}
	return &UserWithRolesResponse{
		ID:    userID,
		Email: userEmail,
		Username:  userName,
		Roles: roles,
	}
}
