package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetUsers struct {
	UserIDS []string `json:"userIds"`
}


type Data_GetUsers struct {
	Data_School21_GetUsers Data_School21_GetUsers `json:"school21"`
}

type Data_School21_GetUsers struct {
	GetUsers []Data_GetUser_GetUsers `json:"getUsers"`
	Typename string    `json:"__typename"`
}

type Data_GetUser_GetUsers struct {
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
		"query getUsers($userIds: [UUID!]!) {\n  school21 {\n    getUsers(userIds: $userIds) {\n      userId\n      login\n      firstName\n      middleName\n      lastName\n      avatarUrl\n      level\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetUsers](ctx, request)
}