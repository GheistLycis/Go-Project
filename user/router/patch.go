package user_router

import (
	user_structs "github.com/go-project/user/structs"
)

func patch(payload user_structs.PatchUser, ID int) user_structs.GetUser {
	return user_structs.GetUser{}
}
