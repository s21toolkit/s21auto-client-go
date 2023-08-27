package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetCredentialsByLogin struct {
	Login string `json:"login"`
}


type Data_GetCredentialsByLogin struct {
	Data_School21_GetCredentialsByLogin Data_School21_GetCredentialsByLogin `json:"school21"`
}

type Data_School21_GetCredentialsByLogin struct {
	Data_GetStudentByLogin_GetCredentialsByLogin Data_GetStudentByLogin_GetCredentialsByLogin `json:"getStudentByLogin"`
	Typename          string            `json:"__typename"`
}

type Data_GetStudentByLogin_GetCredentialsByLogin struct {
	StudentID  string `json:"studentId"`
	UserID     string `json:"userId"`
	SchoolID   string `json:"schoolId"`
	IsActive   bool   `json:"isActive"`
	IsGraduate bool   `json:"isGraduate"`
	Typename   string `json:"__typename"`
}


func (ctx *RequestContext) GetCredentialsByLogin(variables Variables_GetCredentialsByLogin) (Data_GetCredentialsByLogin, error) {
	request := gql.NewQueryRequest[Variables_GetCredentialsByLogin](
		"query getCredentialsByLogin($login: String!) {\n  school21 {\n    getStudentByLogin(login: $login) {\n      studentId\n      userId\n      schoolId\n      isActive\n      isGraduate\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetCredentialsByLogin](ctx, request)
}