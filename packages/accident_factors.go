package packages

import "time"

//AccidentFactors class
type AccidentFactors struct {
	AccidentID int64     `gorm:"primary_key"`
	FactorsID  int64     `gorm:"primary_key"`
	CreateAt   time.Time `gorm:"not null"`
	UpdatedAt  time.Time
}
