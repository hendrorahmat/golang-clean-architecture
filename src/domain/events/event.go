package domain_events

type IEvent interface {
	GetName() string
	GetListenersName() []string
}
