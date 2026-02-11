package user_role

import "fmt"

type Service interface {
	Create(req *UserRoleRequest) (*UserRoleResponse, error)
	Delete(userID, roleID uint) error
	GetByUserID(userID uint) ([]UserRoleResponse, error)
	GetByRoleID(roleID uint) ([]UserRoleResponse, error)
	GetByUserIDWithRole(userID uint) (*UserRoleWithRoleResponse, error)
}

type service struct {
	repo Repository
}

// Create implements [Service].
func (s *service) Create(req *UserRoleRequest) (*UserRoleResponse, error) {
	// 1️⃣ Cek apakah sudah ada
	exists, err := s.repo.Exists(req.UserID, req.RoleID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user role: %w", err)
	}

	if exists {
		return nil, fmt.Errorf("role already assigned to user")
	}

	// 2️⃣ Create kalau belum ada
	userRole := &UserRole{
		UserID: req.UserID,
		RoleID: req.RoleID,
	}

	if err := s.repo.Create(userRole); err != nil {
		return nil, fmt.Errorf("failed to create user role: %w", err)
	}

	return ToUserRoleResponse(userRole), nil
}

// Delete implements [Service].
func (s *service) Delete(userID, roleID uint) error {
	return s.repo.Delete(userID, roleID)
}

// GetByRoleID implements [Service].
func (s *service) GetByRoleID(roleID uint) ([]UserRoleResponse, error) {
	var responses []UserRoleResponse
	userRoles, err := s.repo.FindByRoleID(roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles by role ID: %w", err)
	}
	for _, ur := range userRoles {
		responses = append(responses, *ToUserRoleResponse(&ur))
	}
	return responses, nil
}

// GetByUserID implements [Service].
func (s *service) GetByUserID(userID uint) ([]UserRoleResponse, error) {
	var responses []UserRoleResponse
	userRoles, err := s.repo.FindByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles by user ID: %w", err)
	}
	for _, ur := range userRoles {
		responses = append(responses, *ToUserRoleResponse(&ur))
	}
	return responses, nil
}

// GetByUserIDWithRole implements [Service].
func (s *service) GetByUserIDWithRole(userID uint) (*UserRoleWithRoleResponse, error) {
	userRoles, err := s.repo.FindByUserIDWithRole(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles with role by user ID: %w", err)
	}
	return ToUserRolesResponse(userID, userRoles), nil
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
