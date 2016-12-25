package packages

import (
	"github.com/jinzhu/gorm"
)

//Climate class
type Climate struct {
	gorm.Model
	Name string
}
