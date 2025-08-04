package controllers

import (
	"api/models"
	"api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// RegisterRequest define os dados esperados para registro
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest define os dados esperados para login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register cria um novo usuário
// @Summary Registro de usuário
// @Description Cria um novo usuário com nome, email e senha
// @Tags auth
// @Accept json
// @Produce json
// @Param user body RegisterRequest true "Dados do usuário"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RegisterRequest
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

		token := "dummy-verification-token"
		utils.SendVerificationEmail(user.Email, token)
		c.JSON(201, gin.H{"message": "Registrado. Cheque seu email."})
	}
}

// Login autentica um usuário existente
// @Summary Login
// @Description Autentica o usuário e retorna um token JWT
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body LoginRequest true "Credenciais"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 403 {object} map[string]string
// @Router /login [post]
func Login(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req LoginRequest
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

// VerifyEmail verifica o email com um token fictício
// @Summary Verificar email
// @Description Simula a verificação do email com um token
// @Tags auth
// @Produce json
// @Param token path string true "Token de verificação"
// @Success 200 {object} map[string]string
// @Router /verify-email/{token} [get]
func VerifyEmail(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		// Aqui você poderia validar o token de verdade no banco
		c.JSON(200, gin.H{"message": "Email verificado com token: " + token})
	}
}
