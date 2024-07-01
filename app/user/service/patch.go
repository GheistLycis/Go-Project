package service

import (
	"errors"

	"github.com/go-project/app/user/structs"
	"github.com/go-project/database"
	"gorm.io/gorm"
)

/*
Patch is the usecase for patching an user in the database.
*/
func Patch(payload structs.PatchUser, ID int) (structs.GetUser, error) {
	user := structs.User{}

	res := database.DB.First(&user, ID)

	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return structs.GetUser{}, errors.New("usuário não encontrado")
		}
		return structs.GetUser{}, res.Error
	}

	// database.DB.Model(&user).Updates()

	if err := user.Validate(); err != nil {
		return structs.GetUser{}, err
	}

	// res := database.DB.Create(&user)

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
