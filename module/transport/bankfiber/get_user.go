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

// GetUser func gets user by given ID or 404 error.
// @Description Get user by given ID.
// @Summary get user by given ID
// @Tags User
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func GetUser(sc sctx.ServiceContext) fiber.Handler {
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

		hdl := handlers.NewGetUserHdl(r)
		user, err := hdl.Response(ctx.Context(), data.UserId)
		if err != nil {
			panic(err)
		}

		return ctx.Status(http.StatusOK).JSON(core.SimpleSuccessResponse(user))
	}
}
