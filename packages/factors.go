package packages

//Factors class
type Factors struct {
	ID   int64  `gorm:"primary_key"`
	Name string `gorm:"not null"`
}
