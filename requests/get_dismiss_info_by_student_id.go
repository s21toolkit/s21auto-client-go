package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetDismissInfoByStudentId struct {
	StudentID string `json:"studentId"`
}


type Response_Data_GetDismissInfoByStudentId struct {
	Response_School21_GetDismissInfoByStudentId Response_School21_GetDismissInfoByStudentId `json:"school21"`
}

type Response_School21_GetDismissInfoByStudentId struct {
	GetDismissInfoByStudentID interface{} `json:"getDismissInfoByStudentId"`
	Typename                  string      `json:"__typename"`
}

func (ctx *RequestContext) GetDismissInfoByStudentId(variables Request_Variables_GetDismissInfoByStudentId) (Response_Data_GetDismissInfoByStudentId, error) {
	request := gql.NewQueryRequest[Request_Variables_GetDismissInfoByStudentId](
		"query getDismissInfoByStudentId($studentId: UUID!) {\n  school21 {\n    getDismissInfoByStudentId(studentId: $studentId) {\n      dismissTypeId\n      dismissTs\n      lastStageGroupS21 {\n        waveId\n        waveName\n        eduForm\n        active\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetDismissInfoByStudentId](ctx, request)
}