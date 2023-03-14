package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/RoyerHernandez/socialNetWork.git/middlew"
	"github.com/RoyerHernandez/socialNetWork.git/routers"
)

func HandlerFunc() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.DBCheck(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.DBCheck(routers.Login)).Methods("POST")
	router.HandleFunc("/viewProfile", middlew.DBCheck(middlew.JWTValidate(routers.ViewProfile))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
