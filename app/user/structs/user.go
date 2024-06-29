package user_structs

import (
	"time"

	"gorm.io/gorm"
)

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Other  Gender = "O"
)

type User struct {
	gorm.Model
	BirthDate  time.Time
	Gender     Gender
	GovId      uint64 `gorm:"primaryKey"`
	Name       string
	Profession string
}
