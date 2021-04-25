package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"star-wars/helper"
	"star-wars/models"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) GetPlanets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var planets []models.Planet

	cur, err := s.db.Find(context.TODO(), bson.M{})
	if err != nil {
		helper.GetError(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var planet models.Planet

		err := cur.Decode(&planet)
		if err != nil {
			log.Fatal(err)
		}

		planets = append(planets, planet)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(planets)
}

func (s *Server) GetPlanetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var planet models.Planet
	var params = mux.Vars(r)
	if len(params["id"]) == 0 {
		helper.GetError(fmt.Errorf("No id params"), w)
		return
	}

	id, _ := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}
	err := s.db.FindOne(context.TODO(), filter).Decode(&planet)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(planet)
}

func (s *Server) GetPlanetByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var planet models.Planet
	var params = mux.Vars(r)

	filter := bson.M{"nome": params["nome"]}
	err := s.db.FindOne(context.TODO(), filter).Decode(&planet)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(planet)
}

func (s *Server) CreatePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var planet models.Planet

	err := json.NewDecoder(r.Body).Decode(&planet)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	result, err := s.db.InsertOne(context.TODO(), planet)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func (s *Server) DeletePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)
	if len(params["id"]) == 0 {
		helper.GetError(fmt.Errorf("No id params"), w)
		return
	}

	id, err := primitive.ObjectIDFromHex(params["id"])
	filter := bson.M{"_id": id}

	deleteResult, err := s.db.DeleteOne(context.TODO(), filter)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}
