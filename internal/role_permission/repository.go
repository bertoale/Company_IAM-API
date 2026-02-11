package role_permission

import "gorm.io/gorm"

type Repository interface {
	Create(rolePermission *RolePermission) error
	Delete(roleID uint, permissionID uint) error
	Exists(roleID uint, permissionID uint) (bool, error)
	FindByRoleID(roleID uint) ([]RolePermission, error)
	FindByPermissionID(permissionID uint) ([]RolePermission, error)
	FindByRoleIDWithPermission(roleID uint) ([]RolePermission, error)
}

type repository struct {
	db *gorm.DB
}

// Create implements [Repository].
func (r *repository) Create(rolePermission *RolePermission) error {
	return r.db.Create(rolePermission).Error
}

// Delete implements [Repository].
func (r *repository) Delete(roleID uint, permissionID uint) error {
	return r.db.
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Delete(&RolePermission{}).
		Error
}

// Exists implements [Repository].
func (r *repository) Exists(roleID uint, permissionID uint) (bool, error) {
	var count int64
	err := r.db.
		Model(&RolePermission{}).
		Where("role_id = ? AND permission_id = ?", roleID, permissionID).
		Count(&count).
		Error

	return count > 0, err
}

// FindByPermissionID implements [Repository].
func (r *repository) FindByPermissionID(permissionID uint) ([]RolePermission, error) {
	var rolePermissions []RolePermission
	err := r.db.
		Where("permission_id = ?", permissionID).
		Find(&rolePermissions).
		Error

	return rolePermissions, err
}

// FindByRoleID implements [Repository].
func (r *repository) FindByRoleID(roleID uint) ([]RolePermission, error) {
	var rolePermissions []RolePermission
	err := r.db.
		Where("role_id = ?", roleID).
		Find(&rolePermissions).
		Error

	return rolePermissions, err
}

// FindByRoleIDWithPermission implements [Repository].
func (r *repository) FindByRoleIDWithPermission(roleID uint) ([]RolePermission, error) {
	var rolePermissions []RolePermission
	err := r.db.
		Preload("Permission").
		Where("role_id = ?", roleID).
		Find(&rolePermissions).
		Error

	return rolePermissions, err
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
