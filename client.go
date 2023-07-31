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

func (client *Client) applyAuth(request *resty.Request) (err error) {
	credentials, err := client.authProvider.GetAuthCredentials(request.Context())

	if err != nil {
		return
	}

	request.SetAuthToken(credentials.Token).SetHeader("schoolid", credentials.SchoolId)

	return
}

func (client *Client) R() *resty.Request {
	return client.resty.R()
}

func New(authProvider AuthProvider) *Client {
	client := &Client{
		authProvider: authProvider,
		resty:        resty.New(),
	}

	client.resty.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		return client.applyAuth(r)
	})

	return client
}
