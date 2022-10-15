package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/src/bootstrap"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/utils"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest"
	"github.com/spf13/cobra"
	"net/http"
	"time"
)

var httpCommand = &cobra.Command{
	Use:   "http",
	Short: "Run HTTP Api",
	RunE: func(cmd *cobra.Command, args []string) error {
		application := bootstrap.Boot()
		ctx := context.Background()

		router := rest.NewRoute(application)
		srv := &http.Server{
			Addr:    ":" + application.GetConfig().Http.Port,
			Handler: router,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
				application.GetLogger().Warnf("listen %s\n", err)
			}
		}()
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
