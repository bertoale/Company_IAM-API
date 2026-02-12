package user_application

import "gorm.io/gorm"

type Repository interface {
	Create(userApplication *UserApplication) error
	Delete(userID, applicationID uint) error
	Exists(userID uint, applicationID uint) (bool, error)
	FindByUserID(userID uint) ([]UserApplication, error)
	FindByApplicationID(applicationID uint) ([]UserApplication, error)
}

type repository struct {
	db *gorm.DB
}

// Exists implements [Repository].
func (r *repository) Exists(userID uint, applicationID uint) (bool, error) {
	var count int64
	err := r.db.
		Model(&UserApplication{}).
		Where("user_id = ? AND application_id = ?", userID, applicationID).
		Count(&count).
		Error

	return count > 0, err	
}

// Create implements [Repository].
func (r *repository) Create(userApplication *UserApplication) error {
	return r.db.Create(userApplication).Error
}

// Delete implements [Repository].
func (r *repository) Delete(userID, applicationID uint) error {
	return r.db.
		Where("user_id = ? AND application_id = ?", userID, applicationID).
		Delete(&UserApplication{}).
		Error
}

// FindByApplicationID implements [Repository].
func (r *repository) FindByApplicationID(applicationID uint) ([]UserApplication, error) {
	var userApplications []UserApplication
	err := r.db.
		Preload("User").
		Preload("Application").
		Where("application_id = ?", applicationID).
		Find(&userApplications).
		Error

	return userApplications, err
}

// FindByUserID implements [Repository].
func (r *repository) FindByUserID(userID uint) ([]UserApplication, error) {
	var userApplications []UserApplication
	err := r.db.
		Preload("User").
		Preload("Application").
		Where("user_id = ?", userID).
		Find(&userApplications).
		Error

	return userApplications, err
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
