package bankfiber

import (
	"abank/module/handlers"
	"abank/module/repo"
	"abank/module/storage"
	"abank/plugins/validation"
	"abank/shared/common"
	"github.com/gofiber/fiber/v2"
	sctx "github.com/phathdt/service-context"
	"github.com/phathdt/service-context/component/gormc"
	"github.com/phathdt/service-context/component/redisc"
	"github.com/phathdt/service-context/core"
	"net/http"
)

// GetUserAccount func gets accounts by given user ID or 404 error.
// @Description Get accounts by given user ID.
// @Summary get accounts by given user ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "Account ID"
// @Success 200 {array} models.User
// @Router /users/{id}/accounts [get]
func GetUserAccount(sc sctx.ServiceContext) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		type RequestParams struct {
			UserId int `params:"user_id" validate:"required,number"`
		}

		var data RequestParams

		if err := ctx.ParamsParser(&data); err != nil {
			panic(err)
		}

		if err := validation.Validate(data); err != nil {
			panic(err)
		}

		db := sc.MustGet(common.KeyCompGorm).(gormc.GormComponent).GetDB()
		client := sc.MustGet(common.KeyCompRedis).(redisc.RedisComponent).GetClient()

		store := storage.NewSqlStore(db)
		cacheStore := storage.NewCacheStore(client)
		r := repo.NewRepo(store, cacheStore)

		hdl := handlers.NewGetUserAccountHdl(r)
		accounts, err := hdl.Response(ctx.Context(), data.UserId)
		if err != nil {
			panic(err)
		}

		return ctx.Status(http.StatusOK).JSON(core.SimpleSuccessResponse(accounts))
	}
}
