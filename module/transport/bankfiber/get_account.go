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

func GetAccount(sc sctx.ServiceContext) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		type RequestParams struct {
			AccountId int `params:"account_id" validate:"required,number"`
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

		hdl := handlers.NewGetAccountHdl(r)
		account, err := hdl.Response(ctx.Context(), data.AccountId)
		if err != nil {
			panic(err)
		}

		return ctx.Status(http.StatusOK).JSON(core.SimpleSuccessResponse(account))
	}
}
