package repo

import (
	"core/app"
	"core/internal/model"
)

func CreateAdmin(user model.User) (model.User, error) {
	db := app.Database.DB
	crUser := db.Create(&user)
	return user, crUser.Error
}
