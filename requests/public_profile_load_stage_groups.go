package requests

import "s21client/gql"

type Variables_PublicProfileLoadStageGroups struct {
	UserID   string `json:"userId"`
	SchoolID string `json:"schoolId"`
}


type Data_PublicProfileLoadStageGroups struct {
	School21_PublicProfileLoadStageGroups School21_PublicProfileLoadStageGroups `json:"school21"`
}

type School21_PublicProfileLoadStageGroups struct {
	LoadStudentStageGroupsS21PublicProfile_PublicProfileLoadStageGroups []LoadStudentStageGroupsS21PublicProfile_PublicProfileLoadStageGroups `json:"loadStudentStageGroupsS21PublicProfile"`
	Typename                               string                                   `json:"__typename"`
}

type LoadStudentStageGroupsS21PublicProfile_PublicProfileLoadStageGroups struct {
	StageGroupStudentID string        `json:"stageGroupStudentId"`
	StudentID           string        `json:"studentId"`
	StageGroupS21_PublicProfileLoadStageGroups       StageGroupS21_PublicProfileLoadStageGroups `json:"stageGroupS21"`
	SafeSchool_PublicProfileLoadStageGroups          SafeSchool_PublicProfileLoadStageGroups    `json:"safeSchool"`
	Typename            string        `json:"__typename"`
}

type SafeSchool_PublicProfileLoadStageGroups struct {
	FullName string `json:"fullName"`
	Typename string `json:"__typename"`
}

type StageGroupS21_PublicProfileLoadStageGroups struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Active   bool   `json:"active"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileLoadStageGroups(variables Variables_PublicProfileLoadStageGroups) (Data_PublicProfileLoadStageGroups, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileLoadStageGroups](
		"query publicProfileLoadStageGroups($userId: UUID!, $schoolId: UUID!) {   school21 {     loadStudentStageGroupsS21PublicProfile(userId: $userId, schoolId: $schoolId) {       stageGroupStudentId       studentId       stageGroupS21 {         waveId         waveName         eduForm         active         __typename       }       safeSchool {         fullName         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileLoadStageGroups](ctx, request)
}