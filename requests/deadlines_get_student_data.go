package requests

import "s21client/gql"

type Variables_DeadlinesGetStudentData struct {
}


type Data_DeadlinesGetStudentData struct {
	Student_DeadlinesGetStudentData Student_DeadlinesGetStudentData `json:"student"`
}

type Student_DeadlinesGetStudentData struct {
	GetStudentProfile_DeadlinesGetStudentData    GetStudentProfile_DeadlinesGetStudentData      `json:"getStudentProfile"`
	GetExperience_DeadlinesGetStudentData        GetExperience_DeadlinesGetStudentData          `json:"getExperience"`
	GetExperienceHistory_DeadlinesGetStudentData []GetExperienceHistory_DeadlinesGetStudentData `json:"getExperienceHistory"`
	Typename             string                 `json:"__typename"`
}

type GetExperience_DeadlinesGetStudentData struct {
	ID       string `json:"id"`
	Value    int64  `json:"value"`
	Level_DeadlinesGetStudentData    Level_DeadlinesGetStudentData  `json:"level"`
	Typename string `json:"__typename"`
}

type Level_DeadlinesGetStudentData struct {
	ID       int64  `json:"id"`
	Range_DeadlinesGetStudentData    Range_DeadlinesGetStudentData  `json:"range"`
	Typename string `json:"__typename"`
}

type Range_DeadlinesGetStudentData struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	RightBorder int64  `json:"rightBorder"`
	LeftBorder  int64  `json:"leftBorder"`
	Typename    string `json:"__typename"`
}

type GetExperienceHistory_DeadlinesGetStudentData struct {
	ID                 string `json:"id"`
	AwardDate          string `json:"awardDate"`
	ExperienceReceived int64  `json:"experienceReceived"`
	Typename           string `json:"__typename"`
}

type GetStudentProfile_DeadlinesGetStudentData struct {
	User_DeadlinesGetStudentData     User_DeadlinesGetStudentData   `json:"user"`
	Typename string `json:"__typename"`
}

type User_DeadlinesGetStudentData struct {
	ID         string `json:"id"`
	Login      string `json:"login"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Typename   string `json:"__typename"`
}

func (ctx *RequestContext) DeadlinesGetStudentData(variables Variables_DeadlinesGetStudentData) (Data_DeadlinesGetStudentData, error) {
	request := gql.NewQueryRequest[Variables_DeadlinesGetStudentData](
		"query deadlinesGetStudentData {   student {     getStudentProfile {       user {         id         login         firstName         middleName         lastName         __typename       }       __typename     }     getExperience {       id       value       level {         id         range {           id           levelCode           rightBorder           leftBorder           __typename         }         __typename       }       __typename     }     getExperienceHistory {       id       awardDate       experienceReceived       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_DeadlinesGetStudentData](ctx, request)
}