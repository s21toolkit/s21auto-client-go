package requests

import "s21client/gql"

type Variables_GetUsers struct {
	UserIDS []string `json:"userIds"`
}


type Data_GetUsers struct {
	School21_GetUsers School21_GetUsers `json:"school21"`
}

type School21_GetUsers struct {
	GetUsers []GetUser_GetUsers `json:"getUsers"`
	Typename string    `json:"__typename"`
}

type GetUser_GetUsers struct {
	UserID     string      `json:"userId"`
	Login      string      `json:"login"`
	FirstName  string      `json:"firstName"`
	MiddleName interface{} `json:"middleName"`
	LastName   string      `json:"lastName"`
	AvatarURL  string      `json:"avatarUrl"`
	Level      int64       `json:"level"`
	Typename   string      `json:"__typename"`
}

func (ctx *RequestContext) GetUsers(variables Variables_GetUsers) (Data_GetUsers, error) {
	request := gql.NewQueryRequest[Variables_GetUsers](
		"query getUsers($userIds: [UUID!]!) {   school21 {     getUsers(userIds: $userIds) {       userId       login       firstName       middleName       lastName       avatarUrl       level       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetUsers](ctx, request)
}