package routers

import (
	"encoding/json"
	"github.com/RoyerHernandez/socialNetWork.git/db"
	"net/http"

	"github.com/RoyerHernandez/socialNetWork.git/models"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "received data error "+err.Error(), 400)
		return
	}
	if len(user.Email) == 0 {
		http.Error(w, "user email is required "+err.Error(), 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "password field should have to more than 6 characters "+err.Error(), 400)
		return
	}

	_, findUser, _ := db.AlreadyExistUser(user.Email)

	if findUser {
		http.Error(w, "user already been registered with this email "+err.Error(), 400)
		return
	}
	_, status, err := db.InsertRegister(user)
	if err != nil {
		http.Error(w, "an error occurred while insert the register "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "register not insert "+err.Error(), 400)
		return
	}

}
