package requests

import "github.com/s21toolkit/s21client/gql"

type GetStudentStageGroupS21_Variables struct {
	StudentID string `json:"studentId"`
}


type GetStudentStageGroupS21_Data struct {
	Student GetStudentStageGroupS21_Data_Student `json:"student"`
}

type GetStudentStageGroupS21_Data_Student struct {
	GetStageGroupS21PublicProfile GetStudentStageGroupS21_Data_GetStageGroupS21PublicProfile `json:"getStageGroupS21PublicProfile"`
	Typename                      string                        `json:"__typename"`
}

type GetStudentStageGroupS21_Data_GetStageGroupS21PublicProfile struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetStudentStageGroupS21(variables GetStudentStageGroupS21_Variables) (GetStudentStageGroupS21_Data, error) {
	request := gql.NewQueryRequest[GetStudentStageGroupS21_Variables](
		"query getStudentStageGroupS21($studentId: UUID!) {\n  student {\n    getStageGroupS21PublicProfile(studentId: $studentId) {\n      waveId\n      waveName\n      eduForm\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetStudentStageGroupS21_Data](ctx, request)
}