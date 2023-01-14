package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recruiter struct {
	ID         primitive.ObjectID
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	LastName   string `json:"LastName"`
	Gender     string `json:"Gender"`
	Phone      string `json:"Phone"`
	Email      string `json:"Email"`
	IsActive   bool   `json:"IsActive"`
	CreatedAt  primitive.DateTime
	UpdatedAt  primitive.DateTime
}
