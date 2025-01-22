package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
    UserID    string `json:"user_id"`
    Username  string `json:"username"`
    jwt.RegisteredClaims
}

type JWTService struct {
    secretKey []byte
    expires   time.Duration
}

func NewJWTService(secretKey string, expires time.Duration) *JWTService {
    return &JWTService{
        secretKey: []byte(secretKey),
        expires:   expires,
    }
}

func (s *JWTService) GenerateToken(userID, username string) (string, error) {
    claims := JWTClaims{
        UserID:   userID,
        Username: username,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expires)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(s.secretKey)
}

func (s *JWTService) ValidateToken(tokenString string) (*JWTClaims, error) {
    token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
        return s.secretKey, nil
    })

    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
        return claims, nil
    }

    return nil, errors.New("invalid token")
}
