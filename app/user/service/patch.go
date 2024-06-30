package service

import (
	structs "github.com/go-project/app/user/structs"
)

/*
Patch is the usecase for patching an user in the database.
*/
func Patch(payload structs.PatchUser, ID int) structs.GetUser {
	return structs.GetUser{}
}
