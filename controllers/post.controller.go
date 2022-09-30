package controllers

import (
	"context"
	"log"
	"net/http"
	"trandung/api/configs"
	"trandung/api/models"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var postCollection *mongo.Collection = configs.GetCollection(configs.DB, "posts")

func GetPosts(c echo.Context) error {
	cur, err := postCollection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())
	var results []models.Post
	if err = cur.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusOK, &echo.Map{"posts": results})
}
func CreatePost(c echo.Context) error {
	var newPost models.Post
	newPost.Id = primitive.NewObjectID()
	if err := c.Bind(&newPost); err != nil {
		return err
	}
	res, err := postCollection.InsertOne(context.TODO(), newPost)
	if err != nil {
		log.Fatal(err)
	}
	return c.JSON(http.StatusCreated, &echo.Map{
		"status": http.StatusCreated,
		"_id":    res.InsertedID,
	})
}
func GetPost(c echo.Context) error {
	var id = c.Param("id")
	IdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Id invalid",
		})
	}
	var post models.Post
	err = postCollection.FindOne(context.TODO(), bson.D{{"_id", IdObject}}).Decode(&post)
	if post.Title != "" && post.Description != "" && post.UserId.String() != "" {
		return c.JSON(http.StatusOK, post)
	} else {
		return c.JSON(http.StatusNotFound, &echo.Map{
			"status":  http.StatusNotFound,
			"message": "Post not found",
		})
	}

}
func DeletePost(c echo.Context) error {
	var id = c.Param("id")
	IdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Id invalid",
		})
	}
	var post models.Post
	err = postCollection.FindOne(context.TODO(), bson.D{{"_id", IdObject}}).Decode(&post)
	if post.Title == "" && post.Description == "" && post.UserId.String() == "" {
		return c.JSON(http.StatusNotFound, &echo.Map{
			"status":  http.StatusNotFound,
			"message": "Post not found",
		})
	} else {
		var rs models.Post
		err = postCollection.FindOneAndDelete(context.TODO(), bson.D{{"_id", IdObject}}).Decode(&rs)
		return c.JSON(http.StatusOK, &echo.Map{
			"status":  http.StatusOK,
			"message": "Deleted post successfully!",
		})
	}

}

func UpdatePost(c echo.Context) error {
	var id = c.Param("id")
	IdObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &echo.Map{
			"status":  http.StatusBadRequest,
			"message": "Id invalid",
		})
	}
	var updatePost models.Post
	if err := c.Bind(&updatePost); err != nil {
		return err
	}
	var post models.Post
	err = postCollection.FindOne(context.TODO(), bson.D{{"_id", IdObject}}).Decode(&post)
	if post.Title != "" && post.Description != "" && post.UserId.String() != "" {
		filter := bson.D{{"_id", IdObject}}
		update := bson.D{{"$set", bson.D{{"userId", updatePost.UserId}, {"title", updatePost.Title}, {"description", updatePost.Description}}}}
		result, err := postCollection.UpdateOne(context.TODO(), filter, update)
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
				"message": "Post update failure!",
			})
		}
	} else {
		return c.JSON(http.StatusNotFound, &echo.Map{
			"status":  http.StatusNotFound,
			"message": "Post not found",
		})
	}

}
