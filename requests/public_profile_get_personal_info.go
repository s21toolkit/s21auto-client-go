package requests

import "s21client/gql"

type Variables_PublicProfileGetPersonalInfo struct {
	UserID    string `json:"userId"`
	StudentID string `json:"studentId"`
	SchoolID  string `json:"schoolId"`
	Login     string `json:"login"`
}


type Data_PublicProfileGetPersonalInfo struct {
	Student_PublicProfileGetPersonalInfo Student_PublicProfileGetPersonalInfo `json:"student"`
	User_PublicProfileGetPersonalInfo    User_PublicProfileGetPersonalInfo    `json:"user"`
}

type Student_PublicProfileGetPersonalInfo struct {
	GetAvatarByUserID                 string                            `json:"getAvatarByUserId"`
	GetStageGroupS21PublicProfile_PublicProfileGetPersonalInfo     GetStageGroupS21PublicProfile_PublicProfileGetPersonalInfo     `json:"getStageGroupS21PublicProfile"`
	GetExperiencePublicProfile_PublicProfileGetPersonalInfo        GetExperiencePublicProfile_PublicProfileGetPersonalInfo        `json:"getExperiencePublicProfile"`
	GetEmailbyUserID                  string                            `json:"getEmailbyUserId"`
	GetWorkstationByLogin             interface{}                       `json:"getWorkstationByLogin"`
	GetClassRoomByLogin               interface{}                       `json:"getClassRoomByLogin"`
	GetFeedbackStatisticsAverageScore_PublicProfileGetPersonalInfo GetFeedbackStatisticsAverageScore_PublicProfileGetPersonalInfo `json:"getFeedbackStatisticsAverageScore"`
	Typename                          string                            `json:"__typename"`
}

type GetExperiencePublicProfile_PublicProfileGetPersonalInfo struct {
	Value            int64  `json:"value"`
	Level_PublicProfileGetPersonalInfo            Level_PublicProfileGetPersonalInfo  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CoinsCount       int64  `json:"coinsCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Level_PublicProfileGetPersonalInfo struct {
	LevelCode int64  `json:"levelCode"`
	Range_PublicProfileGetPersonalInfo     Range_PublicProfileGetPersonalInfo  `json:"range"`
	Typename  string `json:"__typename"`
}

type Range_PublicProfileGetPersonalInfo struct {
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}

type GetFeedbackStatisticsAverageScore_PublicProfileGetPersonalInfo struct {
	CountFeedback        int64                  `json:"countFeedback"`
	FeedbackAverageScore_PublicProfileGetPersonalInfo []FeedbackAverageScore_PublicProfileGetPersonalInfo `json:"feedbackAverageScore"`
	Typename             string                 `json:"__typename"`
}

type FeedbackAverageScore_PublicProfileGetPersonalInfo struct {
	CategoryCode string `json:"categoryCode"`
	CategoryName string `json:"categoryName"`
	Value        string `json:"value"`
	Typename     string `json:"__typename"`
}

type GetStageGroupS21PublicProfile_PublicProfileGetPersonalInfo struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Typename string `json:"__typename"`
}

type User_PublicProfileGetPersonalInfo struct {
	GetSchool_PublicProfileGetPersonalInfo GetSchool_PublicProfileGetPersonalInfo `json:"getSchool"`
	Typename  string    `json:"__typename"`
}

type GetSchool_PublicProfileGetPersonalInfo struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
	Address   string `json:"address"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetPersonalInfo(variables Variables_PublicProfileGetPersonalInfo) (Data_PublicProfileGetPersonalInfo, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetPersonalInfo](
		"query publicProfileGetPersonalInfo($userId: UUID!, $studentId: UUID!, $login: String!, $schoolId: UUID!) {   student {     getAvatarByUserId(userId: $userId)     getStageGroupS21PublicProfile(studentId: $studentId) {       waveId       waveName       eduForm       __typename     }     getExperiencePublicProfile(userId: $userId) {       value       level {         levelCode         range {           leftBorder           rightBorder           __typename         }         __typename       }       cookiesCount       coinsCount       codeReviewPoints       __typename     }     getEmailbyUserId(userId: $userId)     getWorkstationByLogin(login: $login) {       workstationId       hostName       row       number       __typename     }     getClassRoomByLogin(login: $login) {       id       number       floor       __typename     }     getFeedbackStatisticsAverageScore(studentId: $studentId) {       countFeedback       feedbackAverageScore {         categoryCode         categoryName         value         __typename       }       __typename     }     __typename   }   user {     getSchool(schoolId: $schoolId) {       id       fullName       shortName       address       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetPersonalInfo](ctx, request)
}