package requests

import "s21client/gql"

type Variables_GetCredentialsByLogin struct {
	Login string `json:"login"`
}


type Data_GetCredentialsByLogin struct {
	School21_GetCredentialsByLogin School21_GetCredentialsByLogin `json:"school21"`
}

type School21_GetCredentialsByLogin struct {
	GetStudentByLogin_GetCredentialsByLogin GetStudentByLogin_GetCredentialsByLogin `json:"getStudentByLogin"`
	Typename          string            `json:"__typename"`
}

type GetStudentByLogin_GetCredentialsByLogin struct {
	StudentID  string `json:"studentId"`
	UserID     string `json:"userId"`
	SchoolID   string `json:"schoolId"`
	IsActive   bool   `json:"isActive"`
	IsGraduate bool   `json:"isGraduate"`
	Typename   string `json:"__typename"`
}

func (ctx *RequestContext) GetCredentialsByLogin(variables Variables_GetCredentialsByLogin) (Data_GetCredentialsByLogin, error) {
	request := gql.NewQueryRequest[Variables_GetCredentialsByLogin](
		"query getCredentialsByLogin($login: String!) {   school21 {     getStudentByLogin(login: $login) {       studentId       userId       schoolId       isActive       isGraduate       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetCredentialsByLogin](ctx, request)
}