package application

type Application struct {
	ID       uint   `gorm:"primaryKey"`
	Code     string `gorm:"uniqueIndex;size:100;not null"`
	Name     string `gorm:"size:255;not null"`
	IsActive bool   `gorm:"default:true"`
}

type ApplicationResponse struct {
	ID       uint   `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}
