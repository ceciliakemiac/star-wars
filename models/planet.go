package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nome      string             `json:"nome,omitempty" bson:"nome,omitempty"`
	Clima     string             `json:"clima" bson:"clima,omitempty"`
	Terreno   string             `json:"terreno" bson:"terreno,omitempty"`
	Aparicoes int                `json:"aparicoes" bson:"aparicoes,omitempty"`
}
