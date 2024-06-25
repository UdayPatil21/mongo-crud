package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string             `json:"name"`
	Company    string             `json:"company"`
	Salary     float64            `json:"salary"`
	Experiance float64            `json:"experiance"`
}
