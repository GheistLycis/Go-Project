package service

import (
	structs "github.com/go-project/app/user/structs"
	db "github.com/go-project/database"
)

/*
Create is the usecase for inserting an user into the database.
*/
func Create(payload structs.CreateUser) (structs.GetUser, error) {
	user := structs.User{
		BirthDate:  payload.BirthDate,
		Gender:     payload.Gender,
		Name:       payload.Name,
		Profession: payload.Profession,
	}

	if err := user.Validate(); err != nil {
		return structs.GetUser{}, err
	}

	res := db.DB.Create(&user)

	if res.Error != nil {
		return structs.GetUser{}, res.Error
	}

	return structs.GetUser{
			BirthDate:  user.BirthDate,
			Gender:     user.Gender,
			ID:         user.ID,
			Name:       user.Name,
			Profession: user.Profession,
		},
		nil
}
