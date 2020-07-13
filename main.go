package main

import (
	//"smarttask/services/firebaseClient"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// v1 := router.Group("/api/v1/todos")
	//  {
	//   v1.POST("/", createTodo)
	//   v1.GET("/", fetchAllTodo)
	//   v1.GET("/:id", fetchSingleTodo)
	//   v1.PUT("/:id", updateTodo)
	//   v1.DELETE("/:id", deleteTodo)
	//  }
	//  router.Run()

	router.GET("/", helloWorld)
	router.GET("/param/:id", ginURLParam)
	router.GET("/query", ginQuery)
	router.POST("/post", ginPostForm)

	router.Run()
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "hello world as a function"})
}

func ginURLParam(c *gin.Context) { // url parameters
	c.JSON(http.StatusOK, gin.H{"param": c.Param("id")})
}

func ginQuery(c *gin.Context) { // get query parameters
	id := c.Query("id")
	page := c.DefaultQuery("page", "No page")
	c.JSON(http.StatusOK, gin.H{"id": id, "page": page})
}

func ginPostForm(c *gin.Context) {
	name := c.PostForm("name")
	id := c.DefaultPostForm("id", "anonymous")
	c.JSON(http.StatusOK, gin.H{"name": name, "id ": id})
}

