package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/ummuys/avito-test-intership/internal/errs"
	"github.com/ummuys/avito-test-intership/internal/models"
	"github.com/ummuys/avito-test-intership/internal/service"
)

type adh struct {
	logger *zerolog.Logger
	srv    service.AdminService
}

func NewAdminHandler(ads service.AdminService, logger *zerolog.Logger) AdminHandler {
	return &adh{logger: logger, srv: ads}
}

func (ad *adh) CreateUser(g *gin.Context) {
	ad.logger.Debug().Str("evt", "call CreateUser").Msg("")
	ctx := g.Request.Context()

	var req models.CreateUserRequest
	if err := g.ShouldBindBodyWithJSON(&req); err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeBadJSON,
			Message: errs.ErrMsgBadJSON,
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		return
	}

	if err := ad.srv.CreateUser(ctx, req.Username, req.Password, req.Role); err != nil {
		g.Set("msg", err.Error())
		switch {
		case errors.Is(err, errs.ErrPGDuplicate):
			err := models.Error{
				Code:    errs.ErrCodeUserExists,
				Message: errs.ErrMsgUserExists,
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

	g.Set("msg", "user created")
	g.JSON(http.StatusCreated, models.CreateUserResponse{Username: req.Username, Role: req.Role})
}
