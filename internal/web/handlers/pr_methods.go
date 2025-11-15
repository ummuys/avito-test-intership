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

type prh struct {
	svc    service.PRService
	logger *zerolog.Logger
}

func NewPRHandler(svc service.PRService, logger *zerolog.Logger) PRHandler {
	return &prh{svc: svc, logger: logger}
}

func (p *prh) Create(g *gin.Context) {
	ctx := g.Request.Context()
	var req models.CreatePRRequest
	if err := g.ShouldBindBodyWithJSON(&req); err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeBadJSON,
			Message: errs.ErrMsgBadJSON,
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		return
	}

	pr, err := p.svc.Create(ctx, req.PRID, req.PRName, req.AuthorID)
	if err != nil {
		g.Set("msg", err.Error())
		switch {
		case errors.Is(err, errs.ErrPGDuplicate):
			err := models.Error{
				Code:    errs.ErrCodePRExists,
				Message: errs.ErrMsgPRExists,
			}
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		case errors.Is(err, errs.ErrPGForeignKey):
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

	g.Set("msg", "pull_request created")
	g.JSON(http.StatusCreated, models.CreatePRWrapper{PR: pr})
}

func (p *prh) Merge(g *gin.Context) {
	ctx := g.Request.Context()
	var req models.MergePRRequest
	if err := g.ShouldBindBodyWithJSON(&req); err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeBadJSON,
			Message: errs.ErrMsgBadJSON,
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		return
	}

	pr, err := p.svc.Merge(ctx, req.PRID)
	if err != nil {
		g.Set("msg", err.Error())
		switch {
		case errors.Is(err, errs.ErrPGNotFound):
			err := models.Error{
				Code:    errs.ErrCodeNotFound,
				Message: errs.ErrMsgNotFound,
			}
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		case errors.Is(err, errs.ErrNotEnoughReviewers):
			err := models.Error{
				Code:    errs.ErrCodeNotEnoughReviewers,
				Message: errs.ErrMsgNotEnoughReviewers,
			}
			g.AbortWithStatusJSON(http.StatusConflict, models.ErrorResponse{Error: err})
		default:
			err := models.Error{
				Code:    errs.ErrCodeInternal,
				Message: errs.ErrMsgInternal,
			}
			g.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: err})
		}
		return
	}
	g.Set("msg", "pull_request merged")
	g.JSON(http.StatusOK, models.MergeRPWrapper{PR: pr})
}

func (p *prh) Reassign(g *gin.Context) {
	ctx := g.Request.Context()
	var req models.ReassignPRRequest
	if err := g.ShouldBindBodyWithJSON(&req); err != nil {
		g.Set("msg", err.Error())
		err := models.Error{
			Code:    errs.ErrCodeBadJSON,
			Message: errs.ErrMsgBadJSON,
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		return
	}

	pr, err := p.svc.Reassign(ctx, req.PRID, req.OldUserID)
	if err != nil {
		g.Set("msg", err.Error())
		switch {
		case errors.Is(err, errs.ErrPGNotFound):
			err := models.Error{
				Code:    errs.ErrCodeNotFound,
				Message: errs.ErrMsgNotFound,
			}
			g.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorResponse{Error: err})
		case errors.Is(err, errs.ErrPRMerged):
			err := models.Error{
				Code:    errs.ErrCodePRMerged,
				Message: errs.ErrMsgPRMerged,
			}
			g.AbortWithStatusJSON(http.StatusConflict, models.ErrorResponse{Error: err})
		case errors.Is(err, errs.ErrRVNotAssigned):
			err := models.Error{
				Code:    errs.ErrCodeRVNotAssigned,
				Message: errs.ErrMsgRVNotAssigned,
			}
			g.AbortWithStatusJSON(http.StatusConflict, models.ErrorResponse{Error: err})
		case errors.Is(err, errs.ErrNoCandidate):
			err := models.Error{
				Code:    errs.ErrCodeNoCandidate,
				Message: errs.ErrMsgNoCandidate,
			}
			g.AbortWithStatusJSON(http.StatusConflict, models.ErrorResponse{Error: err})
		default:
			err := models.Error{
				Code:    errs.ErrCodeInternal,
				Message: errs.ErrMsgInternal,
			}
			g.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorResponse{Error: err})
		}

		return
	}

	g.Set("msg", "reassigned pull_request")
	g.JSON(http.StatusOK, models.ReassignRPWrapper{PR: pr})
}
