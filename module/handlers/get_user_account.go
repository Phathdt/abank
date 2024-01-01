package handlers

import (
	"abank/module/models"
	"abank/shared/errorx"
	"context"
	"github.com/phathdt/service-context/core"
)

type GetUserAccountRepo interface {
	ListAccount(ctx context.Context, userId int) ([]models.Account, error)
}

type getUserAccountHdl struct {
	repo GetUserAccountRepo
}

func NewGetUserAccountHdl(repo GetUserAccountRepo) *getUserAccountHdl {
	return &getUserAccountHdl{repo: repo}
}

func (h *getUserAccountHdl) Response(ctx context.Context, userId int) ([]models.Account, error) {
	accounts, err := h.repo.ListAccount(ctx, userId)
	if err != nil {
		return nil, core.ErrBadRequest.
			WithError(errorx.ErrCannotListAccounts.Error()).
			WithDebug(err.Error())
	}

	return accounts, nil
}
