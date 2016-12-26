package packages

//Severity class
type Severity struct {
	ID   int64  `gorm:"primary_key"`
	Name string `gorm:"not null"`
}
