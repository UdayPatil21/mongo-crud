package cmd

import (
	"log"
	"mongo-crud/api/handler"
	"mongo-crud/internal/database"
	"os"

	"github.com/gorilla/mux"
)

type App struct {
	DB     *database.Database
	Router *mux.Router
}

func NewApp() *App {

	mongoUrl := os.Getenv("MONGO_URI")
	if mongoUrl == "" {
		mongoUrl = "mongodb://localhost:27017"
		log.Println("MONGODB_URL  environment variable not set")
	}
	db, err := database.InitDB(mongoUrl)
	if err != nil {
		panic(err)
	}
	route := mux.NewRouter()
	return &App{DB: db, Router: route}
}
func (a *App) InitializeRoutes() {
	handler := handler.NewApp(a.DB.Client)

	r := a.Router.PathPrefix("/api").Subrouter()
	r.HandleFunc("/employee", handler.CreateHandler).Methods("POST")
	r.HandleFunc("/employee", handler.GetAllHandler).Methods("GET")
	r.HandleFunc("/employee/{id}", handler.UpdateHandler).Methods("PUT")
	r.HandleFunc("/employee/{id}", handler.DeleteHandler).Methods("DELETE")
	r.HandleFunc("/employee/{id}", handler.GetHandler).Methods("GET")

}
