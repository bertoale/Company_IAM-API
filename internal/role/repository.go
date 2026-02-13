package role

import "gorm.io/gorm"

type Repository interface {
	Create(role *Role) error
	FindByID(id uint) (*Role, error)
	FindByName(name string) (*Role, error)
	FindAll() ([]Role, error)
	Update(role *Role) error
	Delete(id uint) (int64, error)
}

type repository struct {
	db *gorm.DB
}

// Create implements [Repository].
func (r *repository) Create(role *Role) error {
	return r.db.Create(role).Error
}

// Delete implements [Repository].
func (r *repository) Delete(id uint) (int64, error) {
	result := r.db.Delete(&Role{}, id)
	return result.RowsAffected, result.Error
}

// GetAll implements [Repository].
func (r *repository) FindAll() ([]Role, error) {
	var roles []Role
	if err := r.db.Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}

// GetByID implements [Repository].
func (r *repository) FindByID(id uint) (*Role, error) {
	var role Role
	if err := r.db.First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// FindByName implements [Repository].
func (r *repository) FindByName(name string) (*Role, error) {
	var role Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

// Update implements [Repository].
func (r *repository) Update(role *Role) error {
	return r.db.Save(role).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
