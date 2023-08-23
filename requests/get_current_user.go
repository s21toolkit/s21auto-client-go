package requests

import "s21client/gql"

type Variables_GetCurrentUser struct {
}


type Data_GetCurrentUser struct {
	User_GetCurrentUser    User_GetCurrentUser    `json:"user"`
	Student_GetCurrentUser Student_GetCurrentUser `json:"student"`
}

type Student_GetCurrentUser struct {
	GetExperience_GetCurrentUser GetExperience_GetCurrentUser `json:"getExperience"`
	Typename      string        `json:"__typename"`
}

type GetExperience_GetCurrentUser struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Level_GetCurrentUser            Level_GetCurrentUser  `json:"level"`
	Typename         string `json:"__typename"`
}

type Level_GetCurrentUser struct {
	ID       int64  `json:"id"`
	Range_GetCurrentUser    Range_GetCurrentUser  `json:"range"`
	Typename string `json:"__typename"`
}

type Range_GetCurrentUser struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type User_GetCurrentUser struct {
	GetCurrentUser_GetCurrentUser GetCurrentUser_GetCurrentUser `json:"getCurrentUser"`
	Typename       string         `json:"__typename"`
}

type GetCurrentUser_GetCurrentUser struct {
	ID                     string `json:"id"`
	AvatarURL              string `json:"avatarUrl"`
	Login                  string `json:"login"`
	FirstName              string `json:"firstName"`
	MiddleName             string `json:"middleName"`
	LastName               string `json:"lastName"`
	CurrentSchoolStudentID string `json:"currentSchoolStudentId"`
	Typename               string `json:"__typename"`
}

func (ctx *RequestContext) GetCurrentUser(variables Variables_GetCurrentUser) (Data_GetCurrentUser, error) {
	request := gql.NewQueryRequest[Variables_GetCurrentUser](
		"query getCurrentUser {   user {     getCurrentUser {       ...CurrentUser       __typename     }     __typename   }   student {     getExperience {       ...CurrentUserExperience       __typename     }     __typename   } }  fragment CurrentUser on User {   id   avatarUrl   login   firstName   middleName   lastName   currentSchoolStudentId   __typename }  fragment CurrentUserExperience on UserExperience {   id   cookiesCount   codeReviewPoints   coinsCount   level {     id     range {       id       levelCode       __typename     }     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetCurrentUser](ctx, request)
}