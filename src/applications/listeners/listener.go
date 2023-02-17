package listeners

import (
	domain_errors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	domain_events "github.com/hendrorahmat/golang-clean-architecture/src/domain/events"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"sync"
)

type IListener interface {
	ShouldHandleAsync() bool
	Handle(event domain_events.IEvent) domain_errors.DomainError
}

type ListenerName string

var listenerObjectOnce sync.Once
var ListenerObject map[ListenerName]IListener

func init() {
	listenerObjectOnce.Do(func() {
		ListenerObject = map[ListenerName]IListener{
			constants.SendEmailListener: &SendEmailConfirmationListener{},
		}
	})
}
