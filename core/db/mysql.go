package db

import (
	"message/models"

	"message/pkg/db"
)

func GetUser(id int64) (user *models.User, err error) {
	err = db.DB.Raw("SELECT * FROM users WHERE id = ?", id).Find(&user).Error
	return user, err
}
