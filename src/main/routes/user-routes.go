package routes

import (
	"github.com/gorilla/mux"
	"go-bb-test/src/app/controllers"
	"net/http"
)


func UserRoutes(r *mux.Router, handle  *controllers.UserController){
	r.HandleFunc("/users", handle.CreateUserController).Methods(http.MethodPost)
	r.HandleFunc("/users", handle.GetAllUsersController).Methods(http.MethodGet)
}