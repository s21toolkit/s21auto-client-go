package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileGetPersonalInfo_Variables struct {
	UserID    string `json:"userId"`
	StudentID string `json:"studentId"`
	SchoolID  string `json:"schoolId"`
	Login     string `json:"login"`
}


type PublicProfileGetPersonalInfo_Data struct {
	Student PublicProfileGetPersonalInfo_Data_Student `json:"student"`
	User    PublicProfileGetPersonalInfo_Data_User    `json:"user"`
}

type PublicProfileGetPersonalInfo_Data_Student struct {
	GetAvatarByUserID                 string                            `json:"getAvatarByUserId"`
	GetStageGroupS21PublicProfile     PublicProfileGetPersonalInfo_Data_GetStageGroupS21PublicProfile     `json:"getStageGroupS21PublicProfile"`
	GetExperiencePublicProfile        PublicProfileGetPersonalInfo_Data_GetExperiencePublicProfile        `json:"getExperiencePublicProfile"`
	GetEmailbyUserID                  string                            `json:"getEmailbyUserId"`
	GetWorkstationByLogin             *PublicProfileGetPersonalInfo_Data_GetWorkstationByLogin            `json:"getWorkstationByLogin"`
	GetClassRoomByLogin               *PublicProfileGetPersonalInfo_Data_GetClassRoomByLogin              `json:"getClassRoomByLogin"`
	GetFeedbackStatisticsAverageScore PublicProfileGetPersonalInfo_Data_GetFeedbackStatisticsAverageScore `json:"getFeedbackStatisticsAverageScore"`
	Typename                          string                            `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_GetClassRoomByLogin struct {
	ID       string `json:"id"`
	Number   string `json:"number"`
	Floor    int64  `json:"floor"`
	Typename string `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_GetExperiencePublicProfile struct {
	Value            int64  `json:"value"`
	Level            PublicProfileGetPersonalInfo_Data_Level  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CoinsCount       int64  `json:"coinsCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_Level struct {
	LevelCode int64  `json:"levelCode"`
	Range     PublicProfileGetPersonalInfo_Data_Range  `json:"range"`
	Typename  string `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_Range struct {
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_GetFeedbackStatisticsAverageScore struct {
	CountFeedback        int64                  `json:"countFeedback"`
	FeedbackAverageScore []PublicProfileGetPersonalInfo_Data_FeedbackAverageScore `json:"feedbackAverageScore"`
	Typename             string                 `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_FeedbackAverageScore struct {
	CategoryCode string `json:"categoryCode"`
	CategoryName string `json:"categoryName"`
	Value        string `json:"value"`
	Typename     string `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_GetStageGroupS21PublicProfile struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Typename string `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_GetWorkstationByLogin struct {
	WorkstationID int64  `json:"workstationId"`
	HostName      string `json:"hostName"`
	Row           string `json:"row"`
	Number        int64  `json:"number"`
	Typename      string `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_User struct {
	GetSchool PublicProfileGetPersonalInfo_Data_GetSchool `json:"getSchool"`
	Typename  string    `json:"__typename"`
}

type PublicProfileGetPersonalInfo_Data_GetSchool struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
	Address   string `json:"address"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetPersonalInfo(variables PublicProfileGetPersonalInfo_Variables) (PublicProfileGetPersonalInfo_Data, error) {
	request := gql.NewQueryRequest[PublicProfileGetPersonalInfo_Variables](
		"query publicProfileGetPersonalInfo($userId: UUID!, $studentId: UUID!, $login: String!, $schoolId: UUID!) {\n  student {\n    getAvatarByUserId(userId: $userId)\n    getStageGroupS21PublicProfile(studentId: $studentId) {\n      waveId\n      waveName\n      eduForm\n      __typename\n    }\n    getExperiencePublicProfile(userId: $userId) {\n      value\n      level {\n        levelCode\n        range {\n          leftBorder\n          rightBorder\n          __typename\n        }\n        __typename\n      }\n      cookiesCount\n      coinsCount\n      codeReviewPoints\n      __typename\n    }\n    getEmailbyUserId(userId: $userId)\n    getWorkstationByLogin(login: $login) {\n      workstationId\n      hostName\n      row\n      number\n      __typename\n    }\n    getClassRoomByLogin(login: $login) {\n      id\n      number\n      floor\n      __typename\n    }\n    getFeedbackStatisticsAverageScore(studentId: $studentId) {\n      countFeedback\n      feedbackAverageScore {\n        categoryCode\n        categoryName\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  user {\n    getSchool(schoolId: $schoolId) {\n      id\n      fullName\n      shortName\n      address\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileGetPersonalInfo_Data](ctx, request)
}