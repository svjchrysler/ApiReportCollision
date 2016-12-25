package packages

import (
	"github.com/jinzhu/gorm"
)

//Photo class
type Photo struct {
	gorm.Model
	Accidents Accident
	Filename  string
}
