package packages

import (
	"github.com/jinzhu/gorm"
)

//Person class
type Person struct {
	gorm.Model
	Users    User
	Name     string
	LastName string
}
