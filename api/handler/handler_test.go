package handler

import (
	"bytes"
	"encoding/json"
	"mongo-crud/internal/database"
	"mongo-crud/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Employee struct {
	Name       string  `json:"name"`
	Company    string  `json:"company"`
	Salary     float64 `json:"salary"`
	Experiance float64 `json:"experiance"`
}

func TestCreateHandler(t *testing.T) {
	mockCollection := new(database.MockCollection)
	app := &App{Collection: mockCollection}

	employee := model.Employee{ID: primitive.NewObjectID(), Name: "Uday", Company: "gslab", Salary: 15, Experiance: 3}
	jsonStr, _ := json.Marshal(employee)
	req, err := http.NewRequest("POST", "/api/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	// setup mock response
	// mockRes := mongo.InsertOneResult{InsertedID: employee.ID}
	mockCollection.On("InsertOne", mock.Anything, mock.Anything, ([]*options.InsertOneOptions)(nil)).Return(&mongo.InsertOneResult{}, nil)

	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/", app.CreateHandler)
	r.ServeHTTP(rr, req)

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler return wrong status code")
	mockCollection.AssertExpectations(t)
}
func TestGetHandler(t *testing.T) {
	mockCollection := new(database.MockCollection)
	app := &App{Collection: mockCollection}

	emp := model.Employee{ID: primitive.NewObjectID(), Name: "Uday", Company: "gslab", Salary: 10, Experiance: 3}

	//setup mock response
	mockRes := mongo.NewSingleResultFromDocument(emp, nil, nil)
	mockCollection.On("FindOne", mock.Anything, bson.D{{"_id", emp.ID}}, ([]*options.FindOneOptions)(nil)).Return(mockRes)

	req, err := http.NewRequest("GET", "/api/"+emp.ID.Hex(), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/{id}", app.GetHandler)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler return wrong status code")
	mockCollection.AssertExpectations(t)
}

func TestDeleteHandler(t *testing.T) {
	mockCollection := new(database.MockCollection)
	app := &App{Collection: mockCollection}

	emp := model.Employee{ID: primitive.NewObjectID(), Name: "Uday", Company: "gslab", Salary: 10, Experiance: 3}

	//setup mock response
	// mockRes := mongo.DeleteResult{DeletedCount: 1}
	mockCollection.On("DeleteOne", mock.Anything, bson.D{{"_id", emp.ID}}, ([]*options.DeleteOptions)(nil)).Return(&mongo.DeleteResult{DeletedCount: 1})

	req, err := http.NewRequest("DELETE", "/api/"+emp.ID.Hex(), nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/{id}", app.DeleteHandler)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler return wrong status code")
	mockCollection.AssertExpectations(t)
}
func TestUpdateHandler(t *testing.T) {
	mockCollection := new(database.MockCollection)
	app := &App{Collection: mockCollection}

	emp := model.Employee{ID: primitive.NewObjectID(), Name: "Uday", Company: "gslab", Salary: 10, Experiance: 3}
	jsonStr, _ := json.Marshal(emp)
	mockCollection.On("UpdateOne", mock.Anything, mock.Anything, ([]*options.UpdateOptions)(nil)).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil)
	req, err := http.NewRequest("PUT", "/api/", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/", app.UpdateHandler)
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "handler return wrong status code")
	mockCollection.AssertExpectations(t)
}
