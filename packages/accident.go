package packages

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Accident class
type Accident struct {
	gorm.Model
	RegistrationDate time.Time
	DateAccident     time.Time
	Time             time.Time
	People           Person
	Climates         Climate
	Severities       Severity
}
