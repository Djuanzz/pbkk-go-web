package main

import (
	"github.com/Djuanzz/pbkk-go-web/config"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = config.ConnectDatabase()

}
