package packages

//AccidentFactors class
type AccidentFactors struct {
	AccidentID Accident `gorm:"primary_key"`
	FactorsID  Factors  `gorm:"primary_key"`
}
