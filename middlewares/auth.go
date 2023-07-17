package middlewares

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/procode2/accunotes/database"
	"github.com/procode2/accunotes/utils"
)

type AuthenticRequestBody struct {
	Sid string `json:"sid"`
}

func AuthenticatedRoutes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Checking if user is authentic")
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		r.Body.Close()
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		body := &AuthenticRequestBody{}
		if err := utils.BodyAsJson(ioutil.NopCloser(bytes.NewBuffer(bodyBytes)), body); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		jwt, err := utils.ValiadateTokenString(body.Sid)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		user, err := database.Db.GetUserById(jwt.Id)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "userId", user.ID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
