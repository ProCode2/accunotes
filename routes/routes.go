package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/procode2/accunotes/handlers"
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

}
