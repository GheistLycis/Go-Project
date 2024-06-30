package structs

import "time"

/*
ListUser is the API response payload for returning multiple users.
*/
type ListUser struct {
	BirthDate  time.Time `json:"birthDate" binding:"required"`
	Gender     Gender    `json:"gender" binding:"required"`
	ID         int       `json:"id" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Profession string    `json:"profession" binding:"required"`
}
