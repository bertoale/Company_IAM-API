package user_application

import "fmt"

type Service interface {
	Create(req *UserApplicationRequest) (*UserApplicationResponse, error)
	Delete(userID, ApplicationID uint) error
	GetByUserID(userID uint) (*UserWithApplicationsResponse, error)
	GetByApplicationID(applicationID uint) (*ApplicationWithUsersResponse, error)
}

type service struct {
	repository Repository
}

// GetByApplicationID implements [Service].
func (s *service) GetByApplicationID(applicationID uint) (*ApplicationWithUsersResponse, error) {
	userApplications, err := s.repository.FindByApplicationID(applicationID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user applications by application ID: %w", err)
	}
	return ToApplicationWithUsersResponse(applicationID, userApplications), nil
}

// GetByUserID implements [Service].
func (s *service) GetByUserID(userID uint) (*UserWithApplicationsResponse, error) {
	userApplications, err := s.repository.FindByUserID(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user applications by user ID: %w", err)
	}
	return ToUserWithApplicationsResponse(userID, userApplications), nil	
}

// Create implements [Service].
func (s *service) Create(req *UserApplicationRequest) (*UserApplicationResponse, error) {
	// 1️⃣ Cek apakah sudah ada
	exists, err := s.repository.Exists(req.UserID, req.ApplicationID)
	if err != nil {
		return nil, fmt.Errorf("failed to check existing user application: %w", err)
	}

	if exists {
		return nil, fmt.Errorf("application already assigned to user")
	}

	// 2️⃣ Create kalau belum ada
	userApplication := &UserApplication{
		UserID:        req.UserID,
		ApplicationID: req.ApplicationID,
	}

	if err := s.repository.Create(userApplication); err != nil {
		return nil, fmt.Errorf("failed to create user application: %w", err)
	}

	return ToUserApplicationResponse(userApplication), nil

	// Delete implements [Service].
}
func (s *service) Delete(userID, ApplicationID uint) error {
	return s.repository.Delete(
		userID,
		ApplicationID,
	)
}

func NewService(repository Repository) Service {
	return &service{repository: repository}
}
