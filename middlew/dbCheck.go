package middlew

import (
	"net/http"

	"github.com/RoyerHernandez/socialNetWork.git/db"
)

func DBCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !db.ConnectionCheck() {
			http.Error(w, "Lost connection to data base", 500)
		}
		next.ServeHTTP(w, r)
	}
}
