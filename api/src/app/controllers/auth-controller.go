package controllers

import (
	"encoding/json"
	"go-bb-test/src/app/utils"
	domain "go-bb-test/src/domain/models"
	"go-bb-test/src/domain/usecases"
	"io/ioutil"
	"log"
	"net/http"
)

type AuthController struct {
	Auth usecases.IAuth
}

func NewAuthController (a usecases.IAuth) *AuthController {
	return &AuthController{
		Auth: a,
	}
}


func (a AuthController) LoginController(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// TODO - reuse - encaps
		log.Printf("Failed on read user body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user domain.User

	if err:= json.Unmarshal(body, &user); err !=nil {
		log.Printf("Failed on unmarshal user body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}


	result, err := a.Auth.Login(user.Username, user.Password)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	utils.JSONResponse(w, result)
}
