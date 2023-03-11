package main

import (
	"log"

	"github.com/RoyerHernandez/socialNetWork.git/db"
	"github.com/RoyerHernandez/socialNetWork.git/handlers"
)

func main() {
	if db.ConnectionCheck() == false {
		log.Fatalf("whitout connection to data base")
		return
	}
	handlers.HandlerFunc()
}
