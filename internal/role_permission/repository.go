package role_permission

import "gorm.io/gorm"

type Repository interface {
	Create(rolePermission *RolePermission) error
	Delete(rolePermission *RolePermission) error
}

type repository struct {
	db *gorm.DB
}

// Create implements [Repository].
func (r *repository) Create(rolePermission *RolePermission) error {
	return r.db.Create(rolePermission).Error
}

// Delete implements [Repository].
func (r *repository) Delete(rolePermission *RolePermission) error {
	return r.db.Delete(rolePermission).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
