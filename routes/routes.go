package routes

import (
    "api/controllers"
    "api/middlewares"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
    r.POST("/register", controllers.Register(db))
    r.POST("/login", controllers.Login(db))
    r.GET("/verify-email/:token", controllers.VerifyEmail(db))

    auth := r.Group("/")
    auth.Use(middlewares.AuthMiddleware())
    auth.GET("/profile", controllers.Profile(db))
    auth.GET("/users/:id", middlewares.RoleMiddleware("admin", "user"), controllers.GetUser(db))
    auth.GET("/admin/dashboard", middlewares.RoleMiddleware("admin"), controllers.AdminDashboard())
}
