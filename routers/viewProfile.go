package routers

import (
	"encoding/json"
	"net/http"

	"github.com/RoyerHernandez/socialNetWork.git/db"
)

func ViewProfile(w http.ResponseWriter, r *http.Request) {
	Id := r.URL.Query().Get("id")
	if len(Id) < 1 {
		http.Error(w, "missing Id parameter", http.StatusBadRequest)
		return
	}
	profile, err := db.SearchProfile(Id)
	if err != nil {
		http.Error(w, "an error occurred while intent search the register", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}
