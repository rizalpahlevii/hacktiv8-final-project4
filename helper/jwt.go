package helper

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var SecretKey = []byte("secret")

func GenerateToken(id uint, email string, role string) string {
	claims := jwt.Claims(jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(SecretKey)
	PanicIfError(err)
	return signedToken
}

func VerifyToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}

	if time.Now().Unix() > int64(claims["exp"].(float64)) {
		return nil, fmt.Errorf("token expired")
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
