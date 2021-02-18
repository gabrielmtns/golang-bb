package main

import (
	"fmt"
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

	log.Fatal(http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), r))
}