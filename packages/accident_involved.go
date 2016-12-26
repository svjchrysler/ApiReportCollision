package packages

//AccidentInvolved class
type AccidentInvolved struct {
	AccidentID       int64  `gorm:"primary_key"`
	InvolvedID       int64  `gorm:"primary_key"`
	Gender           string `gorm:"type:varchar(1)"`
	PrivateInsurance int
	SOATInsurance    int
	NotFigured       int
}
