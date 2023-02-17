package config

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hendrorahmat/golang-clean-architecture/src/infrastructure/utils"
	"sync"
	"time"
)

const OneDay = 24 * time.Hour
const LoginExpirationDuration = OneDay

var JwtSigningMethod = jwt.SigningMethodRS256

var JwtSignatureKey []byte
var jwtConfigOnce sync.Once

func MakeJwtConfig() {
	jwtConfigOnce.Do(func() {
		privateKey, err := utils.GetOauthPrivateKeyFile()
		if err != nil {
			fmt.Println("Error read private key")
			panic(err)
			return
		}
		JwtSignatureKey = privateKey
	})
}
