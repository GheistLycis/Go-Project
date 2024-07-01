package service

import (
	"github.com/go-project/app/user/structs"
	"github.com/go-project/database"
)

/*
List is the usecase for listing users from the database.
*/
func List() ([]structs.ListUser, error) {
	users := []structs.User{}

	if res := database.DB.Find(&users); res.Error != nil {
		return nil, res.Error
	}

	listUsers := make([]structs.ListUser, len(users))

	for i, user := range users {
		listUsers[i] = structs.ListUser{
			BirthDate:  user.BirthDate,
			Gender:     user.Gender,
			ID:         user.ID,
			Name:       user.Name,
			Profession: user.Profession,
		}
	}

	return listUsers, nil
}
