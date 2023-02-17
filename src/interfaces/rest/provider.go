//go:build wireinject
// +build wireinject

package rest

import (
	"github.com/google/wire"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications/usecases"
	handler2 "github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/oauth2/handler"
	"github.com/sirupsen/logrus"
)

func ProvideOauthClientHandler(u *usecases.Usecase, logger *logrus.Logger) *handler2.OauthClientHandler {
	return &handler2.OauthClientHandler{
		Usecase: u.OauthUsecase,
		Logger:  logger,
	}
}

func ProvideOauthTokenHandler(u *usecases.Usecase, logger *logrus.Logger) *handler2.OauthTokenHandler {
	return &handler2.OauthTokenHandler{
		Usecase: u.OauthUsecase,
		Logger:  logger,
	}
}

var (
	ProviderHandlerSet wire.ProviderSet = wire.NewSet(
		ProvideOauthClientHandler,
		ProvideOauthTokenHandler,
		wire.Struct(new(Handler), "*"),
		wire.Struct(new(handler2.Oauth2Handler), "*"),
		wire.Bind(new(handler2.IOauthClientHandler), new(*handler2.OauthClientHandler)),
		wire.Bind(new(handler2.IOauthTokenHandler), new(*handler2.OauthTokenHandler)),
	)
)

func InjectHandler(usecases *usecases.Usecase, logger *logrus.Logger, defaultJoins ...string) *Handler {
	panic(wire.Build(ProviderHandlerSet))
}
