package s21client

import "context"

type AuthCredentials struct {
	Token    string
	SchoolId string
}

type AuthProvider interface {
	GetAuthCredentials(ctx context.Context) AuthCredentials
}
