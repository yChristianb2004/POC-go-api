package main

import (
	"api/models"
	"api/routes"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "api/docs" // Documentação gerada (swag init vai criar)

	swaggerFiles "github.com/swaggo/files"     // Swagger UI files
	ginSwagger "github.com/swaggo/gin-swagger" // Swagger handler
)

// @title API com Swagger - Exemplo
// @version 1.0
// @description API exemplo com documentação Swagger em Go + Gin
// @host localhost:8080
// @BasePath /

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

	// Endpoint do Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r, db)
	r.Run()
}
