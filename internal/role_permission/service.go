package role_permission

import (
	"fmt"
)

type Service interface {
	CreateRolePermission(req *RolePermissionRequest) (*RolePermissionResponse, error)
	DeleteRolePermission(roleID uint, permissionID uint) error
	FindByRoleID(roleID uint) (*RoleWithPermissionsResponse, error)
	FindByPermissionID(permissionID uint) (*PermissionWithRolesResponse, error)
}

type service struct {
	repo Repository
}

// FindByPermissionID implements [Service].
func (s *service) FindByPermissionID(permissionID uint) (*PermissionWithRolesResponse, error) {
	rolePermission, err := s.repo.FindByPermissionID(permissionID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role permissions by permission ID: %w", err)
	}
	return ToPermissionWithRolesResponse(permissionID, rolePermission), nil
}

// FindByRoleID implements [Service].
func (s *service) FindByRoleID(roleID uint) (*RoleWithPermissionsResponse, error) {
	rolePermission, err := s.repo.FindByRoleID(roleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role permissions by role ID: %w", err)
	}
	return ToRoleWithPermissionsResponse(roleID, rolePermission), nil
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
	RowsAffected, err := s.repo.Delete(roleID, permissionID)
	if err != nil {
		return fmt.Errorf("failed to delete role permission: %w", err)
	}
	if RowsAffected == 0 {
		return fmt.Errorf("role permission not found")
	}
	return nil
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
