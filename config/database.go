package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// func ConnectDb() *sql.DB {
// 	cfg := mysql.Config{
// 		User:   os.Getenv("DB_USER"),
// 		Passwd: os.Getenv("DB_PASS"),
// 		Net:    "tcp",
// 		Addr:   os.Getenv("DB_HOST"),
// 		DBName: os.Getenv("DB_NAME"),
// 	}

// 	db, err := sql.Open("mysql", cfg.FormatDSN())
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	pingErr := db.Ping()
// 	if pingErr != nil {
// 		panic(pingErr.Error())
// 	}

// 	fmt.Println("Connected to database")

// 	return db
// }

func ConnectDatabase() *gorm.DB {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
