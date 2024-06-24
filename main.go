package main

import (
	"fmt"
	"log"
	"mongo-crud/cmd"
	"net/http"
)

func main() {
	fmt.Println("Welcome to Golang")
	app := cmd.NewApp()
	app.InitializeRoutes()
	log.Println("Server is running ....")
	http.ListenAndServe(":8080", app.Router)
}
