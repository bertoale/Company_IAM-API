package user

import "gorm.io/gorm"

type Repository interface {
	Create(user *User) error
	FindByID(id uint) (*User, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindAll() ([]User, error)
	Update(user *User) error
	Delete(user *User) error
}

type repository struct {
	db *gorm.DB
}

// Create implements [Repository].
func (r *repository) Create(user *User) error {
	return r.db.Create(user).Error
}

// Delete implements [Repository].
func (r *repository) Delete(user *User) error {
	return r.db.Delete(user).Error
}

// FindAll implements [Repository].
func (r *repository) FindAll() ([]User, error) {
	var users []User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindByEmail implements [Repository].
func (r *repository) FindByEmail(email string) (*User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID implements [Repository].
func (r *repository) FindByID(id uint) (*User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername implements [Repository].
func (r *repository) FindByUsername(username string) (*User, error) {
	var user User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Update implements [Repository].
func (r *repository) Update(user *User) error {
	return r.db.Save(user).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}
