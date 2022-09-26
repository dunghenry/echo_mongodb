package controllers

import (
	"context"
	"log"
	"net/http"
	"trandung/api/configs"
	"trandung/api/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func GetUsers(c echo.Context) error {
	// user := models.User{
	// 	Name:  "Jon",
	// 	Email: "jon@labstack.com",
	// }
	cur, err := userCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var results []models.User
	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, &echo.Map{"users": results})
}
