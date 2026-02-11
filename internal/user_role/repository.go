package user_role

import "gorm.io/gorm"

type Repository interface {
	Create(userRole *UserRole) error
	Delete(userID uint, roleID uint) error
	Exists(userID uint, roleID uint) (bool, error)
	FindByUserID(userID uint) ([]UserRole, error)
	FindByRoleID(roleID uint) ([]UserRole, error)
	FindByUserIDWithRole(userID uint) ([]UserRole, error)
}

type repository struct {
	db *gorm.DB
}

func (r *repository) Create(userRole *UserRole) error {
	return r.db.Create(userRole).Error
}

func (r *repository) Delete(userID uint, roleID uint) error {
	return r.db.
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Delete(&UserRole{}).
		Error
}

func (r *repository) Exists(userID uint, roleID uint) (bool, error) {
	var count int64
	err := r.db.
		Model(&UserRole{}).
		Where("user_id = ? AND role_id = ?", userID, roleID).
		Count(&count).
		Error

	return count > 0, err
}

func (r *repository) FindByUserID(userID uint) ([]UserRole, error) {
	var userRoles []UserRole
	err := r.db.
		Where("user_id = ?", userID).
		Find(&userRoles).
		Error

	return userRoles, err
}

func (r *repository) FindByRoleID(roleID uint) ([]UserRole, error) {
	var userRoles []UserRole
	err := r.db.
		Where("role_id = ?", roleID).
		Find(&userRoles).
		Error

	return userRoles, err
}

func (r *repository) FindByUserIDWithRole(userID uint) ([]UserRole, error) {
	var userRoles []UserRole
	err := r.db.
		Preload("Role").
		Where("user_id = ?", userID).
		Find(&userRoles).
		Error

	return userRoles, err
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
