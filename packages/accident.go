package packages

import "time"

//Accident class
type Accident struct {
	ID           int64 `gorm:"primary_key"`
	WeatherID    int64 `gorm:"not null"`
	PersonID     int64
	SeverityID   int64
	DateAccident time.Time `gorm:"not null"`
	Time         time.Time `gorm:"not null"`
	Latitude     float64   `gorm:"not null"`
	Longitude    float64   `gorm:"not null"`
	CreateAt     time.Time `gorm:"not null"`
	UpdatedAt    time.Time
}
