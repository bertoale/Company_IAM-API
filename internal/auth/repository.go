package auth

import (
	"company_iam/internal/user"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByIdentifier(identifier string) (*user.User, error)
	GetUserRoles(userID uint) ([]string, error)
	GetUserPermissions(userID uint) ([]string, error)
	GetUserApplications(userID uint) ([]string, error)
}

type repository struct {
	db *gorm.DB
}

// GetUserByIdentifier implements [Repository].
func (r *repository) GetUserByIdentifier(identifier string) (*user.User, error) {
	var user user.User
	if err := r.db.
		Where("email = ? OR username = ?", identifier, identifier).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserRoles implements [Repository].
func (r *repository) GetUserRoles(userID uint) ([]string, error) {
	var roles []string
	err := r.db.Table("user_roles").
		Select("roles.name").
		Joins("JOIN roles ON roles.id = user_roles.role_id").
		Where("user_roles.user_id = ?", userID).
		Pluck("roles.name", &roles).Error
	return roles, err
}

// GetUserPermissions implements [Repository].
func (r *repository) GetUserPermissions(userID uint) ([]string, error) {
	var permissions []string
	err := r.db.Table("user_roles").
		Select("DISTINCT permissions.code").
		Joins("JOIN role_permissions ON role_permissions.role_id = user_roles.role_id").
		Joins("JOIN permissions ON permissions.id = role_permissions.permission_id").
		Where("user_roles.user_id = ?", userID).
		Pluck("permissions.permission_name", &permissions).Error
	return permissions, err
}

// GetUserApplications implements [Repository].
func (r *repository) GetUserApplications(userID uint) ([]string, error) {
	var applications []string
	err := r.db.Table("user_applications").
		Select("applications.code").
		Joins("JOIN applications ON applications.id = user_applications.application_id").
		Where("user_applications.user_id = ?", userID).
		Pluck("applications.code", &applications).Error
	return applications, err
}


func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
