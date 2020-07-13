package main

import (
	//"smarttask/services/firebaseClient"
	"net/http"
	"smarttask/models"
	"smarttask/controllers"
	"github.com/gin-gonic/gin"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-contrib/cors"
)
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, &models.AppResponse{0, "Welcome to smarttask api ", ""})
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())

	// use nice.Recovery
	router.Use(nice.Recovery(recoveryHandler))

	// Set CORS allow all origins from parameter
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Origin", "Api-Key", "Content-Type", "X-CSRF-TOKEN"}
	router.Use(cors.New(corsConfig))

	router.GET("/", healthCheck)

	v1 := router.Group("/api/v1")
	v1.GET("/health", healthCheck)

	authRoute := v1.Group("/auth")
	authRoute.POST("/register", controllers.RegisterUser)
	authRoute.POST("/login", controllers.LoginUser)

	taskRoute := v1.Group("/task")
	taskRoute.POST("/add/", controllers.AddTask)
	taskRoute.POST("/list/", controllers.GetTasks)
	taskRoute.POST("/view/:ref_id", controllers.GetTaskInfo)
	taskRoute.POST("/update", controllers.UpdateTask)
	taskRoute.POST("/delete", controllers.DeleteTask)

	profileRoute := v1.Group("/profile")
	profileRoute.POST("/view", controllers.GetUserInfo)
	profileRoute.POST("/update", controllers.UpdateUserInfo)
	profileRoute.POST("/password", controllers.UpdateUserPassword)

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

