package utils

import (
	"time"

	"github.com/belovetech/e-commerce/database/sqlc"
	"github.com/golang-jwt/jwt"
)

var jwtSecretKey []byte

func SetJWTSecretKey(secretKey string) {
	jwtSecretKey = []byte(secretKey)
}

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(user sqlc.GetUserByEmailRow) (string, error) {
	claims := Claims{
		ID:    int(user.ID),
		Email: user.Email,
		Role:  user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func VerifyJWT(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}
