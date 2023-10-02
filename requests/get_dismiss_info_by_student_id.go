package requests

import "github.com/s21toolkit/s21client/gql"

type GetDismissInfoByStudentId_Variables struct {
	StudentID string `json:"studentId"`
}


type GetDismissInfoByStudentId_Data struct {
	School21 GetDismissInfoByStudentId_Data_School21 `json:"school21"`
}

type GetDismissInfoByStudentId_Data_School21 struct {
	GetDismissInfoByStudentID interface{} `json:"getDismissInfoByStudentId"`
	Typename                  string      `json:"__typename"`
}


func (ctx *RequestContext) GetDismissInfoByStudentId(variables GetDismissInfoByStudentId_Variables) (GetDismissInfoByStudentId_Data, error) {
	request := gql.NewQueryRequest[GetDismissInfoByStudentId_Variables](
		"query getDismissInfoByStudentId($studentId: UUID!) {\n  school21 {\n    getDismissInfoByStudentId(studentId: $studentId) {\n      dismissTypeId\n      dismissTs\n      lastStageGroupS21 {\n        waveId\n        waveName\n        eduForm\n        active\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetDismissInfoByStudentId_Data](ctx, request)
}