package requests

import "s21client/gql"

type Request_Variables_BonusesGetUserIdByLogin struct {
	Login string `json:"login"`
}


type Response_Data_BonusesGetUserIdByLogin struct {
	Response_Student_BonusesGetUserIdByLogin Response_Student_BonusesGetUserIdByLogin `json:"student"`
}

type Response_Student_BonusesGetUserIdByLogin struct {
	GetUserIDByLogin string `json:"getUserIdByLogin"`
	Typename         string `json:"__typename"`
}

func (ctx *RequestContext) BonusesGetUserIdByLogin(variables Request_Variables_BonusesGetUserIdByLogin) (Response_Data_BonusesGetUserIdByLogin, error) {
	request := gql.NewQueryRequest[Request_Variables_BonusesGetUserIdByLogin](
		"query bonusesGetUserIdByLogin($login: String!) {   student {     getUserIdByLogin(login: $login)     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_BonusesGetUserIdByLogin](ctx, request)
}