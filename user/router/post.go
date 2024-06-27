package user_router

import user_structs "github.com/go-project/user/structs"

func Create(payload user_structs.CreateUser) user_structs.User {
	return user_structs.User{
		CreateUser: payload,
		ID:         0,
	}
}
