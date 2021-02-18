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

type UserController struct {
	Users usecases.IUsers
}

func NewUserController (u usecases.IUsers) *UserController {
	return &UserController{
		Users: u,
	}
}

func (u UserController) CreateUserController(w http.ResponseWriter, r *http.Request){
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		// TODO - reuse - encaps
		log.Printf("Failed on read user body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user domain.User

	if err:= json.Unmarshal(body, &user); err !=nil {
		// TODO - reuse - encaps
		log.Printf("Failed on unmarshal user body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	u.Users.Add(user)
	utils.JSONResponse(w, user)
}

func (u UserController) GetAllUsersController(w http.ResponseWriter, r *http.Request){

	result, err := u.Users.GetAll()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, result)

}

func (u UserController) GetUserByIdController(w http.ResponseWriter, r *http.Request){

}