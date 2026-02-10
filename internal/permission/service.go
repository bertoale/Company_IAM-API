package permission

type Service interface {
	GetAllPermissions() ([]PermissionResponse, error)
}

type service struct {
	repo Repository
}

// GetAllPermissions implements [Service].
func (s *service) GetAllPermissions() ([]PermissionResponse, error) {
	permissions, err := s.repo.GetAllPermissions()
	if err != nil {
		return nil, err
	}
	var responses []PermissionResponse
	for _, p := range permissions {
		responses = append(responses, ToPermissionResponse(p))
	}
	return responses, nil
}

// GetAllPermissions implements [Service].
func NewService(repo Repository) Service {
	return &service{repo: repo}
}
