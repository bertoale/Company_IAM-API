package user_role

import "gorm.io/gorm"

type Repository interface {
	Create(userRole *UserRole) error
	Delete(userRole *UserRole) error
}
type repository struct {
	db *gorm.DB
}

// Create implements [Repository].
func (r *repository) Create(userRole *UserRole) error {
	return r.db.Create(userRole).Error
}

// Delete implements [Repository].
func (r *repository) Delete(userRole *UserRole) error {
	return r.db.Delete(userRole).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
