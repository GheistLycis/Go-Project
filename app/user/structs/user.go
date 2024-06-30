package structs

import (
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
	Gender     Gender
	ID         uint64
	Name       string
	Profession string
}
