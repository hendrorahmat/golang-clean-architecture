package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/gops/agent"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications"
	"github.com/hendrorahmat/golang-clean-architecture/src/bootstrap"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/utils"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var httpCommand = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP Api",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		application := bootstrap.Boot()

		repositories := database.InjectRepository(application.GetActiveConnection().DB(), application.GetLogger())
		usecases := applications.InjectUsecase(repositories, application.GetLogger())
		controllers := rest.InjectHandler(usecases, application.GetLogger())

		router := rest.NewRoute(ctx, controllers, application.GetConfig())
		srv := &http.Server{
			Addr:    ":" + application.GetConfig().Http.Port,
			Handler: router,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
				application.GetLogger().Warnf("listen %s\n", err)
			}
		}()

		if err := agent.Listen(agent.Options{}); err != nil {
			application.GetLogger().Info("Gops Error")
			application.GetLogger().Fatal(err)
		}
		application.GetLogger().Info("Server listen ", application.GetConfig().Http.Port)
		timeout := time.Duration(application.GetConfig().App.GracefulShutdownTimeout) * time.Second

		operations := map[string]utils.GracefulOperation{}

		for c, connection := range application.GetConnections().Connection {
			operations["database-"+c] = func(ctx context.Context) error {
				return connection.SqlDB().Close()
			}
		}

		operations["http-server"] = func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		}

		wait := utils.GracefulShutdown(ctx, application.GetLogger(), timeout, operations)
		<-wait
		fmt.Println("Closed!")

		return nil
	},
}
