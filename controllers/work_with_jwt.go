package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"referral_system/config"
	"referral_system/models"
	"time"
)

var secretKey = []byte("123") // Так не делаем, лучше в переменное окружение закинуть

func GenerateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func VerifyJWT(c fiber.Ctx) (models.User, error) {
	tokenString := c.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil || !token.Valid {
		return models.User{}, err
	}
	claims := token.Claims.(jwt.MapClaims)
	userId := claims["sub"].(string)
	var user models.User
	if err := config.DB.First(&user, userId).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
