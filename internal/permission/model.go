package permission

type Permission struct {
	ID          uint           `gorm:"primaryKey"`
	Code        string         `gorm:"unique;not null"`
	Description string         `gorm:"nullable"`
}

type PermissionResponse struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
