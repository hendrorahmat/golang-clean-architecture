package listeners

import (
	"fmt"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/services"
	domain_errors "github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/events"
)

type SendEmailConfirmationListener struct{}

func (*SendEmailConfirmationListener) ShouldHandleAsync() bool {
	return false
}

func (*SendEmailConfirmationListener) Handle(event domain_events.IEvent) domain_errors.DomainError {
	objectEvent := event.(*domain_events.OrderCreated)
	objectEvent.GetName()
	fmt.Println(objectEvent.Order)
	emailService := services.NewSendEmailService("halo",
		[]string{""},
		[]string{""},
		[]string{""},
		"",
		"",
		[]string{""},
	)
	emailService.Send()
	return nil
}

var _ IListener = &SendEmailConfirmationListener{}
