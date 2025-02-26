package notion

import (
	"github.com/Hayao0819/Abracadabra/notion/nautils"
	"github.com/jomei/notionapi"
)

var client *nautils.Client

func load() error {
	token, err := getToken()
	if err != nil {
		return err
	}

	client = nautils.NewClient(notionapi.Token(token))

	return nil
}

func Init() error {
	return load()
}

func GetClient() (*nautils.Client, error) {
	if client == nil {
		if err := load(); err != nil {
			return nil, err
		}
	}
	return client, nil
}

func ShouldGetClient() *nautils.Client {
	if client == nil {
		panic("client is not loaded")
	}
	return client
}
