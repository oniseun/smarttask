package main

import (
	//"smarttask/services/firebaseClient"
	"fmt"
	"log"
	"net/http"
	"os"
	"smarttask/services"
	"strconv"

	"smarttask/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	router.GET("/", helloWorld)
	router.POST("/add", ginAddForm)
	router.POST("/update", ginUpdateForm)
	router.GET("/list", ginList)

	router.Run()
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, &models.AppResponse{0, "Welcome to our api ", ""})
}

func getUserId() int {
	user, _ := os.LookupEnv("DEFAULT_USER")
	userID, _ := strconv.Atoi(user)
	return userID
}

type Data struct {
	OwnerId    int    `json:"ownerId"`
	RefId string `json:"refId"`
	Born  int    `json:"born"`
	First string `json:"first"`
	Last  string `json:"last"`
}

func ginAddForm(c *gin.Context) {

	var input Data
	
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, err.Error(), ""})
	}

	client, ctx := services.FirestoreClient()
	defer client.Close()

	ref := client.Collection("users").NewDoc()
	input.OwnerId = getUserId()
	input.RefId = ref.ID
	_, err := ref.Set(ctx, &input)

	if err != nil {
		log.Fatalf("Failed adding : %v", err)
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, fmt.Sprintf("Failed adding : %v", err), ""})

	} else {
		log.Println("data added successfully")
		c.JSON(http.StatusOK, &models.AppResponse{0, "Data added successfully", &input})
	}

}

type UpdateData struct {
	Born  int    `json:"born"`
	First string `json:"first"`
	Last  string `json:"last"`
}

func ginUpdateForm(c *gin.Context) {
	var input Data
	input.OwnerId = getUserId()
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": err.Error()})
		return
	}

	client, ctx := services.FirestoreClient()
	defer client.Close()

	_, err := client.Collection("users").Doc("7vKfjbIhEemZlDdLoib7").Set(ctx, &input)

	if err != nil {
		log.Fatalf("Failed updating data : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": fmt.Sprintf("Failed adding : %v", err)})

	} else {
		log.Println("data updated successfully")
		c.JSON(http.StatusOK, gin.H{"success": 1, "data": &input})
	}
}

func ginList(c *gin.Context) {
	var lists []Data
	var list Data

	client, ctx := services.FirestoreClient()
	defer client.Close()
	//retrieve data

	docs, err := client.Collection("users").Limit(10).Documents(ctx).GetAll()
	for _, doc := range docs {
		doc.DataTo(&list)
		lists = append(lists, list)
	}

	if err != nil {
		log.Fatalf("Failed fetching list : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": 0, "message": fmt.Sprintf("Failed fetching list : %v", err)})

	} else {
		log.Println("List fetched successfully")
		c.JSON(http.StatusOK, gin.H{"success": 1, "data": &lists})
	}

}
