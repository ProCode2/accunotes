package handlers

import (
	"fmt"
	"net/http"

	"github.com/procode2/accunotes/database"
	"github.com/procode2/accunotes/models"
	"github.com/procode2/accunotes/utils"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func HandlePostSignup(w http.ResponseWriter, r *http.Request) error {
	body := &SignUpRequest{}
	if err := utils.BodyAsJson(r.Body, body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	pswd, err := utils.HashPassword(body.Password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	newUser := &models.User{
		Email:    body.Email,
		Name:     body.Name,
		Password: pswd,
	}

	_, err = database.Db.CreateNewUser(newUser)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}

func HandlePostLogin(w http.ResponseWriter, r *http.Request) error {
	body := &LoginRequest{}
	if err := utils.BodyAsJson(r.Body, body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	user, err := database.Db.GetUserByEmail(body.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	if !utils.CheckPasswordHash(body.Password, user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return nil
	}

	// authentic user
	jwtString, err := utils.GetJWTKey(fmt.Sprint(user.ID))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	return utils.WriteJson(w, map[string]string{"sid": jwtString})
}
