package user_application

func ToUserApplicationResponse(userApplication *UserApplication) *UserApplicationResponse {
	return &UserApplicationResponse{
		ID:            userApplication.ID,
		UserID:        userApplication.UserID,
		ApplicationID: userApplication.ApplicationID,
	}
}
