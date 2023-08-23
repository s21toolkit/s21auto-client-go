package requests

import "s21client/gql"

type Variables_GetStudentStageGroupS21 struct {
	StudentID string `json:"studentId"`
}


type Data_GetStudentStageGroupS21 struct {
	Student_GetStudentStageGroupS21 Student_GetStudentStageGroupS21 `json:"student"`
}

type Student_GetStudentStageGroupS21 struct {
	GetStageGroupS21PublicProfile_GetStudentStageGroupS21 GetStageGroupS21PublicProfile_GetStudentStageGroupS21 `json:"getStageGroupS21PublicProfile"`
	Typename                      string                        `json:"__typename"`
}

type GetStageGroupS21PublicProfile_GetStudentStageGroupS21 struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetStudentStageGroupS21(variables Variables_GetStudentStageGroupS21) (Data_GetStudentStageGroupS21, error) {
	request := gql.NewQueryRequest[Variables_GetStudentStageGroupS21](
		"query getStudentStageGroupS21($studentId: UUID!) {   student {     getStageGroupS21PublicProfile(studentId: $studentId) {       waveId       waveName       eduForm       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetStudentStageGroupS21](ctx, request)
}