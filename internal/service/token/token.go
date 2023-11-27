package token

import (
	"fmt"
	"time"

	"github.com/bifidokk/go-api/internal/entity"
	"github.com/golang-jwt/jwt/v5"
)

type JwtCustomClaims struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string, secret string) (*JwtCustomClaims, error) {
	claims := &JwtCustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token is invalid")
	}

	claims, ok := token.Claims.(*JwtCustomClaims)

	if !ok {
		return nil, fmt.Errorf("Token is valid but couldn't parse claims")
	}

	return claims, nil
}

func CreatAccessToken(user *entity.User, secret string, ttlSeconds int) (accessToken string, err error) {
	expiresAt := time.Now().Add(time.Second * time.Duration(ttlSeconds))

	claims := &JwtCustomClaims{
		Email: user.Email,
		ID:    user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return signedToken, err
}
