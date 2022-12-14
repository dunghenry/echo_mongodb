package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"trandung/api/configs"
	"trandung/api/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")

func GetUsers(c echo.Context) error {
	data := c.Request().Header.Get("token")
	fmt.Println(data)
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
func CreateUser(c echo.Context) error {
	var newUser models.User
	newUser.Id = primitive.NewObjectID()
	if err := c.Bind(&newUser); err != nil {
		return err
	}
	res, err := userCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, &echo.Map{
		"status": http.StatusCreated,
		"_id":    res.InsertedID,
	})
}
func GetUser(c echo.Context) error {
	var id = c.Param("id")
	IdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Id invalid",
		})
	}
	var user models.User
	err = userCollection.FindOne(context.TODO(), bson.D{{"_id", IdObject}}).Decode(&user)
	if user.Name != "" && user.Email != "" {
		return c.JSON(http.StatusOK, user)
	} else {
		return c.JSON(http.StatusNotFound, &echo.Map{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
	}

}
func DeleteUser(c echo.Context) error {
	var id = c.Param("id")
	IdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Id invalid",
		})
	}
	var user models.User
	err = userCollection.FindOne(context.TODO(), bson.D{{"_id", IdObject}}).Decode(&user)
	if user.Name == "" && user.Email == "" {
		return c.JSON(http.StatusNotFound, &echo.Map{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
	} else {
		var rs models.User
		err = userCollection.FindOneAndDelete(context.TODO(), bson.D{{"_id", IdObject}}).Decode(&rs)
		return c.JSON(http.StatusOK, &echo.Map{
			"status":  http.StatusOK,
			"message": "Deleted user successfully!",
		})
	}

}

func UpdateUser(c echo.Context) error {
	var id = c.Param("id")
	IdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Id invalid",
		})
	}
	var updateUser models.User
	if err := c.Bind(&updateUser); err != nil {
		return err
	}
	var user models.User
	err = userCollection.FindOne(context.TODO(), bson.D{{"_id", IdObject}}).Decode(&user)
	if user.Email != "" && user.Name != "" {
		filter := bson.D{{"_id", IdObject}}
		update := bson.D{{"$set", bson.D{{"name", updateUser.Name}, {"email", updateUser.Email}}}}
		result, err := userCollection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Fatal(err)
		}
		if result.MatchedCount != 0 {
			return c.JSON(http.StatusOK, &echo.Map{
				"status": http.StatusOK,
				"item":   "Updated todo successfully!",
			})

		} else {
			return c.JSON(http.StatusBadRequest, &echo.Map{
				"status":  http.StatusBadRequest,
				"message": "Todo update failure!",
			})
		}
	} else {
		return c.JSON(http.StatusNotFound, &echo.Map{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
	}

}
