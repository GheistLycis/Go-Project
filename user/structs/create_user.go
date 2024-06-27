package user_structs

import "time"

type CreateUser struct {
	Name       string    `json:"name" binding:"required"`
	Birth_date time.Time `json:"birthDate" binding:"required"`
	Gender     Gender    `json:"gender" binding:"required"`
	Profession string    `json:"profession" binding:"required"`
}
