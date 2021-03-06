package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Dmytro-yakymuk/travelwithme/internal/models"
	"github.com/Dmytro-yakymuk/travelwithme/internal/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "skjdhgjSKALdjhgdsnjkdzjhxfdjk"
	signingKey = "dsjdhghsjaklsdkjfhgsjkldjhgu62u"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"user_id"`
	Role   string `json:"user_role"`
}

type AuthorizationService struct {
	authRepository repository.Authorization
}

func NewAuthorizationService(authRepository repository.Authorization) *AuthorizationService {
	return &AuthorizationService{authRepository: authRepository}
}

func (s *AuthorizationService) CreateUser(user models.User) error {
	user.Password = generatePasswordHash(user.Password)
	return s.authRepository.CreateUser(user)
}

func (s *AuthorizationService) GenerateToken(email, password string) (map[string]interface{}, error) {
	user, err := s.authRepository.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return nil, err
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
		user.Role,
	})

	token, err := tokenClaims.SignedString([]byte(signingKey))
	return map[string]interface{}{
		"token": token,
		"id":    user.Id,
		"role":  user.Role,
	}, err
}

func (s *AuthorizationService) ParseToken(accessToken string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return nil, errors.New("token claims are not of role *tokenClaims")
	}

	return map[string]interface{}{
		"id":   claims.UserId,
		"role": claims.Role,
	}, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}