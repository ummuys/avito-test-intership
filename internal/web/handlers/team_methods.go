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

type th struct {
	svc    service.TeamService
	logger *zerolog.Logger
}

func NewTeamHandler(svc service.TeamService, logger *zerolog.Logger) TeamHandler {
	return &th{svc: svc, logger: logger}
}

func parseAddTeamRequest() {

}

func (t *th) Create(g *gin.Context) {
	ctx := g.Request.Context()
	var req models.AddTeamRequest
	if err := g.ShouldBindBodyWithJSON(&req); err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeBadJSON,
			Message: errs.ErrMsgBadJSON,
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		return
	}

	if err := t.svc.Add(ctx, req); err != nil {
		g.Set("msg", err.Error())
		switch {
		case errors.Is(err, errs.ErrPGDuplicate):
			err := models.Error{
				Code:    errs.ErrCodeTeamExists,
				Message: errs.ErrMsgTeamExists,
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
	g.Set("msg", "team added")
	g.JSON(http.StatusOK, req)
}

func (t *th) Get(g *gin.Context) {
	ctx := g.Request.Context()
	teamName := g.Query("team_name")
	if teamName == "" {
		g.Set("msg", errs.ErrCodeInvalidTeamName)
		err := models.Error{
			Code:    errs.ErrCodeInvalidTeamName,
			Message: errs.ErrMsgInvalidTeamName,
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		return
	}

	team, err := t.svc.Get(ctx, teamName)
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

	g.Set("msg", "team returned")
	g.JSON(http.StatusOK, team)

}
