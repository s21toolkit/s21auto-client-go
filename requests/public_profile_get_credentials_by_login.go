package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_PublicProfileGetCredentialsByLogin struct {
	Login string `json:"login"`
}


type Response_Data_PublicProfileGetCredentialsByLogin struct {
	Response_School21_PublicProfileGetCredentialsByLogin Response_School21_PublicProfileGetCredentialsByLogin `json:"school21"`
}

type Response_School21_PublicProfileGetCredentialsByLogin struct {
	Response_GetStudentByLogin_PublicProfileGetCredentialsByLogin Response_GetStudentByLogin_PublicProfileGetCredentialsByLogin `json:"getStudentByLogin"`
	Typename          string            `json:"__typename"`
}

type Response_GetStudentByLogin_PublicProfileGetCredentialsByLogin struct {
	StudentID  string `json:"studentId"`
	UserID     string `json:"userId"`
	SchoolID   string `json:"schoolId"`
	IsActive   bool   `json:"isActive"`
	IsGraduate bool   `json:"isGraduate"`
	Typename   string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetCredentialsByLogin(variables Request_Variables_PublicProfileGetCredentialsByLogin) (Response_Data_PublicProfileGetCredentialsByLogin, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileGetCredentialsByLogin](
		"query publicProfileGetCredentialsByLogin($login: String!) {   school21 {     getStudentByLogin(login: $login) {       studentId       userId       schoolId       isActive       isGraduate       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileGetCredentialsByLogin](ctx, request)
}