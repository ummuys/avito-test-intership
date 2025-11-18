package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/models"
	"github.com/ummuys/avito-test-intership/internal/secure"
	"github.com/ummuys/avito-test-intership/internal/service"
)

type ah struct {
	logger *zerolog.Logger
	tm     secure.TokenManager
	aus    service.AuthService
}

func NewAuthHandler(tm secure.TokenManager, aus service.AuthService, logger *zerolog.Logger) AuthHandler {
	return &ah{logger: logger, tm: tm, aus: aus}
}

func (ah *ah) UpdateAccessToken(g *gin.Context) {
	ah.logger.Debug().Str("evt", "call UpdateAccessToken").Msg("")
	refreshToken, err := g.Cookie("refresh_token")
	if err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeBadToken,
			Message: errs.ErrMsgBadToken,
		}
		g.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err})
		return
	}

	claims, err := ah.tm.ValidateRefreshToken(refreshToken)
	if err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeBadToken,
			Message: errs.ErrMsgBadToken,
		}
		g.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorResponse{Error: err})
		return
	}

	userID := claims.UserID
	role := claims.Role

	access, err := ah.tm.GenerateAccessToken(userID, role)
	if err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeInternal,
			Message: errs.ErrMsgInternal,
		}
		g.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: err})
		return
	}

	g.Set("msg", "access token is updated")
	g.JSON(http.StatusOK, models.AuthResponse{AccessToken: access})
}

func (ah *ah) Authorization(g *gin.Context) {
	ah.logger.Debug().Str("evt", "call Authorization").Msg("")
	ctx := g.Request.Context()
	var req models.AuthRequest
	if err := g.ShouldBindBodyWithJSON(&req); err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeBadJSON,
			Message: errs.ErrMsgBadJSON,
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		return
	}

	id, role, err := ah.aus.CheckCredentials(ctx, req.Username, req.Password)
	if err != nil {
		g.Set("msg", err.Error())
		switch {
		case errors.Is(err, errs.ErrPGNotFound):
			err := models.Error{
				Code:    errs.ErrCodeNotFound,
				Message: errs.ErrMsgNotFound,
			}
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		case errors.Is(err, errs.ErrInvalidCredentials):
			err := models.Error{
				Code:    errs.ErrCodeInvalidCredentials,
				Message: errs.ErrMsgInvalidCredentials,
			}
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		default:
			err := models.Error{
				Code:    errs.ErrCodeInternal,
				Message: errs.ErrMsgInternal,
			}
			g.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: err})
		}
		return
	}

	access, err := ah.tm.GenerateAccessToken(id, role)
	if err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeInternal,
			Message: errs.ErrMsgInternal,
		}
		g.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: err})
		return
	}

	refresh, err := ah.tm.GenerateRefreshToken(id, role)
	if err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeInternal,
			Message: errs.ErrMsgInternal,
		}
		g.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: err})
		return
	}

	cfg := ah.tm.GetConfiguration()

	g.Set("msg", "auth successful")
	g.SetCookie("refresh_token", refresh, int(cfg.RefreshTokenLimit.Seconds()), "/", "", false, true)
	g.JSON(http.StatusOK, models.AuthResponse{AccessToken: access})
}
