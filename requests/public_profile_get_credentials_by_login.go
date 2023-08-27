package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_PublicProfileGetCredentialsByLogin struct {
	Login string `json:"login"`
}


type Data_PublicProfileGetCredentialsByLogin struct {
	Data_School21_PublicProfileGetCredentialsByLogin Data_School21_PublicProfileGetCredentialsByLogin `json:"school21"`
}

type Data_School21_PublicProfileGetCredentialsByLogin struct {
	Data_GetStudentByLogin_PublicProfileGetCredentialsByLogin Data_GetStudentByLogin_PublicProfileGetCredentialsByLogin `json:"getStudentByLogin"`
	Typename          string            `json:"__typename"`
}

type Data_GetStudentByLogin_PublicProfileGetCredentialsByLogin struct {
	StudentID  string `json:"studentId"`
	UserID     string `json:"userId"`
	SchoolID   string `json:"schoolId"`
	IsActive   bool   `json:"isActive"`
	IsGraduate bool   `json:"isGraduate"`
	Typename   string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetCredentialsByLogin(variables Variables_PublicProfileGetCredentialsByLogin) (Data_PublicProfileGetCredentialsByLogin, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetCredentialsByLogin](
		"query publicProfileGetCredentialsByLogin($login: String!) {\n  school21 {\n    getStudentByLogin(login: $login) {\n      studentId\n      userId\n      schoolId\n      isActive\n      isGraduate\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetCredentialsByLogin](ctx, request)
}