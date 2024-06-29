package user_structs

import "time"

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Other  Gender = "O"
)

type User struct {
	BirthDate  time.Time
	Gender     Gender
	ID         int
	Name       string
	Profession string
}
