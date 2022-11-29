package rest

import "github.com/hendrorahmat/golang-clean-architecture/src/interfaces/rest/routes/v1/simkah_app/handler"

type Handler struct {
	BankHandler  handler.IBankHandler
	EventHandler handler.IBankHandler
}
