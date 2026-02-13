package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	CreateUser(req *UserRequest) (*UserResponse, error)
	GetUserByID(id uint) (*UserResponse, error)
	GetAllUsers() ([]UserResponse, error)
	UpdateUser(id uint, req *UserUpdateRequest) (*UserResponse, error)
	DeleteUser(id uint) error
}

type service struct {
	repo Repository
}

// CreateUser implements [Service].
func (s *service) CreateUser(req *UserRequest) (*UserResponse, error) {
	var existingEmail, _ = s.repo.FindByEmail(req.Email)
	if existingEmail != nil {
		return nil, fmt.Errorf("user with email '%s' already exists", req.Email)
	}
	var existingUsername, _ = s.repo.FindByUsername(req.Username)
	if existingUsername != nil {
		return nil, fmt.Errorf("user with username '%s' already exists", req.Username)
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	newUser := &User{
		Email:    req.Email,
		Username: req.Username,
		Password: string(hashedPassword),
	}
	if err := s.repo.Create(newUser); err != nil {
		return nil, err
	}
	return ToUserResponse(newUser), nil
}

// DeleteUser implements [Service].
func (s *service) DeleteUser(id uint) error {
	rowsAffected, err := s.repo.Delete(&User{ID: id})
	if err != nil {
		return fmt.Errorf("failed to delete user with id '%d': %w", id, err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user with id '%d' not found", id)
	}
	return nil
}

// GetAllUsers implements [Service].
func (s *service) GetAllUsers() ([]UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var responses []UserResponse
	for _, user := range users {
		responses = append(responses, *ToUserResponse(&user))
	}
	return responses, nil
}

// GetUserByID implements [Service].
func (s *service) GetUserByID(id uint) (*UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("user with id '%d' not found", id)
	}
	return ToUserResponse(user), nil
}

// UpdateUser implements [Service].
func (s *service) UpdateUser(id uint, req *UserUpdateRequest) (*UserResponse, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("user with id '%d' not found", id)
	}
	if req.Email != nil {
		existingEmail, _ := s.repo.FindByEmail(*req.Email)
		if existingEmail != nil && existingEmail.ID != id {
			return nil, fmt.Errorf("user with email '%s' already exists", *req.Email)
		}
		user.Email = *req.Email
	}
	if req.Username != nil {
		existingUsername, _ := s.repo.FindByUsername(*req.Username)
		if existingUsername != nil && existingUsername.ID != id {
			return nil, fmt.Errorf("user with username '%s' already exists", *req.Username)
		}
		user.Username = *req.Username
	}
	if req.Password != nil {
		user.Password = *req.Password
	}
	if req.Status != nil {
		user.Status = StatusEnum(*req.Status)
	}
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}
	return ToUserResponse(user), nil
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
