package user_router

import (
	"time"

	user_structs "github.com/go-project/user/structs"
)

var mockUser = user_structs.User{
	ID: 0,
	CreateUser: user_structs.CreateUser{
		Name:       "Bruno",
		Birth_date: time.Now(),
		Gender:     user_structs.Male,
		Profession: "Programmer",
	},
}

func Get(ID string) user_structs.User {
	return mockUser
}

func List() []user_structs.User {
	return []user_structs.User{mockUser}
}
