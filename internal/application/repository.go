package application

import "gorm.io/gorm"

type Repository interface {
	GetAllApplications() ([]Application, error)
}

type repository struct {
	db *gorm.DB
}

// GetAllApplications implements [Repository].
func (r *repository) GetAllApplications() ([]Application, error) {
	var applications []Application
	if err := r.db.Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
