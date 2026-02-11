package role_permission

import (
	"fmt"
)

type Service interface {
	CreateRolePermission(req *RolePermissionRequest) (*RolePermissionResponse, error)
	DeleteRolePermission(roleID uint, permissionID uint) error
	FindByRoleID(roleID uint) ([]RolePermissionResponse, error)
	FindByPermissionID(permissionID uint) ([]RolePermissionResponse, error)
	FindByRoleIDWithPermission(roleID uint) (*RolePermissionWithPermissionResponse, error)
}

type service struct {
	repo Repository
}

// CreateRolePermission implements [Service].
func (s *service) CreateRolePermission(req *RolePermissionRequest) (*RolePermissionResponse, error) {
	exists, err := s.repo.Exists(req.RoleID, req.PermissionID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing role permission: %w", err)
	}

	if exists {
		return nil, fmt.Errorf("permission already assigned to role")
	}
	rolePermission := &RolePermission{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
	}
	if err := s.repo.Create(rolePermission); err != nil {
		return nil, fmt.Errorf("failed to create role permission: %w", err)
	}
	return ToRolePermissionResponse(rolePermission), nil
}

// DeleteRolePermission implements [Service].
func (s *service) DeleteRolePermission(roleID uint, permissionID uint) error {
	return s.repo.Delete(roleID, permissionID)
}

// FindByPermissionID implements [Service].
func (s *service) FindByPermissionID(permissionID uint) ([]RolePermissionResponse, error) {
	var responses []RolePermissionResponse
	rolePermissions, err := s.repo.FindByPermissionID(permissionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role permissions by permission ID: %w", err)
	}
	for _, rp := range rolePermissions {
		responses = append(responses, *ToRolePermissionResponse(&rp))
	}
	return responses, nil
}

// FindByRoleID implements [Service].
func (s *service) FindByRoleID(roleID uint) ([]RolePermissionResponse, error) {
	var responses []RolePermissionResponse
	rolePermissions, err := s.repo.FindByRoleID(roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role permissions by role ID: %w", err)
	}
	for _, rp := range rolePermissions {
		responses = append(responses, *ToRolePermissionResponse(&rp))
	}
	return responses, nil
}

// FindByRoleIDWithPermission implements [Service].
func (s *service) FindByRoleIDWithPermission(roleID uint) (*RolePermissionWithPermissionResponse, error) {
	rolePermissions, err := s.repo.FindByRoleIDWithPermission(roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role permissions with permission by role ID: %w", err)
	}
	return ToRolePermissionWithPermissionResponse(roleID, rolePermissions), nil
}


func NewService(repo Repository) Service {
	return &service{repo: repo}
}
