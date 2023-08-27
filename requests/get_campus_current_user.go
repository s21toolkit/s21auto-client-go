package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetCampusCurrentUser struct {
}


type Data_GetCampusCurrentUser struct {
	Data_User_GetCampusCurrentUser Data_User_GetCampusCurrentUser `json:"user"`
}

type Data_User_GetCampusCurrentUser struct {
	Data_GetCurrentUser_GetCampusCurrentUser Data_GetCurrentUser_GetCampusCurrentUser `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type Data_GetCurrentUser_GetCampusCurrentUser struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetCampusCurrentUser(variables Variables_GetCampusCurrentUser) (Data_GetCampusCurrentUser, error) {
	request := gql.NewQueryRequest[Variables_GetCampusCurrentUser](
		"query getCampusCurrentUser {\n  user {\n    getCurrentUser {\n      id\n      login\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetCampusCurrentUser](ctx, request)
}