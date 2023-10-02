package requests

import "github.com/s21toolkit/s21client/gql"

type GetCredentialsByLogin_Variables struct {
	Login string `json:"login"`
}


type GetCredentialsByLogin_Data struct {
	School21 GetCredentialsByLogin_Data_School21 `json:"school21"`
}

type GetCredentialsByLogin_Data_School21 struct {
	GetStudentByLogin GetCredentialsByLogin_Data_GetStudentByLogin `json:"getStudentByLogin"`
	Typename          string            `json:"__typename"`
}

type GetCredentialsByLogin_Data_GetStudentByLogin struct {
	StudentID  string `json:"studentId"`
	UserID     string `json:"userId"`
	SchoolID   string `json:"schoolId"`
	IsActive   bool   `json:"isActive"`
	IsGraduate bool   `json:"isGraduate"`
	Typename   string `json:"__typename"`
}


func (ctx *RequestContext) GetCredentialsByLogin(variables GetCredentialsByLogin_Variables) (GetCredentialsByLogin_Data, error) {
	request := gql.NewQueryRequest[GetCredentialsByLogin_Variables](
		"query getCredentialsByLogin($login: String!) {\n  school21 {\n    getStudentByLogin(login: $login) {\n      studentId\n      userId\n      schoolId\n      isActive\n      isGraduate\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetCredentialsByLogin_Data](ctx, request)
}