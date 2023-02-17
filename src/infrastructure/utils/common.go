package utils

import (
	"github.com/joho/godotenv"
	"math/rand"
	"os/exec"
	"strings"
)

func GetRootPath() string {
	path, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err == nil {
		return strings.TrimSuffix(string(path), "\n")
	}
	return ""
}

func LoadEnv() error {
	rootPath := GetRootPath()
	var err error
	if rootPath == "" {
		rootPath += ".env"
		err = godotenv.Load(".env")
	} else {
		rootPath = rootPath + "/" + ".env"
		err = godotenv.Load(rootPath)
	}

	if err != nil {
		return err
	}
	return nil
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}
