package user

import (
	"time"

	"gorm.io/gorm"
)

type StatusEnum string

const (
	Active    StatusEnum = "active"
	Inactive  StatusEnum = "inactive"
	Suspended StatusEnum = "suspended"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`
	Email     string         `gorm:"unique;not null"`
	Username  string         `gorm:"not null"`
	Password  string         `gorm:"not null"`
	Status    StatusEnum     `gorm:"type:enum('active', 'inactive', 'suspended');default:'active'"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserResponse struct {
	ID         uint       `json:"id"`
	Email      string     `json:"email"`
	Name       string     `json:"name"`
	StatusEnum StatusEnum `json:"status"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Username string `json:"name" `
	Password string `json:"password"`
}

type UserUpdateRequest struct {
	Email    *string `json:"email"`
	Username *string `json:"name" `
	Status   *string `json:"status"`
	Password *string `json:"password"`
}

type LoginRequest struct {
	Identifier string `json:"email"`
	Password   string `json:"password"`
}

type LoginResponse struct {
	ID         uint   `json:"id"`
	Identifier string `json:"email"`
	Token      string `json:"token"`
}
