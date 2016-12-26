package packages

import "time"

//Photo class
type Photo struct {
	ID         int64     `gorm:"primary_key"`
	AccidentID int64     `gorm:"not null"`
	Filename   string    `gorm:"not null"`
	CreateAt   time.Time `gorm:"not null"`
	UpdatedAt  time.Time
}
