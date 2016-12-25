package packages

import (
	"github.com/jinzhu/gorm"
)

//User class
type User struct {
	gorm.Model
	CI       int32
	Password string
}
