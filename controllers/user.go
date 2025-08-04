package controllers

import (
	"api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Profile retorna o perfil do usuário autenticado
// @Summary Ver perfil
// @Description Retorna os dados do usuário autenticado via JWT
// @Tags user
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /profile [get]
func Profile(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt("user_id")
		var user models.User
		if err := db.First(&user, userID).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(200, gin.H{"user": user})
	}
}

// GetUser retorna um usuário por ID
// @Summary Buscar usuário por ID
// @Description Retorna os dados de um usuário específico
// @Tags user
// @Produce json
// @Param id path int true "ID do usuário"
// @Success 200 {object} map[string]interface{}
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /users/{id} [get]
func GetUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var user models.User
		if err := db.First(&user, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(200, gin.H{"user": user})
	}
}

// AdminDashboard retorna informações administrativas
// @Summary Dashboard do administrador
// @Description Retorna uma mensagem de boas-vindas ao admin
// @Tags admin
// @Produce json
// @Success 200 {object} map[string]string
// @Security BearerAuth
// @Router /admin/dashboard [get]
func AdminDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Admin dashboard"})
	}
}
