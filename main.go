package main

import (
	"api/models"
	"api/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=123456 dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	db.AutoMigrate(&models.User{})
	return db
}

func main() {
	db := InitDB()
	r := gin.Default()
	routes.SetupRoutes(r, db)
	r.Run()
}
