package requests

import "github.com/s21toolkit/s21client/gql"

type GetCampusCurrentUser_Variables struct {
}


type GetCampusCurrentUser_Data struct {
	User GetCampusCurrentUser_Data_User `json:"user"`
}

type GetCampusCurrentUser_Data_User struct {
	GetCurrentUser GetCampusCurrentUser_Data_GetCurrentUser `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type GetCampusCurrentUser_Data_GetCurrentUser struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetCampusCurrentUser(variables GetCampusCurrentUser_Variables) (GetCampusCurrentUser_Data, error) {
	request := gql.NewQueryRequest[GetCampusCurrentUser_Variables](
		"query getCampusCurrentUser {\n  user {\n    getCurrentUser {\n      id\n      login\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetCampusCurrentUser_Data](ctx, request)
}