package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/RoyerHernandez/socialNetWork.git/db"
	"github.com/RoyerHernandez/socialNetWork.git/jwt"
	"github.com/RoyerHernandez/socialNetWork.git/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "user and/or password incorrect, please try again "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "user email is required ", 400)
		return
	}
	document, userExist, _ := db.AlreadyExistUser(user.Email)
	if !userExist {
		http.Error(w, "user and/or password incorrect, please try again ", 400)
		return
	}
	jwtKey, err := jwt.GeneratedJwt(document)
	if err != nil {
		http.Error(w, "an error has occurred while generating the corresponding token  "+err.Error(), 400)
		return
	}
	response := models.LoginResponse{
		Token: jwtKey,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
