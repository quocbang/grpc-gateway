package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"github.com/quocbang/grpc-gateway/server/utils/roles"
)

type TokenMaker interface {
	GenerateToken() (string, error)
	VerifyToken(string) (any, error)
}

type JWTClaimCustom struct {
	SessionID uuid.UUID
	Username  string
	Role      roles.Roles
	jwt.RegisteredClaims
}

type UserInfo struct {
	Username string
	Role     roles.Roles
}

type JWT struct {
	SecretKey     string
	User          UserInfo
	TokenLifeTime time.Duration
}

func (j JWT) GenerateToken() (string, *JWTClaimCustom, error) {
	if j.SecretKey == "" {
		return "", nil, fmt.Errorf("missing secret key")
	}

	claims := &JWTClaimCustom{
		SessionID: uuid.New(),
		Username:  j.User.Username,
		Role:      j.User.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(j.TokenLifeTime),
			},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", nil, err
	}

	return signedToken, claims, nil
}

func (j JWT) VerifyToken(token string) (*JWTClaimCustom, error) {
	if token == "" {
		return nil, fmt.Errorf("missing token")
	}

	if j.SecretKey == "" {
		return nil, fmt.Errorf("missing secret key")
	}

	claims := JWTClaimCustom{}
	_, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return []byte(j.SecretKey), nil
	})
	if err != nil {
		if e, ok := err.(*jwt.ValidationError); ok {
			if e.Errors == jwt.ValidationErrorExpired {
				return nil, fmt.Errorf("token expired")
			}
			return nil, fmt.Errorf("invalid token")
		}
		return nil, err
	}

	return &claims, nil
}
