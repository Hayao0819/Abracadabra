package nautils

import "github.com/jomei/notionapi"

type Client struct {
	c *notionapi.Client
}

func NewClientFromNotionAPIClient(c *notionapi.Client) *Client {
	return &Client{c: c}
}

func (c *Client) RawClient() *notionapi.Client {
	if c == nil {
		panic("client is not loaded")
	}
	return c.c
}

func NewClient(t notionapi.Token, opts ...notionapi.ClientOption) *Client {
	rawClient := notionapi.NewClient(t, opts...)
	return &Client{c: rawClient}
}
