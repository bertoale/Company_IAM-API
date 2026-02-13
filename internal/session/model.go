package session

type Session struct {
	ID     uint `gorm:"primaryKey;autoIncrement"`
	UserID string
}
