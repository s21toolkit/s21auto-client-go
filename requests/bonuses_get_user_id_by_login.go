package requests

import "github.com/s21toolkit/s21client/gql"

type BonusesGetUserIdByLogin_Variables struct {
	Login string `json:"login"`
}


type BonusesGetUserIdByLogin_Data struct {
	Student BonusesGetUserIdByLogin_Data_Student `json:"student"`
}

type BonusesGetUserIdByLogin_Data_Student struct {
	GetUserIDByLogin string `json:"getUserIdByLogin"`
	Typename         string `json:"__typename"`
}


func (ctx *RequestContext) BonusesGetUserIdByLogin(variables BonusesGetUserIdByLogin_Variables) (BonusesGetUserIdByLogin_Data, error) {
	request := gql.NewQueryRequest[BonusesGetUserIdByLogin_Variables](
		"query bonusesGetUserIdByLogin($login: String!) {\n  student {\n    getUserIdByLogin(login: $login)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[BonusesGetUserIdByLogin_Data](ctx, request)
}