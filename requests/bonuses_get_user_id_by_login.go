package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_BonusesGetUserIdByLogin struct {
	Login string `json:"login"`
}


type Data_BonusesGetUserIdByLogin struct {
	Data_Student_BonusesGetUserIdByLogin Data_Student_BonusesGetUserIdByLogin `json:"student"`
}

type Data_Student_BonusesGetUserIdByLogin struct {
	GetUserIDByLogin string `json:"getUserIdByLogin"`
	Typename         string `json:"__typename"`
}


func (ctx *RequestContext) BonusesGetUserIdByLogin(variables Variables_BonusesGetUserIdByLogin) (Data_BonusesGetUserIdByLogin, error) {
	request := gql.NewQueryRequest[Variables_BonusesGetUserIdByLogin](
		"query bonusesGetUserIdByLogin($login: String!) {\n  student {\n    getUserIdByLogin(login: $login)\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_BonusesGetUserIdByLogin](ctx, request)
}