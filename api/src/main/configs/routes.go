package configs

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go-bb-test/src/app/middlewares"
	"go-bb-test/src/main/instances"
	"go-bb-test/src/main/routes"
)

func MapRoutes (conn *sql.DB) *mux.Router {
	userController := instances.UserControllerFactory(conn)
	authController := instances.AuthControllerFactory(conn)
	router := mux.NewRouter()
	noAuthRouter := router.PathPrefix("/api/v1").Subrouter()

	authRouter := router.PathPrefix("/api/v1").Subrouter()
	authRouter.Use(middlewares.AuthMiddleware)

	routes.AuthRoutes(noAuthRouter, authController)
	routes.UserRoutes(authRouter, userController)



	return router
}


