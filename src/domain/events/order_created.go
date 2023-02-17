package domain_events

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
)

type OrderCreated struct {
	Order string
}

func (c *OrderCreated) GetListenersName() []string {
	return []string{
		constants.SendEmailListener,
	}
}

func (*OrderCreated) GetName() string {
	return constants.OrderCreatedEvent
}

func OrderCreatedEvent(order string) IEvent {
	return &OrderCreated{Order: order}
}

var _ IEvent = &OrderCreated{}
