package structs

import "time"

/*
PatchUser is the API request payload for patching an user.
*/
type PatchUser struct {
	BirthDate  time.Time `json:"birthDate"`
	Gender     Gender    `json:"gender"`
	Name       string    `json:"name"`
	Profession string    `json:"profession"`
}
