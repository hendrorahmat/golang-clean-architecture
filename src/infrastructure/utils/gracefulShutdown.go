package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type GracefulOperation func(ctx context.Context) error

func GracefulShutdown(ctx context.Context, log *logrus.Logger, timeout time.Duration, ops map[string]GracefulOperation) <-chan struct{} {
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
