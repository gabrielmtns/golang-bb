package main

import (
	"fmt"
	"github.com/rs/cors"
	"go-bb-test/src/infra/database"
	"go-bb-test/src/main/configs"
	"go-bb-test/src/main/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Rodando API")

	utils.LoadEnvs()

	databaseConnection, err := database.Handler()

	if err != nil {
		log.Fatalf("Error - Database connect %v", err)
	}



	r := configs.MapRoutes(databaseConnection)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(r)

	log.Fatal(
		http.ListenAndServe(":"+os.Getenv("SERVER_PORT"),
		handler),
	)
}