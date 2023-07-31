package s21client

import (
	"context"
	"s21client/auth"
)

type defaultAuthProvider struct {
	token    auth.Token
	schoolId string
}

func (provider *defaultAuthProvider) refreshCredentials(ctx context.Context) {
	err := provider.token.Refresh(ctx)

	if err != nil {
		panic(err)
	}

	if provider.schoolId != "" {
		return
	}

	user, err := auth.RequestUserData(provider.token, ctx)

	if err != nil {
		panic(err)
	}

	provider.schoolId = user.Roles[0].SchoolID
}

func (provider *defaultAuthProvider) GetAuthCredentials(ctx context.Context) AuthCredentials {
	provider.refreshCredentials(ctx)

	return AuthCredentials{
		Token:    provider.token.AccessToken,
		SchoolId: provider.schoolId,
	}
}

func DefaultAuth(username, password string) *defaultAuthProvider {
	return &defaultAuthProvider{
		token: auth.Token{
			Username: username,
			Password: password,
		},
	}
}
