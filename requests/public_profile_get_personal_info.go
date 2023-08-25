package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_PublicProfileGetPersonalInfo struct {
	UserID    string `json:"userId"`
	StudentID string `json:"studentId"`
	SchoolID  string `json:"schoolId"`
	Login     string `json:"login"`
}


type Response_Data_PublicProfileGetPersonalInfo struct {
	Response_Student_PublicProfileGetPersonalInfo Response_Student_PublicProfileGetPersonalInfo `json:"student"`
	Response_User_PublicProfileGetPersonalInfo    Response_User_PublicProfileGetPersonalInfo    `json:"user"`
}

type Response_Student_PublicProfileGetPersonalInfo struct {
	GetAvatarByUserID                 string                            `json:"getAvatarByUserId"`
	Response_GetStageGroupS21PublicProfile_PublicProfileGetPersonalInfo     Response_GetStageGroupS21PublicProfile_PublicProfileGetPersonalInfo     `json:"getStageGroupS21PublicProfile"`
	Response_GetExperiencePublicProfile_PublicProfileGetPersonalInfo        Response_GetExperiencePublicProfile_PublicProfileGetPersonalInfo        `json:"getExperiencePublicProfile"`
	GetEmailbyUserID                  string                            `json:"getEmailbyUserId"`
	GetWorkstationByLogin             interface{}                       `json:"getWorkstationByLogin"`
	GetClassRoomByLogin               interface{}                       `json:"getClassRoomByLogin"`
	Response_GetFeedbackStatisticsAverageScore_PublicProfileGetPersonalInfo Response_GetFeedbackStatisticsAverageScore_PublicProfileGetPersonalInfo `json:"getFeedbackStatisticsAverageScore"`
	Typename                          string                            `json:"__typename"`
}

type Response_GetExperiencePublicProfile_PublicProfileGetPersonalInfo struct {
	Value            int64  `json:"value"`
	Response_Level_PublicProfileGetPersonalInfo            Response_Level_PublicProfileGetPersonalInfo  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CoinsCount       int64  `json:"coinsCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Response_Level_PublicProfileGetPersonalInfo struct {
	LevelCode int64  `json:"levelCode"`
	Response_Range_PublicProfileGetPersonalInfo     Response_Range_PublicProfileGetPersonalInfo  `json:"range"`
	Typename  string `json:"__typename"`
}

type Response_Range_PublicProfileGetPersonalInfo struct {
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}

type Response_GetFeedbackStatisticsAverageScore_PublicProfileGetPersonalInfo struct {
	CountFeedback        int64                  `json:"countFeedback"`
	Response_FeedbackAverageScore_PublicProfileGetPersonalInfo []Response_FeedbackAverageScore_PublicProfileGetPersonalInfo `json:"feedbackAverageScore"`
	Typename             string                 `json:"__typename"`
}

type Response_FeedbackAverageScore_PublicProfileGetPersonalInfo struct {
	CategoryCode string `json:"categoryCode"`
	CategoryName string `json:"categoryName"`
	Value        string `json:"value"`
	Typename     string `json:"__typename"`
}

type Response_GetStageGroupS21PublicProfile_PublicProfileGetPersonalInfo struct {
	WaveID   int64  `json:"waveId"`
	WaveName string `json:"waveName"`
	EduForm  string `json:"eduForm"`
	Typename string `json:"__typename"`
}

type Response_User_PublicProfileGetPersonalInfo struct {
	Response_GetSchool_PublicProfileGetPersonalInfo Response_GetSchool_PublicProfileGetPersonalInfo `json:"getSchool"`
	Typename  string    `json:"__typename"`
}

type Response_GetSchool_PublicProfileGetPersonalInfo struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	ShortName string `json:"shortName"`
	Address   string `json:"address"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetPersonalInfo(variables Request_Variables_PublicProfileGetPersonalInfo) (Response_Data_PublicProfileGetPersonalInfo, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileGetPersonalInfo](
		"query publicProfileGetPersonalInfo($userId: UUID!, $studentId: UUID!, $login: String!, $schoolId: UUID!) {\n  student {\n    getAvatarByUserId(userId: $userId)\n    getStageGroupS21PublicProfile(studentId: $studentId) {\n      waveId\n      waveName\n      eduForm\n      __typename\n    }\n    getExperiencePublicProfile(userId: $userId) {\n      value\n      level {\n        levelCode\n        range {\n          leftBorder\n          rightBorder\n          __typename\n        }\n        __typename\n      }\n      cookiesCount\n      coinsCount\n      codeReviewPoints\n      __typename\n    }\n    getEmailbyUserId(userId: $userId)\n    getWorkstationByLogin(login: $login) {\n      workstationId\n      hostName\n      row\n      number\n      __typename\n    }\n    getClassRoomByLogin(login: $login) {\n      id\n      number\n      floor\n      __typename\n    }\n    getFeedbackStatisticsAverageScore(studentId: $studentId) {\n      countFeedback\n      feedbackAverageScore {\n        categoryCode\n        categoryName\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  user {\n    getSchool(schoolId: $schoolId) {\n      id\n      fullName\n      shortName\n      address\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileGetPersonalInfo](ctx, request)
}