package structs

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

/*
const (

	Male   Gender = "M"
	Female Gender = "F"
	Other  Gender = "O"

)
*/
type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Other  Gender = "O"
)

/*
User is the user representation in the database.
*/
type User struct {
	gorm.Model
	BirthDate  time.Time
	Gender     Gender `gorm:"type:gender"`
	ID         uint64
	Name       string
	Profession string
}

/*
Validate checks for any business rules broken by the instance
that the model itself is unaware of.
*/
func (u *User) Validate() error {
	// Checking for min BirthDate
	minDate := time.Date(1900, 1, 1, 0, 0, 0, 0, time.Local)

	if u.BirthDate.Before(minDate) {
		return errors.New("data de nascimento não pode ser anterior a 01/01/1900")
	}

	return nil
}
