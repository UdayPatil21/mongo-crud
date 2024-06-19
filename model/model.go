package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string
	Company    string
	Salary     float64
	Experiance float64
}
