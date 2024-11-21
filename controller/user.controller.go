package controller

import (
	"net/http"

	"github.com/Djuanzz/pbkk-go-web/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success create user", "data": user})
}

func (uc *UserController) GetUsers(c *gin.Context) {
	var users []model.User

	if err := uc.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (uc *UserController) GetUsersHTML(c *gin.Context) {
	var users []model.User

	if err := uc.DB.Find(&users).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "users.html", gin.H{
			"error": "Failed to fetch users",
		})
		return
	}

	c.HTML(http.StatusOK, "users.html", gin.H{"data": users})
}

func (uc *UserController) ShowAddUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "addUser.html", nil)
}

func (uc *UserController) AddUser(c *gin.Context) {
	var user model.User

	user.Email = c.PostForm("email")
	user.Password = c.PostForm("password")
	user.Role = c.PostForm("role")

	if err := uc.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.HTML(http.StatusOK, "addUser.html", gin.H{
		"message": "User successfully added!",
	})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	if err := uc.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := uc.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success delete user", "data": user})
}

func (uc *UserController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user model.User

	if err := uc.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := uc.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success update user", "data": user})
}
