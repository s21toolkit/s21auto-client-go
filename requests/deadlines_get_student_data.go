package requests

import "s21client/gql"

type Request_Variables_DeadlinesGetStudentData struct {
}


type Response_Data_DeadlinesGetStudentData struct {
	Response_Student_DeadlinesGetStudentData Response_Student_DeadlinesGetStudentData `json:"student"`
}

type Response_Student_DeadlinesGetStudentData struct {
	Response_GetStudentProfile_DeadlinesGetStudentData    Response_GetStudentProfile_DeadlinesGetStudentData      `json:"getStudentProfile"`
	Response_GetExperience_DeadlinesGetStudentData        Response_GetExperience_DeadlinesGetStudentData          `json:"getExperience"`
	Response_GetExperienceHistory_DeadlinesGetStudentData []Response_GetExperienceHistory_DeadlinesGetStudentData `json:"getExperienceHistory"`
	Typename             string                 `json:"__typename"`
}

type Response_GetExperience_DeadlinesGetStudentData struct {
	ID       string `json:"id"`
	Value    int64  `json:"value"`
	Response_Level_DeadlinesGetStudentData    Response_Level_DeadlinesGetStudentData  `json:"level"`
	Typename string `json:"__typename"`
}

type Response_Level_DeadlinesGetStudentData struct {
	ID       int64  `json:"id"`
	Response_Range_DeadlinesGetStudentData    Response_Range_DeadlinesGetStudentData  `json:"range"`
	Typename string `json:"__typename"`
}

type Response_Range_DeadlinesGetStudentData struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	RightBorder int64  `json:"rightBorder"`
	LeftBorder  int64  `json:"leftBorder"`
	Typename    string `json:"__typename"`
}

type Response_GetExperienceHistory_DeadlinesGetStudentData struct {
	ID                 string `json:"id"`
	AwardDate          string `json:"awardDate"`
	ExperienceReceived int64  `json:"experienceReceived"`
	Typename           string `json:"__typename"`
}

type Response_GetStudentProfile_DeadlinesGetStudentData struct {
	Response_User_DeadlinesGetStudentData     Response_User_DeadlinesGetStudentData   `json:"user"`
	Typename string `json:"__typename"`
}

type Response_User_DeadlinesGetStudentData struct {
	ID         string `json:"id"`
	Login      string `json:"login"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Typename   string `json:"__typename"`
}

func (ctx *RequestContext) DeadlinesGetStudentData(variables Request_Variables_DeadlinesGetStudentData) (Response_Data_DeadlinesGetStudentData, error) {
	request := gql.NewQueryRequest[Request_Variables_DeadlinesGetStudentData](
		"query deadlinesGetStudentData {   student {     getStudentProfile {       user {         id         login         firstName         middleName         lastName         __typename       }       __typename     }     getExperience {       id       value       level {         id         range {           id           levelCode           rightBorder           leftBorder           __typename         }         __typename       }       __typename     }     getExperienceHistory {       id       awardDate       experienceReceived       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_DeadlinesGetStudentData](ctx, request)
}