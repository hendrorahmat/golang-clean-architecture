package repositories

type IOauthAuthCodeRepository interface {
	FindByCode(code string)
}
