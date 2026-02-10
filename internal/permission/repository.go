package permission

import "gorm.io/gorm"

type Repository interface {
	GetAllPermissions() ([]Permission, error)
}

type repository struct {
	db *gorm.DB
}

// GetAllPermissions implements [Repository].
func (r *repository) GetAllPermissions() ([]Permission, error) {
	var permissions []Permission
	if err := r.db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	return permissions, nil
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
