package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

type AccessClaims struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	jwt.RegisteredClaims
}
type refreshClaims struct {
	jwt.RegisteredClaims
}

var (
	accessSecret  = []byte(os.Getenv("JWT_ACCESS_SECRET"))
	refreshSecret = []byte(os.Getenv("JWT_REFRESH_SECRET"))
)

func CreateJWTTokens(id uuid.UUID, email string, access bool) (*string, error) {
	if access {
		return createAccessToken(id, email)
	} else {
		return createRefreshToken()
	}
}

func createAccessToken(id uuid.UUID, email string) (*string, error) {
	claims := &AccessClaims{
		id,
		email,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	at, err := token.SignedString(accessSecret)
	if err != nil {
		return nil, fmt.Errorf("Access Token 발행에 실패했습니다. :: %v", err)
	}

	return &at, nil
}

func createRefreshToken() (*string, error) {
	claims := &refreshClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	rt, err := token.SignedString(refreshSecret)
	if err != nil {
		return nil, fmt.Errorf("Refresh Token 발행에 실패했습니다. :: %v", err)
	}
	return &rt, nil
}

func ParseAccessToken(tokenString string) (*AccessClaims, error) {
	claims := new(AccessClaims)
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		if t.Method.Alg() != jwt.SigningMethodES256.Alg() {
			return nil, fmt.Errorf("unexpected signing method :: %v", t.Header["alg"])
		}
		return accessSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token :: %v", err)
	}
	return claims, nil
}

func EchoJwtConfig() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			return new(AccessClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_ACCESS_SECRET")),
	}
}
