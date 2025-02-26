package notion

import (
	"errors"
	"os"

	"github.com/Hayao0819/Abracadabra/conf"
)

func getToken() (string, error) {
	config, err := conf.Get()
	if err != nil {
		return "", err
	}
	varName := config.TokenVariableName

	token := os.Getenv(varName)
	if token == "" {
		return "", errors.New("token is not set")
	}

	return token, nil
}
