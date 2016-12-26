package packages

//Weather class
type Weather struct {
	ID   int64  `gorm:"primary_key"`
	Name string `gorm:"not null"`
}
