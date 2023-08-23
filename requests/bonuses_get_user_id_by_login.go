package requests

import "s21client/gql"

type Variables_BonusesGetUserIdByLogin struct {
	Login string `json:"login"`
}


type Data_BonusesGetUserIdByLogin struct {
	Student_BonusesGetUserIdByLogin Student_BonusesGetUserIdByLogin `json:"student"`
}

type Student_BonusesGetUserIdByLogin struct {
	GetUserIDByLogin string `json:"getUserIdByLogin"`
	Typename         string `json:"__typename"`
}

func (ctx *RequestContext) BonusesGetUserIdByLogin(variables Variables_BonusesGetUserIdByLogin) (Data_BonusesGetUserIdByLogin, error) {
	request := gql.NewQueryRequest[Variables_BonusesGetUserIdByLogin](
		"query bonusesGetUserIdByLogin($login: String!) {   student {     getUserIdByLogin(login: $login)     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_BonusesGetUserIdByLogin](ctx, request)
}