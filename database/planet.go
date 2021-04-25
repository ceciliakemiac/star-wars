package database

import (
	"context"
	"fmt"
	"os"
	"star-wars/models"
	"strconv"

	"github.com/peterhellberg/swapi"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertPlanets(db *mongo.Collection) {
	c := swapi.DefaultClient

	totalPlanets, _ := strconv.Atoi(os.Getenv("NUMBER_OF_PLANETS"))

	for i := 1; i <= totalPlanets; i++ {
		fmt.Println("Inserting planets...")

		if planet, err := c.Planet(i); err == nil {
			planetStruct := models.Planet{
				Nome:      planet.Name,
				Clima:     planet.Climate,
				Terreno:   planet.Terrain,
				Aparicoes: len(planet.FilmURLs),
			}

			_, err := db.InsertOne(context.TODO(), planetStruct)
			if err != nil {
				fmt.Println("Error inserting " + planetStruct.Nome)
			}
		}
	}
}
