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

	ref := client.Collection("tasks").NewDoc()
	input.OwnerId = "7vKfjbIhEemZlDdLoib7"
	input.RefId = ref.ID
	_, err := ref.Set(ctx, &input)

	if err != nil {
		log.Fatalf("Failed adding : %v", err)
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, fmt.Sprintf("Failed adding : %v", err), ""})

	} 

	log.Printf("Task added successfully, data: %v", &input)
	c.JSON(http.StatusOK, &models.AppResponse{0, fmt.Sprintf("Task Added successfully: ref: %v",ref.ID), &input})


}

func GetTaskInfo(c *gin.Context) {

	var input models.Task
	var taskInfo models.Task
	ownerId := "7vKfjbIhEemZlDdLoib7"

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, err.Error(), ""})
	}

	client, ctx := services.FirestoreClient()
	defer client.Close()
	//retrieve data

	dsnap, err := client.Collection("tasks").Doc(input.RefId)..Get(ctx)
	if err != nil {
		log.Fatalf("Failed fetching taskinfo : %v", err)
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, fmt.Sprintf("Failed fetching list : %v", err), ""})
	}
	
	dsnap.DataTo(&taskInfo)
	fmt.Printf("taskInfo fetched successfully, data =>  %#v\n", taskInfo)
	c.JSON(http.StatusOK, &models.AppResponse{0, fmt.Sprintf("TaskInfo Fetched Successfully"), taskInfo})

}

func GetTasks(c *gin.Context) {
	var lists []models.Task
	var list models.Task
	ownerId := "7vKfjbIhEemZlDdLoib7"

	client, ctx := services.FirestoreClient()
	defer client.Close()
	//retrieve data

	docs, err := client.Collection("tasks").Where("owner_id", "==", ownerId).Documents(ctx).GetAll()
	for _, doc := range docs {
		doc.DataTo(&list)
		lists = append(lists, list)
	}

	if err != nil {
		log.Fatalf("Failed fetching list : %v", err)
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, fmt.Sprintf("Failed fetching list : %v", err), ""})

	} 

	log.Printf("List fetched successfully, list => %v",&lists)
	c.JSON(http.StatusOK, &models.AppResponse{0, fmt.Sprintf("Task Fetched Successfully"), &lists})

}


func UpdateTask(c *gin.Context) {
	var input models.TaskUpdate

	input.OwnerId = "7vKfjbIhEemZlDdLoib7"

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, err.Error(), ""})
	}

	client, ctx := services.FirestoreClient()
	defer client.Close()

	_, err := client.Collection("tasks").Doc(input.RefId).Set(ctx, &input)

	if err != nil {
		log.Fatalf("Failed updating data : %v", err)
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, fmt.Sprintf("Failed updating task : %v", err), ""})

	}

	log.Println("Task updated successfully")
	c.JSON(http.StatusOK, &models.AppResponse{0, fmt.Sprintf("Task Updated Successfully"), &input})

}

func DeleteTask(c *gin.Context) {

	var input models.TaskDelete

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, err.Error(), ""})
	}

	client, ctx := services.FirestoreClient()
	defer client.Close()

	_, err := client.Collection("tasks").Doc(input.RefId).Delete(ctx)
	if err != nil {
		log.Fatalf("Failed deleteing data : %v", err)
		c.JSON(http.StatusBadRequest, &models.AppResponse{-1, fmt.Sprintf("Failed Deleting Data : %v", err), ""})

	}

	log.Println("Task Deleted successfully")
	c.JSON(http.StatusOK, &models.AppResponse{0, fmt.Sprintf("Task Deleted successfully: ref: %v",input.RefId), &input})
	

}