package services

import (
	"context"
	"log"
	"time"

	"manage-user/appconstants"
	"manage-user/middlewares"
	"manage-user/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

func (srv *userService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	users, err := srv.UserRepo.GetUserByFilter(ctx, repositories.User{
		Email: req.Email,
	})
	if err != nil {
		log.Println(err)

		return nil, err
	}

	if !checkPassword(users[0].Password, req.Password) {
		log.Println("can not login wrong email or password")

		return nil, appconstants.WrongKeyLoginError
	}

	token, err := srv.generateToken(users[0].ID, users[0].Name, users[0].Email)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	return &LoginResponse{
		Code:    appconstants.SuccessCode,
		Message: appconstants.SuccessMessage,
		Token:   token,
	}, nil
}

func checkPassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}

	return true
}

func (srv *userService) generateToken(id, name, email string) (string, error) {
	now := time.Now()
	expiredTime := now.Add(srv.Config.TokenTime * time.Minute)
	claims := &middlewares.Claims{
		ID:    id,
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
			IssuedAt:  jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(srv.Config.JwtSecret))
}
