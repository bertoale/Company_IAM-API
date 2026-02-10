package user

func ToUserResponse(user *User) *UserResponse {
	return &UserResponse{
		ID:         user.ID,
		Email:      user.Email,
		Name:       user.Username,
		StatusEnum: user.Status,
	}
}
