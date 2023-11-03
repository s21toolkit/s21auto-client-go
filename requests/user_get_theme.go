package requests

import "github.com/s21toolkit/s21client/gql"

type UserGetTheme_Variables struct {
	UserID string `json:"userId"`
}


type UserGetTheme_Data struct {
	User UserGetTheme_Data_User `json:"user"`
}

type UserGetTheme_Data_User struct {
	GetUserViewSettings UserGetTheme_Data_GetUserViewSettings `json:"getUserViewSettings"`
	Typename            string              `json:"__typename"`
}

type UserGetTheme_Data_GetUserViewSettings struct {
	IsDarkThemeEnabled bool   `json:"isDarkThemeEnabled"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) UserGetTheme(variables UserGetTheme_Variables) (UserGetTheme_Data, error) {
	request := gql.NewQueryRequest[UserGetTheme_Variables](
		"query userGetTheme($userId: UUID!) {\n  user {\n    getUserViewSettings(userId: $userId) {\n      isDarkThemeEnabled\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[UserGetTheme_Data](ctx, request)
}