package requests

import "github.com/s21toolkit/s21client/gql"

type GetCurrentUser_Variables struct {
}


type GetCurrentUser_Data struct {
	User    GetCurrentUser_Data_User    `json:"user"`
	Student GetCurrentUser_Data_Student `json:"student"`
}

type GetCurrentUser_Data_Student struct {
	GetExperience GetCurrentUser_Data_GetExperience `json:"getExperience"`
	Typename      string        `json:"__typename"`
}

type GetCurrentUser_Data_GetExperience struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Level            GetCurrentUser_Data_Level  `json:"level"`
	Typename         string `json:"__typename"`
}

type GetCurrentUser_Data_Level struct {
	ID       int64  `json:"id"`
	Range    GetCurrentUser_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type GetCurrentUser_Data_Range struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type GetCurrentUser_Data_User struct {
	GetCurrentUser GetCurrentUser_Data_GetCurrentUser `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type GetCurrentUser_Data_GetCurrentUser struct {
	ID                     string `json:"id"`
	AvatarURL              string `json:"avatarUrl"`
	Login                  string `json:"login"`
	FirstName              string `json:"firstName"`
	MiddleName             string `json:"middleName"`
	LastName               string `json:"lastName"`
	CurrentSchoolStudentID string `json:"currentSchoolStudentId"`
	Typename               string `json:"__typename"`
}


func (ctx *RequestContext) GetCurrentUser(variables GetCurrentUser_Variables) (GetCurrentUser_Data, error) {
	request := gql.NewQueryRequest[GetCurrentUser_Variables](
		"query getCurrentUser {\n  user {\n    getCurrentUser {\n      ...CurrentUser\n      __typename\n    }\n    __typename\n  }\n  student {\n    getExperience {\n      ...CurrentUserExperience\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CurrentUser on User {\n  id\n  avatarUrl\n  login\n  firstName\n  middleName\n  lastName\n  currentSchoolStudentId\n  __typename\n}\n\nfragment CurrentUserExperience on UserExperience {\n  id\n  cookiesCount\n  codeReviewPoints\n  coinsCount\n  level {\n    id\n    range {\n      id\n      levelCode\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetCurrentUser_Data](ctx, request)
}