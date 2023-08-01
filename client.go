package s21client

import (
	"s21client/requests"

	"github.com/go-resty/resty/v2"
)

var S21GqlUrl = "https://edu.21-school.ru/services/graphql"

type Client struct {
	gqlUrl string

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

func (client *Client) R() *requests.RequestContext {
	request := client.resty.R()

	return requests.NewRequestContext(request, client.gqlUrl)
}

func New(authProvider AuthProvider) *Client {
	client := &Client{
		gqlUrl:       S21GqlUrl,
		authProvider: authProvider,
		resty:        resty.New(),
	}

	client.resty.OnBeforeRequest(func(c *resty.Client, r *resty.Request) error {
		return client.applyAuth(r)
	})

	return client
}
