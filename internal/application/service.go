package application

import "fmt"

type Service interface {
	GetAllApplications() ([]ApplicationResponse, error)
}

type service struct {
	repo Repository
}

// GetAllApplications implements [Service].
func (s *service) GetAllApplications() ([]ApplicationResponse, error) {
	applications, err := s.repo.GetAllApplications()
	if err != nil {
		return nil, fmt.Errorf("failed to get all applications: %w", err)
	}
	var responses []ApplicationResponse
	for _, app := range applications {
		responses = append(responses, ToApplicationResponse(app))
	}
	return responses, nil
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
