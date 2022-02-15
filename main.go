package main

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO
	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("D:\\sqlite\\woaamazingprojekt.db"), &gorm.Config{})
	if err != nil {
		println(err)
	}
	db.AutoMigrate(&User{})

	// Add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&message{})

	r.GET("/ping", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/koel", func(c *gin.Context) {

	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
