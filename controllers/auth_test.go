package controllers

import (
	"api/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// setupTestDB inicializa e limpa o banco de teste
func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	dsn := "host=localhost user=postgres password=123456 dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	db.Exec("DELETE FROM users") // limpa os usuários antes do teste
	return db
}

// TestRegister testa a criação de um novo usuário
//
// @Summary Registro de usuário
// @Description Testa o endpoint de registro
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body models.User true "Usuário"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
func TestRegister(t *testing.T) {
	db := setupTestDB(t)
	router := gin.Default()
	router.POST("/register", Register(db))

	payload := `{"name":"Test","email":"test@example.com","password":"123456"}`
	req := httptest.NewRequest("POST", "/register", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Esperado status 201, recebido %d", w.Code)
		var errResp map[string]interface{}
		json.Unmarshal(w.Body.Bytes(), &errResp)
		t.Logf("Resposta de erro: %v", errResp)
		return
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Erro ao decodificar resposta JSON: %v", err)
	}
	if _, ok := resp["message"]; !ok {
		t.Error("Esperado campo 'message' na resposta")
	}
}
