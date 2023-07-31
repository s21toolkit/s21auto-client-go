package s21client

import (
	"context"
	"s21client/auth"
)

type defaultAuthProvider struct {
	token    auth.Token
	schoolId string
}

func (provider *defaultAuthProvider) refreshCredentials(ctx context.Context) (err error) {
	err = provider.token.Refresh(ctx)

	if err != nil {
		return err
	}

	if provider.schoolId != "" {
		return
	}

	user, err := auth.RequestUserData(provider.token, ctx)

	if err != nil {
		return
	}

	provider.schoolId = user.Roles[0].SchoolID

	return
}

func (provider *defaultAuthProvider) GetAuthCredentials(ctx context.Context) (credentials AuthCredentials, err error) {
	err = provider.refreshCredentials(ctx)

	if err != nil {
		return
	}

	credentials = AuthCredentials{
		Token:    provider.token.AccessToken,
		SchoolId: provider.schoolId,
	}

	return
}

func DefaultAuth(username, password string) *defaultAuthProvider {
	return &defaultAuthProvider{
		token: auth.Token{
			Username: username,
			Password: password,
		},
	}
}
