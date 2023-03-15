package routers

import (
	"encoding/json"
	"github.com/RoyerHernandez/socialNetWork.git/db"
	"github.com/RoyerHernandez/socialNetWork.git/models"
	"net/http"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {
	var (
		user   models.User
		status bool
	)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "incorrect data !!! "+err.Error(), 400)
		return
	}
	status, err = db.UpdateRegister(user, UserId)
	if err != nil {
		http.Error(w, "an error occurred while intent modify the register, please try again ", 400)
		return
	}
	if !status {
		http.Error(w, "it has not been modify the register", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
