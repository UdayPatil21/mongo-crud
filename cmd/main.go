package main

import (
	"fmt"
	"log"
	"mongo-crud/api/handler"
	"mongo-crud/internal/database"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	DB     *database.Database
	Router *mux.Router
}

func NewApp() *App {
	db, err := database.InitDB("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	route := mux.NewRouter()
	return &App{DB: db, Router: route}
}
func (a *App) InitializeRoutes() {
	handler := handler.NewApp(a.DB.Client)
	r := a.Router.PathPrefix("/api").Subrouter()
	r.HandleFunc("/", handler.CreateHandler).Methods("POST")
	r.HandleFunc("/", handler.GetAllHandler).Methods("GET")
	r.HandleFunc("/{id}", handler.UpdateHandler).Methods("PUT")
	r.HandleFunc("/{id}", handler.DeleteHandler).Methods("DELETE")
	r.HandleFunc("/{id}", handler.GetHandler).Methods("GET")

}

func main() {
	fmt.Println("Welcome to Golang")
	app := NewApp()
	app.InitializeRoutes()
	log.Println("Server is running ....")
	http.ListenAndServe(":8080", app.Router)
}
