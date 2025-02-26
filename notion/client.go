package notion

import "github.com/jomei/notionapi"

var client *notionapi.Client

func load() error {
	token, err := getToken()
	if err != nil {
		return err
	}

	client = notionapi.NewClient(notionapi.Token(token))

	return nil
}

func Init() error {
	return load()
}

func GetClient() (*notionapi.Client, error) {
	if client == nil {
		if err := load(); err != nil {
			return nil, err
		}
	}
	return client, nil
}

func ShouldGetClient() *notionapi.Client {
	if client == nil {
		panic("client is not loaded")
	}
	return client
}
