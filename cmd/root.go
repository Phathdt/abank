package cmd

import (
	"abank/module/banktransport/bankfiber"
	"fmt"
	"os"
	"time"

	"github.com/phathdt/service-context/component/gormc"
	"github.com/phathdt/service-context/component/redisc"

	"github.com/phathdt/service-context/component/fiberc"

	"abank/shared/common"

	sctx "github.com/phathdt/service-context"

	"github.com/spf13/cobra"
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

var rootCmd = &cobra.Command{
	Use:   serviceName,
	Short: fmt.Sprintf("start %s", serviceName),
	Run: func(cmd *cobra.Command, args []string) {
		sc := newServiceCtx()

		logger := sctx.GlobalLogger().GetLogger("service")

		time.Sleep(time.Second * 5)

		if err := sc.Load(); err != nil {
			logger.Fatal(err)
		}

		fiberComp := sc.MustGet(common.KeyCompFiber).(fiberc.FiberComponent)

		app := fiberComp.GetApp()

		app.Get("/users/:user_id", bankfiber.GetUser(sc))
		app.Get("/users/:user_id/accounts", bankfiber.GetUserAccount(sc))
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
