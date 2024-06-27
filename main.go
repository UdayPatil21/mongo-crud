package main

import (
	"fmt"
	"log"
	"mongo-crud/cmd"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to Golang")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := cmd.NewApp()
	app.InitializeRoutes()
	log.Println("Server is running ....")
	http.ListenAndServe(":8080", app.Router)
}
