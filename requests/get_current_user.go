package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetCurrentUser struct {
}


type Response_Data_GetCurrentUser struct {
	Response_User_GetCurrentUser    Response_User_GetCurrentUser    `json:"user"`
	Response_Student_GetCurrentUser Response_Student_GetCurrentUser `json:"student"`
}

type Response_Student_GetCurrentUser struct {
	Response_GetExperience_GetCurrentUser Response_GetExperience_GetCurrentUser `json:"getExperience"`
	Typename      string        `json:"__typename"`
}

type Response_GetExperience_GetCurrentUser struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Response_Level_GetCurrentUser            Response_Level_GetCurrentUser  `json:"level"`
	Typename         string `json:"__typename"`
}

type Response_Level_GetCurrentUser struct {
	ID       int64  `json:"id"`
	Response_Range_GetCurrentUser    Response_Range_GetCurrentUser  `json:"range"`
	Typename string `json:"__typename"`
}

type Response_Range_GetCurrentUser struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type Response_User_GetCurrentUser struct {
	Response_GetCurrentUser_GetCurrentUser Response_GetCurrentUser_GetCurrentUser `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type Response_GetCurrentUser_GetCurrentUser struct {
	ID                     string `json:"id"`
	AvatarURL              string `json:"avatarUrl"`
	Login                  string `json:"login"`
	FirstName              string `json:"firstName"`
	MiddleName             string `json:"middleName"`
	LastName               string `json:"lastName"`
	CurrentSchoolStudentID string `json:"currentSchoolStudentId"`
	Typename               string `json:"__typename"`
}

func (ctx *RequestContext) GetCurrentUser(variables Request_Variables_GetCurrentUser) (Response_Data_GetCurrentUser, error) {
	request := gql.NewQueryRequest[Request_Variables_GetCurrentUser](
		"query getCurrentUser {   user {     getCurrentUser {       ...CurrentUser       __typename     }     __typename   }   student {     getExperience {       ...CurrentUserExperience       __typename     }     __typename   } }  fragment CurrentUser on User {   id   avatarUrl   login   firstName   middleName   lastName   currentSchoolStudentId   __typename }  fragment CurrentUserExperience on UserExperience {   id   cookiesCount   codeReviewPoints   coinsCount   level {     id     range {       id       levelCode       __typename     }     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_GetCurrentUser](ctx, request)
}