package user_service

import (
	user_structs "github.com/go-project/app/user/structs"
)

func Get(ID int) user_structs.GetUser {
	return user_structs.GetUser{}
}
