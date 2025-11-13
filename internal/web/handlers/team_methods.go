package handlers

import (
	"errors"
	"fmt"
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

func (t *th) Add(g *gin.Context) {
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

	mbrs, err := t.svc.Get(ctx, teamName)
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

	badData := false
	badMID := 0
	resp := models.GetTeamResponse{TeamName: teamName}
	resp.Members = make([]models.Member, len(mbrs))

	for i, m := range mbrs {

		userID, ok := m[0].(string)
		if !ok {
			badMID = i
			badData = true
			break
		}

		username, ok := m[1].(string)
		if !ok {
			badMID = i
			badData = true
			break
		}

		isActive, ok := m[2].(bool)
		if !ok {
			badMID = i
			badData = true
			break
		}

		resp.Members[i] = models.Member{
			UserID:   userID,
			Username: username,
			IsActive: isActive,
		}

	}

	if badData {
		msg := fmt.Sprintf("can't convert field/s of models.User into go type {mbrs[%d] = %v}",
			badMID, mbrs[badMID])
		g.Set("msg", msg)
		err := models.Error{
			Code:    errs.ErrCodeInternal,
			Message: errs.ErrMsgInternal,
		}
		g.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: err})
		return
	}

	g.Set("msg", "team returned")
	g.JSON(http.StatusOK, resp)

}
