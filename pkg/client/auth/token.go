package auth

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"

	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

const (
	kkBaseUrl         = "https://auth.sberclass.ru/auth/realms/EduPowerKeycloak"
	cookieUrlTemplate = kkBaseUrl + "/protocol/openid-connect/auth?client_id=school21&redirect_uri=https%%3A%%2F%%2Fedu.21-school.ru%%2F&state=%s&response_mode=fragment&response_type=code&scope=openid&nonce=%s"
	tokenUrl          = kkBaseUrl + "/protocol/openid-connect/token"
)

var (
	loginActionRegex = regexp.MustCompile(`(https:\/\/.+?)"`)
	oauthCodeRegex   = regexp.MustCompile(`code=(.+)[&$]?`)
)

type JSONToken struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int64  `json:"expires_in"`
	RefreshExpiresIn int64  `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	IDToken          string `json:"id_token"`
	NotBeforePolicy  int64  `json:"not-before-policy"`
	SessionState     string `json:"session_state"`
	Scope            string `json:"scope"`
}

type S21Token struct {
	token    string
	code     string
	User     string
	Password string
	cookies  []*http.Cookie

	expires_in int64
	issued_at  int64
}

func (t *S21Token) Refresh(ctx context.Context) (err error) {
	if !(t.code == "" || (time.Now().Unix()-t.issued_at) > t.expires_in) {
		return
	}

	err = t.getAuthCode(t.User, t.Password, ctx)
	if err != nil {
		return err
	}

	client := resty.New()
	client.SetCookies(t.cookies)

	t.issued_at = time.Now().Unix()
	res, err := client.R().SetFormData(map[string]string{
		"code":         t.code,
		"grant_type":   "authorization_code",
		"client_id":    "school21",
		"redirect_uri": "https://edu.21-school.ru/",
	}).Post(tokenUrl)
	if err != nil {
		return err
	}

	var r JSONToken
	err = json.Unmarshal(res.Body(), &r)
	if err != nil {
		return err
	}

	t.expires_in = r.ExpiresIn
	t.token = r.AccessToken

	return
}

func (t *S21Token) Apply(r *http.Request) (*http.Request, error) {
	err := t.Refresh(r.Context())
	if err != nil {
		return r, err
	}

	for _, c := range t.cookies {
		r.AddCookie(c)
	}
	r.Header.Add("Authorization", "Bearer "+t.token)

	return r, err
}

func (t *S21Token) getAuthCode(user string, password string, ctx context.Context) (err error) {
	state := uuid.New().String()
	nonce := uuid.New().String()

	client := resty.New()
	client.SetRedirectPolicy(resty.NoRedirectPolicy())

	res, err := client.R().SetContext(ctx).Get(fmt.Sprintf(cookieUrlTemplate, state, nonce))
	if err != nil {
		return
	}
	client.SetCookies(res.Cookies())

	loginAction := loginActionRegex.FindString(string(res.Body()))
	loginAction = strings.ReplaceAll(loginAction[0:len(loginAction)-1], "amp;", "")

	res, err = client.R().SetContext(ctx).SetFormData(map[string]string{
		"username": user,
		"password": password,
	}).Post(loginAction)
	if res.StatusCode() != 302 && err != nil {
		return
	}
	client.SetCookies(res.Cookies())

	location := res.Header().Get("location")
	res, err = client.R().SetContext(ctx).Post(location)
	if res.StatusCode() != 302 && err != nil {
		return
	}

	location = res.Header().Get("location")
	t.code = oauthCodeRegex.FindString(location)[5:]
	t.cookies = client.Cookies

	err = nil

	return
}
