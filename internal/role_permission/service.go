package role_permission

type Service interface {
	CreateRolePermission(req *RolePermissionRequest) (*RolePermissionResponse, error)
	DeleteRolePermission(rolePermissionID uint) error
}

type service struct {
	repo Repository
}

// DeleteRolePermission implements [Service].
func (s *service) DeleteRolePermission(rolePermissionID uint) error {
	return s.repo.Delete(&RolePermission{ID: rolePermissionID})
}

// CreateRolePermission implements [Service].
func (s *service) CreateRolePermission(req *RolePermissionRequest) (*RolePermissionResponse, error) {
	rolePermission := &RolePermission{
		RoleID:       req.RoleID,
		PermissionID: req.PermissionID,
	}
	if err := s.repo.Create(rolePermission); err != nil {
		return nil, err
	}
	return ToRolePermissionResponse(rolePermission), nil
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
