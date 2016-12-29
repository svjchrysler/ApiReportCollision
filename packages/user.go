package packages

import "time"

//User class
type User struct {
	ID        int64     `gorm:"primary_key"`
	CI        int       `gorm:"not null;unique"`
	Password  string    `gorm:"not null"`
	CreateAt  time.Time `gorm:"not null"`
	UpdatedAt time.Time
}
