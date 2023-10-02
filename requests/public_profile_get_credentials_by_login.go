package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileGetCredentialsByLogin_Variables struct {
	Login string `json:"login"`
}


type PublicProfileGetCredentialsByLogin_Data struct {
	School21 PublicProfileGetCredentialsByLogin_Data_School21 `json:"school21"`
}

type PublicProfileGetCredentialsByLogin_Data_School21 struct {
	GetStudentByLogin PublicProfileGetCredentialsByLogin_Data_GetStudentByLogin `json:"getStudentByLogin"`
	Typename          string            `json:"__typename"`
}

type PublicProfileGetCredentialsByLogin_Data_GetStudentByLogin struct {
	StudentID  string `json:"studentId"`
	UserID     string `json:"userId"`
	SchoolID   string `json:"schoolId"`
	IsActive   bool   `json:"isActive"`
	IsGraduate bool   `json:"isGraduate"`
	Typename   string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetCredentialsByLogin(variables PublicProfileGetCredentialsByLogin_Variables) (PublicProfileGetCredentialsByLogin_Data, error) {
	request := gql.NewQueryRequest[PublicProfileGetCredentialsByLogin_Variables](
		"query publicProfileGetCredentialsByLogin($login: String!) {\n  school21 {\n    getStudentByLogin(login: $login) {\n      studentId\n      userId\n      schoolId\n      isActive\n      isGraduate\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileGetCredentialsByLogin_Data](ctx, request)
}