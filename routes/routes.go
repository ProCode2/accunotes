package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/procode2/accunotes/handlers"
	"github.com/procode2/accunotes/middlewares"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

func apiFunc(f CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SetupRoutes(m *mux.Router) {
	m.HandleFunc("/signup", apiFunc(handlers.HandlePostSignup)).Methods("POST")
	m.HandleFunc("/login", apiFunc(handlers.HandlePostLogin)).Methods("POST")

	note := m.PathPrefix("/notes").Subrouter()
	note.Use(middlewares.AuthenticatedRoutes)
	note.HandleFunc("", apiFunc(handlers.HandleGetNotes)).Methods("GET")
	note.HandleFunc("", apiFunc(handlers.HandlePostNote)).Methods("POST")
	note.HandleFunc("", apiFunc(handlers.HandleDeleteNote)).Methods("DELETE")
}
