package controller

import (
	"github.com/cos-plat/backend/internal/model"
	"github.com/cos-plat/backend/internal/pkg/auth"
	"github.com/cos-plat/backend/internal/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
    userRepo   *repository.UserRepository
    jwtService *auth.JWTService
}

func NewAuthService(userRepo *repository.UserRepository, jwtService *auth.JWTService) *AuthService {
    return &AuthService{
        userRepo:   userRepo,
        jwtService: jwtService,
    }
}

// Register メソッドを追加
func (s *AuthService) Register(req model.RegisterRequest) (*model.User, error) {
    // パスワードのハッシュ化
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    // 新しいユーザーを作成
    user := &model.User{
        ID:       uuid.New(),
        Username: req.Username,
        Email:    req.Email,
        Password: string(hashedPassword),
    }

    // データベースに保存
    if err := s.userRepo.Create(user); err != nil {
        return nil, err
    }

    return user, nil
}

// Login メソッドを追加
func (s *AuthService) Login(req model.LoginRequest) (string, error) {
    // メールアドレスでユーザーを検索
    user, err := s.userRepo.FindByEmail(req.Email)
    if err != nil {
        return "", model.ErrInvalidCredentials
    }

    // パスワードを検証
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
        return "", model.ErrInvalidCredentials
    }

    // JWTトークンを生成
    token, err := s.jwtService.GenerateToken(user.ID.String(), user.Username)
    if err != nil {
        return "", err
    }

    return token, nil
}
