package utils

import "os"

func GetEnv(name string) string {
	return os.Getenv(name)
}

func GetEnvWithDefaultValue(name string, defaultValue string) string {
	if os.Getenv(name) == "" {
		return defaultValue
	}
	return os.Getenv(name)
}

func GetOauthPrivateKeyFile() ([]byte, error) {
	rootPath := GetRootPath()
	path := rootPath + "/storage/oauth-private.key"
	privateKey, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func GetOauthPublicKeyFile() ([]byte, error) {
	rootPath := GetRootPath()
	path := rootPath + "/storage/oauth-public.key"
	publicKey, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
