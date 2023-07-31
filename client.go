package s21client

import (
	"github.com/go-resty/resty/v2"
)

type Client struct {
	resty *resty.Client

	authProvider AuthProvider
}

func (client *Client) UseAuth(authProvider AuthProvider) {
	client.authProvider = authProvider
}

func (client *Client) applyAuth(request *resty.Request) {
	credentials := client.authProvider.GetAuthCredentials(request.Context())

	request.SetAuthToken(credentials.Token).SetHeader("schoolid", credentials.SchoolId)
}

func (client *Client) R() *resty.Request {
	request := client.resty.R()

	client.applyAuth(request)

	return request
}

func New(authProvider AuthProvider) *Client {
	return &Client{
		authProvider: authProvider,
		resty:        resty.New(),
	}
}
