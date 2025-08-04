package utils

import (
    "time"
    "github.com/golang-jwt/jwt/v5"
    "api/models"
)

func GenerateJWT(user models.User) (string, error) {
    claims := jwt.MapClaims{
        "sub":  user.ID,
        "role": user.Role,
        "exp":  time.Now().Add(time.Hour * 24).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte("SECRET_KEY"))
}
