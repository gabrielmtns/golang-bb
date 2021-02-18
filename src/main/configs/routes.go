package configs

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go-bb-test/src/main/instances"
	"go-bb-test/src/main/routes"
)

func MapRoutes (conn *sql.DB) *mux.Router {
	router := mux.NewRouter().PathPrefix("/api/v1").Subrouter()
	userController := instances.UserControllerFactory(conn)
	routes.UserRoutes(router, userController)

	return router
}