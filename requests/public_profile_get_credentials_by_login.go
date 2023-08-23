package requests

import "s21client/gql"

type Variables_PublicProfileGetCredentialsByLogin struct {
	Login string `json:"login"`
}


type Data_PublicProfileGetCredentialsByLogin struct {
	School21_PublicProfileGetCredentialsByLogin School21_PublicProfileGetCredentialsByLogin `json:"school21"`
}

type School21_PublicProfileGetCredentialsByLogin struct {
	GetStudentByLogin_PublicProfileGetCredentialsByLogin GetStudentByLogin_PublicProfileGetCredentialsByLogin `json:"getStudentByLogin"`
	Typename          string            `json:"__typename"`
}

type GetStudentByLogin_PublicProfileGetCredentialsByLogin struct {
	StudentID  string `json:"studentId"`
	UserID     string `json:"userId"`
	SchoolID   string `json:"schoolId"`
	IsActive   bool   `json:"isActive"`
	IsGraduate bool   `json:"isGraduate"`
	Typename   string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetCredentialsByLogin(variables Variables_PublicProfileGetCredentialsByLogin) (Data_PublicProfileGetCredentialsByLogin, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetCredentialsByLogin](
		"query publicProfileGetCredentialsByLogin($login: String!) {   school21 {     getStudentByLogin(login: $login) {       studentId       userId       schoolId       isActive       isGraduate       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetCredentialsByLogin](ctx, request)
}