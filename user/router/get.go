package user_router

import (
	user_structs "github.com/go-project/user/structs"
)

func get(ID int) user_structs.GetUser {
	return user_structs.GetUser{}
}

func list() []user_structs.ListUser {
	return []user_structs.ListUser{}
}
