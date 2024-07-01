package service

import (
	"errors"

	"github.com/go-project/app/user/structs"
	"github.com/go-project/database"
	"gorm.io/gorm"
)

/*
Get is the usecase for finding an user in the database.
*/
func Get(ID int) (structs.GetUser, error) {
	user := structs.User{}

	res := database.DB.First(&user, ID)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return structs.GetUser{}, errors.New("usuário não encontrado")
		}
		return structs.GetUser{}, res.Error
	}

	return structs.GetUser{
		BirthDate:  user.BirthDate,
		Gender:     user.Gender,
		ID:         user.ID,
		Name:       user.Name,
		Profession: user.Profession,
	}, nil
}
