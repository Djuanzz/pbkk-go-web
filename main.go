package main

import (
	"os"

	"github.com/Djuanzz/pbkk-go-web/config"
	"github.com/Djuanzz/pbkk-go-web/controller"
	"github.com/Djuanzz/pbkk-go-web/model"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = config.ConnectDatabase()

	model.Migration(db)

	router := gin.Default()

	router.LoadHTMLGlob("page/*")

	userController := controller.NewUserController(db)

	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.GET("/users", userController.GetUsersHTML)
	router.GET("/users/add", userController.ShowAddUserForm)
	router.POST("/users/add", userController.AddUser)

	router.POST("/api/user", userController.CreateUser)
	router.GET("/api/user", userController.GetUsers)
	router.DELETE("/api/user/:id", userController.DeleteUser)
	router.PATCH("/api/user/:id", userController.UpdateUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if err := router.Run(":" + port); err != nil {
		panic(err.Error())
	}

}
