package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetStudentStageGroupS21 struct {
	StudentID string `json:"studentId"`
}


type Data_GetStudentStageGroupS21 struct {
	Data_Student_GetStudentStageGroupS21 Data_Student_GetStudentStageGroupS21 `json:"student"`
}

type Data_Student_GetStudentStageGroupS21 struct {
	Data_GetStageGroupS21PublicProfile_GetStudentStageGroupS21 Data_GetStageGroupS21PublicProfile_GetStudentStageGroupS21 `json:"getStageGroupS21PublicProfile"`
	Typename                      string                        `json:"__typename"`
}

type Data_GetStageGroupS21PublicProfile_GetStudentStageGroupS21 struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetStudentStageGroupS21(variables Variables_GetStudentStageGroupS21) (Data_GetStudentStageGroupS21, error) {
	request := gql.NewQueryRequest[Variables_GetStudentStageGroupS21](
		"query getStudentStageGroupS21($studentId: UUID!) {\n  student {\n    getStageGroupS21PublicProfile(studentId: $studentId) {\n      waveId\n      waveName\n      eduForm\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetStudentStageGroupS21](ctx, request)
}