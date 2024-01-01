package cmd

import (
	"abank/module/transport/bankfiber"
	"fmt"
	"github.com/gofiber/swagger"
	"os"
	"time"

	"github.com/phathdt/service-context/component/gormc"
	"github.com/phathdt/service-context/component/redisc"

	"github.com/phathdt/service-context/component/fiberc"

	"abank/shared/common"

	sctx "github.com/phathdt/service-context"

	"github.com/spf13/cobra"

	_ "abank/docs"
)

const (
	serviceName = "abank"
)

func newServiceCtx() sctx.ServiceContext {
	return sctx.NewServiceContext(
		sctx.WithName(serviceName),
		sctx.WithComponent(fiberc.New(common.KeyCompFiber)),
		sctx.WithComponent(gormc.NewGormDB(common.KeyCompGorm, "")),
		sctx.WithComponent(redisc.New(common.KeyCompRedis)),
	)
}

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:4000
// @BasePath /
var rootCmd = &cobra.Command{
	Use:   serviceName,
	Short: fmt.Sprintf("start %s", serviceName),
	Run: func(cmd *cobra.Command, args []string) {
		sc := newServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		time.Sleep(time.Second * 1)

		if err := sc.Load(); err != nil {
			logger.Fatal(err)
		}

		fiberComp := sc.MustGet(common.KeyCompFiber).(fiberc.FiberComponent)

		app := fiberComp.GetApp()

		app.Get("/swagger/*", swagger.HandlerDefault) // default

		app.Get("/users/:user_id/accounts", bankfiber.GetUserAccount(sc))
		app.Get("/users/:user_id", bankfiber.GetUser(sc))
		app.Get("/accounts/:account_id", bankfiber.GetAccount(sc))

		if err := app.Listen(fmt.Sprintf(":%d", fiberComp.GetPort())); err != nil {
			logger.Fatal(err)
		}
	},
}

func Execute() {
	rootCmd.AddCommand(outEnvCmd)
	rootCmd.AddCommand(migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
