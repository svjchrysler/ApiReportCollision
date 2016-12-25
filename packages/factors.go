package packages

import (
	"github.com/jinzhu/gorm"
)

//Factors class
type Factors struct {
	gorm.Model
	Name string
}
