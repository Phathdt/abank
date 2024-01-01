package handlers

import (
	"abank/module/models"
	"abank/shared/errorx"
	"context"
	"github.com/phathdt/service-context/core"
)

type GetAccountRepo interface {
	GetAccount(ctx context.Context, accountId int) (*models.Account, error)
}

type getAccountHdl struct {
	repo GetAccountRepo
}

func NewGetAccountHdl(repo GetAccountRepo) *getAccountHdl {
	return &getAccountHdl{repo: repo}
}

func (h *getAccountHdl) Response(ctx context.Context, accountId int) (*models.Account, error) {
	account, err := h.repo.GetAccount(ctx, accountId)
	if err != nil {
		return nil, core.ErrBadRequest.
			WithError(errorx.ErrGetAccount.Error()).
			WithDebug(err.Error())
	}

	return account, nil
}
