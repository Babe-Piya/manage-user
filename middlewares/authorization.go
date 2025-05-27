package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type Authorization interface {
	AuthorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type authorization struct {
	JwtSecret []byte
}

func NewAuthorization(jwtSecret string) Authorization {
	return &authorization{
		JwtSecret: []byte(jwtSecret),
	}
}

type Claims struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func (auth *authorization) AuthorizationMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": http.StatusText(http.StatusUnauthorized),
			})
		}
		if !strings.HasPrefix(authHeader, "Bearer ") {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": http.StatusText(http.StatusUnauthorized),
			})
		}

		accessToken := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := validateJWT(accessToken, auth.JwtSecret)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": http.StatusText(http.StatusUnauthorized),
			})
		}

		c.Set("email", claims.Email)
		c.Set("name", claims.Name)
		c.Set("id", claims.ID)

		return next(c)
	}
}

func validateJWT(tokenString string, jwtSecret []byte) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
