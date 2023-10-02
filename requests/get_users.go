package requests

import "github.com/s21toolkit/s21client/gql"

type GetUsers_Variables struct {
	UserIDS []string `json:"userIds"`
}


type GetUsers_Data struct {
	School21 GetUsers_Data_School21 `json:"school21"`
}

type GetUsers_Data_School21 struct {
	GetUsers []GetUsers_Data_GetUser `json:"getUsers"`
	Typename string    `json:"__typename"`
}

type GetUsers_Data_GetUser struct {
	UserID     string      `json:"userId"`
	Login      string      `json:"login"`
	FirstName  string      `json:"firstName"`
	MiddleName interface{} `json:"middleName"`
	LastName   string      `json:"lastName"`
	AvatarURL  string      `json:"avatarUrl"`
	Level      int64       `json:"level"`
	Typename   string      `json:"__typename"`
}


func (ctx *RequestContext) GetUsers(variables GetUsers_Variables) (GetUsers_Data, error) {
	request := gql.NewQueryRequest[GetUsers_Variables](
		"query getUsers($userIds: [UUID!]!) {\n  school21 {\n    getUsers(userIds: $userIds) {\n      userId\n      login\n      firstName\n      middleName\n      lastName\n      avatarUrl\n      level\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetUsers_Data](ctx, request)
}