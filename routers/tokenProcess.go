package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"github.com/RoyerHernandez/socialNetWork.git/db"
	"github.com/RoyerHernandez/socialNetWork.git/models"
)

var Email string
var UserId string

func TokenProcess(token string) (*models.Claim, bool, string, error) {

	myKey := []byte("Israfel_RoyerHernandez")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, "", errors.New("invalid token format")
	}
	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err == nil {
		_, found, _ := db.AlreadyExistUser(claims.Email)
		if found {
			Email = claims.Email
			UserId = claims.Id.Hex()
		}
		return claims, true, UserId, nil
	}
	if !tkn.Valid {
		return claims, false, "", errors.New("invalid token")
	}
	return claims, false, "", err
}
