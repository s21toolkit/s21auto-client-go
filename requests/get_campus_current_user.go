package requests

import "s21client/gql"

type Variables_GetCampusCurrentUser struct {
}


type Data_GetCampusCurrentUser struct {
	User_GetCampusCurrentUser User_GetCampusCurrentUser `json:"user"`
}

type User_GetCampusCurrentUser struct {
	GetCurrentUser_GetCampusCurrentUser GetCurrentUser_GetCampusCurrentUser `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type GetCurrentUser_GetCampusCurrentUser struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetCampusCurrentUser(variables Variables_GetCampusCurrentUser) (Data_GetCampusCurrentUser, error) {
	request := gql.NewQueryRequest[Variables_GetCampusCurrentUser](
		"query getCampusCurrentUser {   user {     getCurrentUser {       id       login       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetCampusCurrentUser](ctx, request)
}