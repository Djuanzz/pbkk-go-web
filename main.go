package main

import (
	"fmt"

	"github.com/Djuanzz/pbkk-go-web/config"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

var db *gorm.DB

type Album struct {
	ID     int
	Title  string
	Artist string
	Price  float32
}

func addAlbum(alb Album) (int, error) {
	result := db.Create(&alb)
	if result.Error != nil {
		return 0, result.Error
	}
	return alb.ID, nil
}

func albumByID(id int) (Album, error) {
	var alb Album
	result := db.First(&alb, id)
	if result.Error != nil {
		return alb, result.Error
	}
	return alb, nil
}

func main() {
	db = config.ConnectDatabase()

	// Add an album to the database.
	// alb := Album{Title: "The Modern Sound of Betty Carter", Artist: "Betty Carter", Price: 49.95}
	// id, err := addAlbum(alb)
	// if err != nil {
	// 	panic(err)
	// }
	// println("New album ID is", id)

	// Retrieve an album from the database.
	alb, err := albumByID(2)
	if err != nil {
		panic(err)
	}
	fmt.Println("Album is", alb)

}
