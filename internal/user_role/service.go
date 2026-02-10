package user_role

type Service interface {
	Create(req *UserRoleRequest) (*UserRoleResponse, error)
	Delete(userRoleID uint) error
}

type service struct {
	repo Repository
}

// Create implements [Service].
func (s *service) Create(req *UserRoleRequest) (*UserRoleResponse, error) {
	userRole := &UserRole{
		UserID: req.UserID,
		RoleID: req.RoleID,
	}
	if err := s.repo.Create(userRole); err != nil {
		return nil, err
	}
	return ToUserRoleResponse(userRole), nil
}

// Delete implements [Service].
func (s *service) Delete(userRoleID uint) error {
	return s.repo.Delete(&UserRole{ID: userRoleID})
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
