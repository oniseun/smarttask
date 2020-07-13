package controllers

import (
	"smarttask/services/firebaseClient"
	"github.com/gin-gonic/gin"
	"smarttask/models"
)

func AddTask(c *gin.Context) {

	var input models.Task
	
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

func GetTaskInfo(c *gin.Context) {
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

func GetTasks(c *gin.Context) {
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


func UpdateTask(c *gin.Context) {
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

func DeleteTask(c *gin.Context) {
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