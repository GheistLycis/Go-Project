package structs

import "time"

/*
CreateUser is the API request payload for creating an user.
*/
type CreateUser struct {
	Name       string    `json:"name" binding:"required"`
	BirthDate  time.Time `json:"birthDate" binding:"required"`
	Gender     Gender    `json:"gender" binding:"required"`
	Profession string    `json:"profession" binding:"required"`
}
