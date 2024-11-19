package main

import (
	"net/http"
	"os"

	"github.com/Djuanzz/pbkk-go-web/config"
	"github.com/Djuanzz/pbkk-go-web/model"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

var db *gorm.DB

func createUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success create user", "data": user})
}

func getUsers(c *gin.Context) {
	var users []model.User

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success delete user", "data": user})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success update user", "data": user})
}

func main() {
	db = config.ConnectDatabase()

	model.Migration(db)

	router := gin.Default()
	router.GET("/api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	router.POST("/api/user", createUser)
	router.GET("/api/user", getUsers)
	router.DELETE("/api/user/:id", deleteUser)
	router.PATCH("/api/user/:id", updateUser)

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if err := router.Run(":" + port); err != nil {
		panic(err.Error())
	}

}
