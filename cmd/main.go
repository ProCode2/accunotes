package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	"github.com/procode2/accunotes/database"
	"github.com/procode2/accunotes/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.Db.Init()

	rtr := mux.NewRouter()
	routes.SetupRoutes(rtr)
	http.Handle("/", rtr)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
