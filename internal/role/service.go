package role

import "fmt"

type Service interface {
	CreateRole(req *RoleRequest) (*RoleResponse, error)
	GetRoleByID(id uint) (*RoleResponse, error)
	GetAllRoles() ([]RoleResponse, error)
	UpdateRole(id uint, req *UpdateRoleRequest) (*RoleResponse, error)
	DeleteRole(id uint) error
}

type service struct {
	repo Repository
}

// CreateRole implements [Service].
func (s *service) CreateRole(req *RoleRequest) (*RoleResponse, error) {
	existingName, _ := s.repo.FindByName(req.Name)
	if existingName != nil {
		return nil, fmt.Errorf("role with name '%s' already exists", req.Name)
	}
	newRole := &Role{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := s.repo.Create(newRole); err != nil {
		return nil, err
	}

	return toRoleResponse(newRole), nil
}

// DeleteRole implements [Service].
func (s *service) DeleteRole(id uint) error {
	err := s.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("failed to delete role with ID %d: %w", id, err)
	}
	return nil
}

// GetAllRoles implements [Service].
func (s *service) GetAllRoles() ([]RoleResponse, error) {
	roles, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve roles: %w", err)
	}
	var responses []RoleResponse
	for _, role := range roles {
		responses = append(responses, *toRoleResponse(&role))
	}
	return responses, nil
}

// GetRoleByID implements [Service].
func (s *service) GetRoleByID(id uint) (*RoleResponse, error) {
	role, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve role with ID %d: %w", id, err)
	}
	return toRoleResponse(role), nil
}

// UpdateRole implements [Service].
func (s *service) UpdateRole(id uint, req *UpdateRoleRequest) (*RoleResponse, error) {
	role, err := s.repo.FindByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve role with ID %d: %w", id, err)
	}
	if req.Name != nil {
		existingName, _ := s.repo.FindByName(*req.Name)
		if existingName != nil && existingName.ID != id {
			return nil, fmt.Errorf("role with name '%s' already exists", *req.Name)
		}
		role.Name = *req.Name
	}
	if req.Description != nil {
		role.Description = *req.Description
	}
	if err := s.repo.Update(role); err != nil {
		return nil, fmt.Errorf("failed to update role with ID %d: %w", id, err)
	}
	return toRoleResponse(role), nil
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
