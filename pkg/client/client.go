package S21Client

import (
	"context"
	"school-21/pkg/client/auth"
)

type S21Client struct {
	schoolId string
	Token    *auth.S21Token
}

func NewS21Client(user string, password string, ctx context.Context) (*S21Client, error) {
	s := &S21Client{
		schoolId: "6bfe3c56-0211-4fe1-9e59-51616caac4dd",
		Token: &auth.S21Token{
			User:     user,
			Password: password,
		},
	}

	err := s.Token.Refresh(ctx)
	if err != nil {
		return nil, err
	}

	return s, err
}
