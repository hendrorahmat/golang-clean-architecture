package oauth

import (
	"github.com/hendrorahmat/golang-clean-architecture/src/domain/errors"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"net/http"
)

type ClientIdAndSecretNotFoundError struct {
	clientId     string
	clientSecret string
}

func (*ClientIdAndSecretNotFoundError) GetStatusCode() int {
	return http.StatusNotFound
}

func (c *ClientIdAndSecretNotFoundError) GetCode() uint {
	//TODO implement me
	panic("implement me")
}

func (c *ClientIdAndSecretNotFoundError) GetTitle() string {
	//TODO implement me
	panic("implement me")
}

func (c *ClientIdAndSecretNotFoundError) Error() string {
	return constants.ClientIdAndSecretNotFound
}

func ThrowClientIdAndSecretNotFound(clientId, clientSecret string) errors.DomainError {
	return &ClientIdAndSecretNotFoundError{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}
