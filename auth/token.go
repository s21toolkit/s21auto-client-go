package auth

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"s21client/util"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

type tokenResponse struct {
	Error            string `json:"error"`
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

type Token struct {
	AccessToken string

	Code string

	Username string
	Password string

	Cookies []*http.Cookie

	IssueTime  int64
	ExpiryTime int64
}

const (
	kсBaseUrl         = "https://auth.sberclass.ru/auth/realms/EduPowerKeycloak"
	cookieUrlTemplate = kсBaseUrl + "/protocol/openid-connect/auth?client_id=school21&redirect_uri=https%%3A%%2F%%2Fedu.21-school.ru%%2F&state=%s&response_mode=fragment&response_type=code&scope=openid&nonce=%s"
	tokenUrl          = kсBaseUrl + "/protocol/openid-connect/token"
)

var (
	loginActionPattern = regexp.MustCompile(`(?P<LoginActionURL>https:\/\/.+?)"`)
	oauthCodePattern   = regexp.MustCompile(`code=(?P<OAuthCode>.+)[&$]?`)
)

func getLoginActionUrl(data []byte) string {
	rawUrl := loginActionPattern.FindStringSubmatch(string(data))[loginActionPattern.SubexpIndex("LoginActionURL")]

	return strings.ReplaceAll(rawUrl, "amp;", "")
}

func getAuthData(username, password string, ctx context.Context) (authCode string, kcCookies []*http.Cookie, err error) {
	state := uuid.New().String()
	nonce := uuid.New().String()

	client := resty.New()
	client.SetRedirectPolicy(resty.NoRedirectPolicy())

	cookieUrl := fmt.Sprintf(cookieUrlTemplate, state, nonce)

	res, err := client.R().SetContext(ctx).Get(cookieUrl)

	if err != nil {
		return
	}

	client.SetCookies(res.Cookies())

	loginAction := getLoginActionUrl(res.Body())

	res, err = client.R().SetContext(ctx).SetFormData(map[string]string{
		"username": username,
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

	authCode = oauthCodePattern.FindStringSubmatch(location)[oauthCodePattern.SubexpIndex("OAuthCode")]
	kcCookies = client.Cookies

	err = nil

	return
}

func (token *Token) Refresh(ctx context.Context) (err error) {
	if !(token.Code == "" || (time.Now().Unix()-token.IssueTime) > token.ExpiryTime) {
		return
	}

	token.Code, token.Cookies, err = getAuthData(token.Username, token.Password, ctx)

	if err != nil {
		return
	}

	token.IssueTime = time.Now().Unix()

	client := resty.New()
	client.SetCookies(token.Cookies)

	res, err := client.R().SetContext(ctx).SetFormData(map[string]string{
		"code":         token.Code,
		"grant_type":   "authorization_code",
		"client_id":    "school21",
		"redirect_uri": "https://edu.21-school.ru/",
	}).Post(tokenUrl)

	if err != nil {
		return
	}

	tokenResponse, err := util.UnmarshalJson[tokenResponse](res.Body())

	if err != nil {
		return
	}

	if tokenResponse.Error != "" {
		err = fmt.Errorf("unable to get access token: %s", tokenResponse.Error)

		return
	}

	token.ExpiryTime = tokenResponse.ExpiresIn
	token.AccessToken = tokenResponse.AccessToken

	return
}

func RequestToken(username, password string, ctx context.Context) (token Token, err error) {
	token.Username, token.Password = username, password

	err = token.Refresh(ctx)

	return
}
