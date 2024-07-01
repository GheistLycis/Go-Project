package service

import (
	"errors"

	"github.com/go-project/app/shared"
	"github.com/go-project/app/user/structs"
	"github.com/go-project/database"
	"gorm.io/gorm"
)

/*
Patch is the usecase for patching an user in the database.
*/
func Patch(payload structs.PatchUser, ID int) (structs.GetUser, error) {
	user := structs.User{}

	// Finding user
	findTx := database.DB.First(&user, ID)

	if findTx.Error != nil {
		if errors.Is(findTx.Error, gorm.ErrRecordNotFound) {
			return structs.GetUser{}, errors.New("usuário não encontrado")
		}
		return structs.GetUser{}, findTx.Error
	}

	// Patching user
	shared.PatchStruct(user, payload)

	if err := user.Validate(); err != nil {
		return structs.GetUser{}, err
	}

	updateTx := database.DB.Model(&user).Updates()

	if updateTx.Error != nil {
		return structs.GetUser{}, updateTx.Error
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
