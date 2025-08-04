package main

import (
    "os"
    "api/models"
    "api/routes"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/gin-gonic/gin"
)

func InitDB() *gorm.DB {
    dsn := os.Getenv("DATABASE_DSN")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
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
