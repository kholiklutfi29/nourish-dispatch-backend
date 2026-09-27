package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kholiklutfi29/nourish-dispatch/internal/models"
)

type JWTService struct {
	secretKey []byte
	expiresIn  time.Duration
}

func NewJWTService(
	secretKey string,
	expiresIn time.Duration,
) *JWTService {
	return &JWTService{
		secretKey: []byte(secretKey),
		expiresIn:  expiresIn,
	}
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   models.UserRole `json:"role"`

	jwt.RegisteredClaims // embedded struct, so we can use the registered claims
} 

func (j *JWTService) GenerateJWTToken(user *models.User) (string, error) {
	now := time.Now()

	claims := Claims{
		UserID: user.ID.String(),
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: user.ID.String(),
			IssuedAt: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(j.expiresIn)), // now + expiresIn duration
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(j.secretKey)
}

func (j *JWTService) ValidateJWTToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{}, // send null struct claims to parse the token into
		func(token *jwt.Token) (interface{}, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("unexpected signing method")
			}
			return j.secretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	// type assertion to get the claims from the token, to pointer of Claims struct
	claims, ok := token.Claims.(*Claims) 

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}