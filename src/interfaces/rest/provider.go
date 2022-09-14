//go:build wireinject
// +build wireinject

package rest

import (
	"github.com/google/wire"
	"github.com/hendrorahmat/golang-clean-architecture/src/applications"
	"github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/v1/simkah_app/handler"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var BankHandlerSet = wire.NewSet(wire.Struct(new(handler.BankHandler), "*"))

var (
	ProviderHandlerSet wire.ProviderSet = wire.NewSet(
		BankHandlerSet,
		applications.ProviderUsecaseSet,
		wire.Struct(new(Handler), "*"),
		wire.Bind(new(handler.IBankHandler), new(*handler.BankHandler)),
	)
)

func InjectHandler(db *gorm.DB, logger *logrus.Logger, defaultJoins ...string) *Handler {
	panic(wire.Build(ProviderHandlerSet))
}
