package packages

//Involved class
type Involved struct {
	ID   int64  `gorm:"primary_key"`
	Name string `gorm:"not null"`
}
