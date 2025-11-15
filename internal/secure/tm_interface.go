package secure

import "github.com/ummuys/avito-test-intership/internal/config"

type TokenManager interface {
	GetConfiguration() config.TMConfig
	GenerateRefreshToken(user_id int64, role string) (string, error)
	GenerateAccessToken(user_id int64, role string) (string, error)
	ValidateAccessToken(rawToken string) (AccessClaims, error)
	ValidateRefreshToken(rawToken string) (RefreshClaims, error)
}
