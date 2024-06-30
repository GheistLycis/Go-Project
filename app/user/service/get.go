package service

import (
	structs "github.com/go-project/app/user/structs"
)

/*
Get is the usecase for finding an user in the database.
*/
func Get(ID int) structs.GetUser {
	return structs.GetUser{}
}
