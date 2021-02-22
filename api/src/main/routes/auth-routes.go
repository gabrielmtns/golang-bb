package routes

import (
	"github.com/gorilla/mux"
	"go-bb-test/src/app/controllers"
	"net/http"
)

func AuthRoutes(r *mux.Router, handle  *controllers.AuthController){
	r.HandleFunc("/login", handle.LoginController).Methods(http.MethodPost)
}