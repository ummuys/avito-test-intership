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

type uh struct {
	svc    service.UserService
	logger *zerolog.Logger
}

func NewUserHandler(svc service.UserService, logger *zerolog.Logger) UserHandler {
	return &uh{svc: svc, logger: logger}
}

func (u *uh) SetState(g *gin.Context) {
	ctx := g.Request.Context()
	var req models.SetUserStateRequest
	if err := g.ShouldBindBodyWithJSON(&req); err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeBadJSON,
			Message: errs.ErrMsgBadJSON,
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		return
	}

	username, teamName, err := u.svc.SetState(ctx, req.UserID, req.IsActive)
	if err != nil {
		g.Set("msg", err.Error())
		switch {
		case errors.Is(err, errs.ErrPGNotFound):
			err := models.Error{
				Code:    errs.ErrCodeNotFound,
				Message: errs.ErrMsgNotFound,
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

	user := models.User{
		UserID:   req.UserID,
		Username: username,
		TeamName: teamName,
		IsActive: req.IsActive,
	}
	g.Set("msg", "user state changed")
	g.JSON(http.StatusOK, models.SetUserStateResponse{User: user})
}

func (u *uh) Get(g *gin.Context) {

}
