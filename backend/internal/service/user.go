package service

import (
	"backend/internal/entity"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func (s *userImpl) CreateUser(
	ctx context.Context,
	data CreateUserBody,
) (*LoginResponse, error) {

	user, err := s.repo.CreateUser(entity.User{
		ID:    uuid.New().String(),
		Email: data.Email,
		Name:  data.Name,
		Role:  data.Role,
	})

	if err != nil {
		return nil, err
	}

	expiredTime := time.Duration(24) * time.Hour
	jwtSecret := os.Getenv("JWT_SECRET")

	claims := TokenClaims{
		jwt.RegisteredClaims{
			Issuer:    "Edukita-Backend",
			Subject:   user.ID,
			Audience:  []string{"user"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Message:    "Success",
		Data:       LoginData{Token: signedToken, User: user},
		StatusCode: http.StatusCreated,
		Success:    true,
	}, nil
}

func (s *userImpl) Login(ctx context.Context, email string) (*LoginResponse, error) {
	expiredTime := time.Duration(24) * time.Hour
	jwtSecret := os.Getenv("JWT_SECRET")

	user, err := s.repo.FindUserByEmail(email)

	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, fmt.Errorf("User not Found")
	}

	claims := TokenClaims{
		jwt.RegisteredClaims{
			Issuer:    "Edukita-Backend",
			Subject:   user.ID,
			Audience:  []string{"user"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiredTime)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		*user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Message:    "Success",
		Data:       LoginData{Token: signedToken, User: *user},
		StatusCode: http.StatusOK,
		Success:    true,
	}, nil
}
