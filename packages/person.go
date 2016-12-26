package packages

import "time"

//Person class
type Person struct {
	ID        int64 `gorm:"primary_key"`
	UserID    int64
	Name      string    `gorm:"not null"`
	Lastname  string    `gorm:"not null"`
	CreateAt  time.Time `gorm:"not null"`
	UpdatedAt time.Time
}
