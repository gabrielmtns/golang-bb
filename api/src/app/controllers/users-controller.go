package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-bb-test/src/app/utils"
	domain "go-bb-test/src/domain/models"
	"go-bb-test/src/domain/usecases"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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

	fmt.Println("AEOho", body)

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

	result, _ := u.Users.Add(user)

	fmt.Println(result)

	utils.JSONResponse(w, domain.User{
		Id: result.Id,
		Name: result.Name,
		Lastname: result.Lastname,
		Username: result.Username,
	})
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
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 0)
	result, err :=  u.Users.Get(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, result)
}

func (u UserController) UpdateUserController(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 0)
	var user domain.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	fmt.Print(user)
	result, err :=  u.Users.Update(id, user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, result)
}

func (u UserController) DeleteUserController(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 0)
	err :=  u.Users.Delete(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, nil)
}