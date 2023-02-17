package errors

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
	"strings"
)

type EventDuplicateError struct {
	eventName string
}

func (e *EventDuplicateError) Error() string {
	return strings.Replace(constants.ErrorEventNameAlreadyAdded, "{name}", e.eventName, 1)
}

func (e *EventDuplicateError) GetTitle() string {
	return constants.ErrorEventAlreadyAdded
}

func (e *EventDuplicateError) GetCode() uint {
	return constants.InvalidHeaderCode
}

func (e *EventDuplicateError) GetStatusCode() int {
	return http.StatusInternalServerError
}

func ThrowEventDuplicateError(eventName string) DomainError {
	return &EventDuplicateError{eventName: eventName}
}
