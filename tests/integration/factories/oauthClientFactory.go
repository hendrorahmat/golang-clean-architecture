package factories

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/gofrs/uuid"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/constants"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/persistance/database/models"
	"gorm.io/gorm"
	"time"
)

func NewOauthClientFactory() models.OauthClient {
	timeNow := time.Now()

	return models.OauthClient{
		ID:               uuid.FromStringOrNil(gofakeit.UUID()),
		Name:             gofakeit.Username(),
		EnabledGrantType: []string{constants.ClientCredentialsGrantType},
		Secret:           gofakeit.Password(true, true, true, true, true, 2),
		Redirect:         gofakeit.URL(),
		CreatedAt:        &timeNow,
		UpdatedAt:        &timeNow,
		DeletedAt:        gorm.DeletedAt{},
	}
}
