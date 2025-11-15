package secure

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ummuys/avito-test-intership/internal/config"
)

type tokMan struct {
	config config.TMConfig
}

func NewTokenManager() (TokenManager, error) {
	c, err := config.ParseTMConfig()
	if err != nil {
		return nil, err
	}
	return &tokMan{config: c}, nil
}

func (tm *tokMan) GenerateRefreshToken(user_id int64, role string) (string, error) {
	claims := RefreshClaims{
		UserID: user_id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.config.RefreshTokenLimit)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "rp-service-token-manager",
			Audience:  []string{"rp-service-user"},
		},
	}
	refresh := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return refresh.SignedString([]byte(tm.config.RefreshSecret))
}

func (tm *tokMan) GetConfiguration() config.TMConfig {
	return tm.config
}

func (tm *tokMan) GenerateAccessToken(user_id int64, role string) (string, error) {
	claims := AccessClaims{
		UserID: user_id,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tm.config.AccessTokenLimit)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "rp-service-token-manager",
			Audience:  []string{"rp-service-user"},
		},
	}
	access := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return access.SignedString([]byte(tm.config.AccessSecret))
}

func (tm *tokMan) ValidateAccessToken(rawToken string) (AccessClaims, error) {
	token, claims, err := tm.unhashAccessToken(rawToken)
	if err != nil {
		return AccessClaims{}, err
	}
	if !token.Valid {
		return AccessClaims{}, fmt.Errorf("invalid token")
	}
	return claims, err
}

func (tm *tokMan) ValidateRefreshToken(rawToken string) (RefreshClaims, error) {
	token, claims, err := tm.unhashRefreshToken(rawToken)
	if err != nil {
		return RefreshClaims{}, err
	}
	if !token.Valid {
		return RefreshClaims{}, fmt.Errorf("invalid token")
	}
	return claims, err
}

func (tm *tokMan) unhashAccessToken(token string) (*jwt.Token, AccessClaims, error) {
	var claims AccessClaims
	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(tm.config.AccessSecret), nil
	})
	return jwtToken, claims, err
}

func (tm *tokMan) unhashRefreshToken(token string) (*jwt.Token, RefreshClaims, error) {
	var claims RefreshClaims
	jwtToken, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(tm.config.RefreshSecret), nil
	})
	return jwtToken, claims, err
}
