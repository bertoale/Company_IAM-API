package auth

type LoginRequest struct {
	Identifier string `json:"identifier"` // email atau username
	Password   string `json:"password"`
}

type LoginResponse struct {
	ID         uint   `json:"id"`
	Identifier string `json:"email"`
	Token      string `json:"token"`
}
