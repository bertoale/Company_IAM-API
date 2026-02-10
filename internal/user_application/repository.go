package user_application

import "gorm.io/gorm"

type Repository interface {
	Create(userApplication *UserApplication) error
	Delete(userApplication *UserApplication) error
}

type repository struct {
	db *gorm.DB
}

// Create implements [Repository].
func (r *repository) Create(userApplication *UserApplication) error {
	return r.db.Create(userApplication).Error
}

// Delete implements [Repository].
func (r *repository) Delete(userApplication *UserApplication) error {
	return r.db.Delete(userApplication).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
