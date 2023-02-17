package utils

import (
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/listeners"
	domain_events "github.com/hendrorahmat/golang-clean-architecture/src/domain/events"
	"github.com/sirupsen/logrus"
)

func DispatchEvent(logger *logrus.Logger, events ...domain_events.IEvent) {
	for _, event := range events {
		logger.Info(fmt.Sprintf("Event name %s", event.GetName()))
		listenersName := event.GetListenersName()

		for _, listenerName := range listenersName {
			logger.Info(fmt.Sprintf("Listener name %s", listenerName))

			listener, ok := listeners.ListenerObject[listeners.ListenerName(listenerName)]
			if !ok {
				logger.Info(fmt.Sprintf("Listener %s not found", listenerName))
				continue
			}

			if listener.ShouldHandleAsync() {
				logger.Info("Handle Async")
				go listener.Handle(event)
			} else {
				logger.Info("Handle Sync")
				listener.Handle(event)
			}

		}
	}
}
