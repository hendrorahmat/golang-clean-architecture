package main

import (
	"context"
	"github.com/hendrorahmat/golang-clean-architecture/cmd"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/config"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructures/utils"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	cmd.Execute()
	//if err != nil {
	//	panic(err)
	//}
	//err := godotenv.Load()

	// repositories
	//bankRepository := repositories.NewBankRepository(db.Postgres.DB(), conf.Log.Logger)

	// usecase
	//bankUsecase := usecases.NewBankUsecase(
	//	bankRepository,
	//)

	//usecase := usecases.UsecaseList{
	//	BankUsecase: bankUsecase,
	//}
	//
	//router := rest.NewRoute(conf, usecase)

	//srv := &http.Server{
	//	Addr:    ":" + conf.Http.Port,
	//	Handler: gin.Default(),
	//}
	//
	//go func() {
	//	if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
	//		//logger.Warnf("listen %s\n", err)
	//	}
	//}()
	//
	//logger.Info("Server listen ", conf.Http.Port)
	//timeout := time.Duration(conf.App.GracefulShutdownTimeout) * time.Second
	//wait := gracefulShutdown(ctx, logger, timeout, map[string]operation{
	//	"db-postgres": func(ctx context.Context) error {
	//		return db.Postgres.SqlDB().Close()
	//	},
	//	"http-server": func(ctx context.Context) error {
	//		return srv.Shutdown(ctx)
	//	},
	//})
	//<-wait
	//fmt.Println("Closed!")
}

func gracefulShutdown(ctx context.Context, log *logrus.Logger, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		data := <-s
		/**
		program in below will never run if <-s not return a value because assigning value as channel will run
		syncronizely
		*/

		log.Warnf("Signal %s pid : %d", data.String(), data.Signal)
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Warnf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()
		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()
				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Fatalf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}
				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}
		wg.Wait()
		close(wait)
	}()
	return wait
}

func setupLogger(conf *config.Config) *logrus.Logger {
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name

	isProd := false

	if conf.App.Environment == "production" {
		isProd = true
	}

	logger := utils.NewLogInstance(
		utils.LogName(conf.Log.Name),
		utils.IsProduction(isProd),
		utils.LogAdditionalFields(m))
	return logger
}

type operation func(ctx context.Context) error
