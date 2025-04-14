package pkg

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService interface {
	GetUserID(token *jwt.Token) (uuid.UUID, error)
}

type JWT struct{}

func NewJWT() JWTService {
	return &JWT{}
}

func (j *JWT) GetUserID(token *jwt.Token) (uuid.UUID, error) {
	userIDStr, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}
