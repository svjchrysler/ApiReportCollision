package packages

import (
	"github.com/jinzhu/gorm"
)

//Involved class
type Involved struct {
	gorm.Model
	Name string
}
