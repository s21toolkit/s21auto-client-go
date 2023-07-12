package S21Client

import (
	"context"
	"net/http"
	"school-21/pkg/client/auth"
)

type S21Client struct {
	schoolId string
	token    *auth.S21Token
}

func NewS21Client(user string, password string, ctx context.Context) (*S21Client, error) {
	s := &S21Client{
		schoolId: "6bfe3c56-0211-4fe1-9e59-51616caac4dd",
		token: &auth.S21Token{
			User:     user,
			Password: password,
		},
	}

	err := s.token.Refresh(ctx)
	if err != nil {
		return nil, err
	}

	return s, err
}

func (s *S21Client) Do(r *http.Request) (*http.Response, error) {
	s.token.Apply(r)
	res, err := http.DefaultClient.Do(r)

	return res, err
}
