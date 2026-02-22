package rbac

type Role struct {
	ID          uint
	Name        string `gorm:"uniqueIndex"`
	Application string
	Permissions []Permission `gorm:"many2many:role_permissions;"`
}

type Permission struct {
	ID   uint   `gorm:"column:id"`
	Code string `gorm:"column:code"`
}