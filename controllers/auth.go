package controllers

import (
	"api/models"
	"api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Requisição inválida"})
			return
		}
		hash, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user := models.User{
			Name:     req.Name,
			Email:    req.Email,
			Password: string(hash),
			Role:     "client",
		}
		if err := db.Create(&user).Error; err != nil {
			c.JSON(400, gin.H{"error": "Email já registrado"})
			return
		}

		// Simular token de verificação
		token := "dummy-verification-token"
		utils.SendVerificationEmail(user.Email, token)
		c.JSON(201, gin.H{"message": "Registrado. Cheque seu email."})
	}
}

func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "Requisição inválida"})
			return
		}
		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			c.JSON(401, gin.H{"error": "Credenciais inválidas"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			c.JSON(401, gin.H{"error": "Credenciais inválidas"})
			return
		}
		if !user.IsEmailVerified {
			c.JSON(403, gin.H{"error": "Email não verificado"})
			return
		}
		token, _ := utils.GenerateJWT(user)
		c.JSON(200, gin.H{"token": token})
	}
}

func VerifyEmail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simular verificação
		c.JSON(200, gin.H{"message": "Email verificado (POC)"})
	}
}
