package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	todo "github.com/StanislavDimitrenco/restapi"
	"github.com/StanislavDimitrenco/restapi/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const salt = "jhgghendgfhfp[267akjhfnv"
const siningKey = "gferjkasolp34,ll"
const tokenTTL = 12 * time.Hour

type AuthService struct {
	repo repository.Authorisation
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorisation) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (id int, err error) {
	user.Password = generateHashPassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, generateHashPassword(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(siningKey))
}

func (s *AuthService) ParseToken(assesToken string) (int, error) {
	token, err := jwt.ParseWithClaims(assesToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(siningKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return 0, errors.New("token claim are not of type")
	}

	return claims.UserId, nil

}

func generateHashPassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))

}
