package user_structs

import "time"

type PatchUser struct {
	BirthDate  time.Time `json:"birthDate"`
	Gender     Gender    `json:"gender"`
	Name       string    `json:"name"`
	Profession string    `json:"profession"`
}
