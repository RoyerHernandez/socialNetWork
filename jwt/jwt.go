package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/RoyerHernandez/socialNetWork.git/models"
)

func GeneratedJwt(user models.User) (string, error) {
	myKey := []byte("Israfel_RoyerHernandez")
	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastName":  user.LastName,
		"birthDate": user.BirthDate,
		"Biography": user.Biography,
		"Location":  user.Location,
		"WebSite":   user.WebSite,
		"_id":       user.ID,
		"exp":       time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
