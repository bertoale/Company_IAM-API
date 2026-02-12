package user_application

func ToUserApplicationResponse(userApplication *UserApplication) *UserApplicationResponse {
	return &UserApplicationResponse{
		UserID:        userApplication.UserID,
		ApplicationID: userApplication.ApplicationID,
	}
}

func ToApplicationWithUsersResponse(applicationID uint, userApplications []UserApplication) *ApplicationWithUsersResponse {
	var users []SimpleUserResponse
	var appCode string
	var appName string
	if len(userApplications) > 0 {
		appCode = userApplications[0].Application.Code
		appName = userApplications[0].Application.Name
	}
	for _, ua := range userApplications {
		if ua.User.ID != 0 {
			users = append(users, SimpleUserResponse{
				ID:       ua.User.ID,
				Email:    ua.User.Email,
				Username: ua.User.Username,
			})
		}
	}
	return &ApplicationWithUsersResponse{
		ID:    applicationID,
		Code:  appCode,
		Name:  appName,
		Users: users,
	}
}

func ToUserWithApplicationsResponse(userID uint, userApplications []UserApplication) *UserWithApplicationsResponse {
	var applications []SimpleApplicationResponse
	var userEmail string
	var userName string
	if len(userApplications) > 0 {
		userEmail = userApplications[0].User.Email
		userName = userApplications[0].User.Username
	}
	for _, ua := range userApplications {
		if ua.Application.ID != 0 {
			applications = append(applications, SimpleApplicationResponse{
				ID:   ua.Application.ID,
				Code: ua.Application.Code,
				Name: ua.Application.Name,
			})
		}
	}
	return &UserWithApplicationsResponse{
		ID:           userID,
		Email:        userEmail,
		Username:     userName,
		Applications: applications,
	}
}	