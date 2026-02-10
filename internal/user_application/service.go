package user_application

type Service interface {
	Create(req *UserApplicationRequest) error
	Delete(id uint) error
}

type service struct {
	repository Repository
}

// Create implements [Service].
func (s *service) Create(req *UserApplicationRequest) error {
	userApplication := &UserApplication{
		UserID:        req.UserID,
		ApplicationID: req.ApplicationID,
	}
	return s.repository.Create(userApplication)
}

// Delete implements [Service].
func (s *service) Delete(id uint) error {
	return s.repository.Delete(&UserApplication{ID: id})
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}
