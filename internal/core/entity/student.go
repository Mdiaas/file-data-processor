package entity

import "time"

const BirthDateFormat = "01/02/2006"

var DefaultDateWhenItsInvalid = time.Date(0001, time.January, 1, 0, 0, 0, 0, time.UTC)

type Student struct {
	Id             string
	FullName       string
	DocumentNumber string
	BirthDate      time.Time
	MotherName     string
	FatherName     string
	Errors         []string
}
