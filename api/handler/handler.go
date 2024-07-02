// Package handler Employee API.
//
// the purpose of this application is to provide an application
// that is using go code to define an  Rest API
//
//	Schemes: http, https
//	Host: localhost:8080
//	BasePath: /api
//	Version: 0.0.1
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"mongo-crud/internal/database"
	"mongo-crud/model"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type App struct {
	Collection database.Collection
}

func NewApp(collection database.Collection) *App {
	return &App{Collection: collection}
}

// swagger:operation POST /employee Employee CreateHandler
//
// # Add new Employee
//
// # Returns new Employee
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
//   - name: employee
//     in: body
//     description: add employee data
//     required: true
//     schema:
//     "$ref": "#/definitions/Employee"
//
// responses:
//
//	'200':
//	  description: Employee response
//	  schema:
//	    "$ref": "#/definitions/Employee"
//	'409':
//	  description: Conflict
//	'405':
//	  description: Method Not Allowed, likely url is not correct
//	'403':
//	  description: Forbidden, you are not allowed to undertake this operation
//
// CreateHandler: insert employee data into the database
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
	emp.ID = primitive.NewObjectID()
	res, err := a.Collection.InsertOne(ctx, emp)
	if err != nil {
		response.WriteHeader(http.StatusExpectationFailed)
		response.Write([]byte(err.Error()))
		return
	}
	response.WriteHeader(http.StatusOK)
	resByte, _ := json.Marshal(res)
	response.Write([]byte(string(resByte)))

}

// swagger:operation GET /employee Employee GetAllHandler
//
// Get Employee
//
// Returns existing Employees
//
// ---
// produces:
// - application/json
// responses:
//   '200':
//     description: employee data
//     schema:
//      $ref: "#/definitions/Employee"
//   '405':
//     description: Method Not Allowed, likely url is not correct
//   '403':
//     description: Forbidden, you are not allowed to undertake this operation

// GetAllHandler: fetch employee data from database
func (a *App) GetAllHandler(response http.ResponseWriter, request *http.Request) {
	var empData []model.Employee
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := a.Collection.Find(ctx, bson.D{})
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

// swagger:operation GET /employee/{id} Employee GetHandler
//
// # Get Employee
//
// # Returns existing Employee filtered by fname
//
// ---
// produces:
// - application/json
// parameters:
//   - name: fname
//     type: string
//     in: path
//     required: true
//
// responses:
//
//	'200':
//	  description: employee data
//	  schema:
//	   "$ref": "#/definitions/Employee"
//	'405':
//	  description: Method Not Allowed, likely url is not correct
//	'403':
//	  description: Forbidden, you are not allowed to undertake this operation
//
// GetHandler: fetch employee by id
func (a *App) GetHandler(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var emp model.Employee
	_ = a.Collection.FindOne(ctx, bson.D{{"_id", id}}).Decode(&emp)
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

// swagger:operation PUT /employee/{id} Employee UpdateHandler
//
// Update Employee
//
// Update existing Employee filtered by id
//
// ---
// consumes:
// - application/json
// produces:
// - application/json
// parameters:
// - name: fname
//   type: string
//   in: path
//   required: true
// - name: employee
//   in: body
//   description: add employee data
//   required: true
//   schema:
//     "$ref": "#/definitions/Employee"
// responses:
//   '200':
//     description: Employee response
//     schema:
//       "$ref": "#/definitions/Employee"
//   '409':
//     description: Conflict
//   '405':
//     description: Method Not Allowed, likely url is not correct
//   '403':
//     description: Forbidden, you are not allowed to undertake this operation

// UpdateHandler: Update employee data by id
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
	res, err := a.Collection.UpdateOne(ctx, bson.D{{"_id", id}}, update)
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

// swagger:operation DELETE /employee/{id} Employee DeleteHandler
//
// Delete Employee
//
// Delete existing Employee filtered by id
//
// ---
// produces:
// - application/json
// parameters:
//  - name: fname
//    type: string
//    in: path
//    required: true
// responses:
//   '200':
//     description: delete employee sucessfully
//     schema:
//       "$ref": "#/definitions/Employee"
//   '405':
//     description: Method Not Allowed, likely url is not correct
//   '403':
//     description: Forbidden, you are not allowed to undertake this operation

// DeleteHandler: Delete employee by id
func (a *App) DeleteHandler(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := a.Collection.DeleteOne(ctx, bson.D{{"_id", id}})
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
