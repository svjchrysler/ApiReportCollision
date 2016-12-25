package packages

import (
	"github.com/jinzhu/gorm"
)

//Severity class
type Severity struct {
	gorm.Model
	Name string
}
