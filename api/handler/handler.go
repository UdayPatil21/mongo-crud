package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"mongo-crud/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	DB *mongo.Client
}

func NewApp(db *mongo.Client) *App {
	return &App{DB: db}
}

// Insert employee data into the database
func (a *App) CreateHandler(response http.ResponseWriter, request *http.Request) {
	var emp = model.Employee{}
	reqData, _ := io.ReadAll(request.Body)
	err := json.Unmarshal(reqData, &emp)
	if err != nil {
		log.Println("Error binding request data")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := a.DB.Database("gslab").Collection("employee").InsertOne(ctx, emp)
	if err != nil {
		response.WriteHeader(http.StatusExpectationFailed)
		response.Write([]byte(err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	resByte, _ := json.Marshal(res)
	response.Write([]byte(string(resByte)))

}

// Get all employee data from database
func (a *App) GetAllHandler(response http.ResponseWriter, request *http.Request) {
	var empData []model.Employee
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := a.DB.Database("gslab").Collection("employee").Find(ctx, bson.D{})
	if err != nil {
		response.WriteHeader(http.StatusExpectationFailed)
		response.Write([]byte(err.Error()))
		return
	}
	err = res.All(ctx, &empData)
	if err != nil {
		response.WriteHeader(http.StatusExpectationFailed)
		response.Write([]byte(err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	resByte, _ := json.Marshal(empData)
	response.Write(resByte)
}

// Get employee by id
func (a *App) GetHandler(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var emp model.Employee
	_ = a.DB.Database("gslab").Collection("employee").FindOne(ctx, bson.D{{"_id", id}}).Decode(&emp)
	response.WriteHeader(http.StatusOK)
	resByte, _ := json.Marshal(emp)
	response.Write(resByte)
}

type employee struct {
	Name       string
	Company    string
	Salary     float64
	Experiance float64
}

// Update employee data by id
func (a *App) UpdateHandler(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var emp = employee{}
	err := json.NewDecoder(request.Body).Decode(&emp)
	if err != nil {
		log.Println("Error binding request data")
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.M{"$set": emp}
	res, err := a.DB.Database("gslab").Collection("employee").UpdateOne(ctx, bson.D{{"_id", id}}, update)
	if err != nil {
		response.WriteHeader(http.StatusExpectationFailed)
		response.Write([]byte(err.Error()))
		return
	}
	if res.MatchedCount == 0 && res.ModifiedCount == 0 {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("No Document Found"))
		return
	}
	response.WriteHeader(http.StatusOK)
	response.Write([]byte("Successfuly Updated"))
}

// Delete employee by id
func (a *App) DeleteHandler(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := a.DB.Database("gslab").Collection("employee").DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		response.WriteHeader(http.StatusExpectationFailed)
		response.Write([]byte(err.Error()))
		return
	}
	if res.DeletedCount == 0 {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte("No Document Found"))
		return
	}
	response.WriteHeader(http.StatusOK)
	resByte, _ := json.Marshal(res)
	response.Write(resByte)
}
