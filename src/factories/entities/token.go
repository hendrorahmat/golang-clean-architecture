package entities

//import (
//	"errors"
//	"github.com/gin-gonic/gin"
//	"github.com/hendrorahmat/golang-clean-architecture/src/applications/dto"
//	"github.com/hendrorahmat/golang-clean-architecture/src/domain/model/aggregates"
//	"github.com/hendrorahmat/golang-clean-architecture/src/domain/model/aggregates/contracts"
//	"github.com/hendrorahmat/golang-clean-architecture/src/domain/repositories"
//	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
//)
//
//func TokenFactory(data *dto.IssueToken, clientRepository repositories.IOauthClientRepository, ctx *gin.Context) (contracts.IIssuedToken, error) {
//	switch data.GrantType {
//	case constants.ClientCredentialsGrantType:
//		result, _ := aggregates.NewTokenClientCredentials()
//		return &result, nil
//	default:
//		return nil, errors.New("grant type not found")
//	}
//}
