// internal/service/auth_service.go
package service

import (
	"github.com/user-authentication-go/backend/internal/model"
	"github.com/user-authentication-go/backend/internal/pkg/auth"
	"github.com/user-authentication-go/backend/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    userRepo    *repository.UserRepository
    jwtService  *auth.JWTService
}

func NewAuthService(userRepo *repository.UserRepository, jwtService *auth.JWTService) *AuthService {
    return &AuthService{
        userRepo:    userRepo,
        jwtService:  jwtService,
    }
}

func (s *AuthService) Login(req model.LoginRequest) (string, error) {
    user, err := s.userRepo.FindByEmail(req.Email)
    if err != nil {
        return "", err
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return "", err
    }

    token, err := s.jwtService.GenerateToken(user.ID.String(), user.Username)
    if err != nil {
        return "", err
    }

    return token, nil
}
