package db

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/RoyerHernandez/socialNetWork.git/models"
)

func IntentLogin(email string, password string) (models.User, bool) {
	user, found, _ := AlreadyExistUser(email)
	if !found {
		return user, false
	}
	dbPassBytes := []byte(password)
	dbPass := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(dbPass, dbPassBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
