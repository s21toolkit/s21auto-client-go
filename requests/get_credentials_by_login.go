package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetCredentialsByLogin struct {
	Login string `json:"login"`
}


type Response_Data_GetCredentialsByLogin struct {
	Response_School21_GetCredentialsByLogin Response_School21_GetCredentialsByLogin `json:"school21"`
}

type Response_School21_GetCredentialsByLogin struct {
	Response_GetStudentByLogin_GetCredentialsByLogin Response_GetStudentByLogin_GetCredentialsByLogin `json:"getStudentByLogin"`
	Typename          string            `json:"__typename"`
}

type Response_GetStudentByLogin_GetCredentialsByLogin struct {
	StudentID  string `json:"studentId"`
	UserID     string `json:"userId"`
	SchoolID   string `json:"schoolId"`
	IsActive   bool   `json:"isActive"`
	IsGraduate bool   `json:"isGraduate"`
	Typename   string `json:"__typename"`
}

func (ctx *RequestContext) GetCredentialsByLogin(variables Request_Variables_GetCredentialsByLogin) (Response_Data_GetCredentialsByLogin, error) {
	request := gql.NewQueryRequest[Request_Variables_GetCredentialsByLogin](
		"query getCredentialsByLogin($login: String!) {\n  school21 {\n    getStudentByLogin(login: $login) {\n      studentId\n      userId\n      schoolId\n      isActive\n      isGraduate\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetCredentialsByLogin](ctx, request)
}