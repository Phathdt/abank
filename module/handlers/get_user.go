package handlers

import (
	"abank/module/models"
	"abank/shared/errorx"
	"context"
	"github.com/phathdt/service-context/core"
)

type GetUserRepo interface {
	GetUser(ctx context.Context, userId int) (*models.User, error)
}

type getUserHdl struct {
	repo GetUserRepo
}

func NewGetUserHdl(repo GetUserRepo) *getUserHdl {
	return &getUserHdl{repo: repo}
}

func (h *getUserHdl) Response(ctx context.Context, userId int) (*models.User, error) {
	user, err := h.repo.GetUser(ctx, userId)
	if err != nil {
		return nil, core.ErrBadRequest.
			WithError(errorx.ErrGetUser.Error()).
			WithDebug(err.Error())
	}

	return user, nil
}
