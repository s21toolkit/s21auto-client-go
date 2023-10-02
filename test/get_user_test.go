package test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/s21toolkit/s21client"
	"github.com/s21toolkit/s21client/requests"
)

func loadDotenv(t *testing.T) {
	d, err := os.ReadFile("../.env")
	if err != nil {
		return
	}
	s := strings.Split(string(d), "\n")
	for _, l := range s {
		v := strings.Split(l, "=")
		t.Setenv(v[0], v[1])
	}
}

func TestGetUser(t *testing.T) {
	loadDotenv(t)
	cli := s21client.New(s21client.DefaultAuth(os.Getenv("USERNAME"), os.Getenv("PASSWORD")))

	u, err := cli.R().GetCurrentUser(requests.GetCurrentUser_Variables{})
	if err != nil {
		t.Error(err)
	}

	fmt.Println(u)
}
