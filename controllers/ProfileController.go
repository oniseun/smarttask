package controllers

import (
	"smarttask/services/firebaseClient"
	"github.com/gin-gonic/gin"
	"smarttask/models"
)

func GetUserInfo(c *gin.Context) {
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

func UpdateUserInfo(c *gin.Context) {
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

func UpdateUserPassword(c *gin.Context) {
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
