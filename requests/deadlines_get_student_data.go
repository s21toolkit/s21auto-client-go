package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_DeadlinesGetStudentData struct {
}


type Data_DeadlinesGetStudentData struct {
	Data_Student_DeadlinesGetStudentData Data_Student_DeadlinesGetStudentData `json:"student"`
}

type Data_Student_DeadlinesGetStudentData struct {
	Data_GetStudentProfile_DeadlinesGetStudentData    Data_GetStudentProfile_DeadlinesGetStudentData      `json:"getStudentProfile"`
	Data_GetExperience_DeadlinesGetStudentData        Data_GetExperience_DeadlinesGetStudentData          `json:"getExperience"`
	Data_GetExperienceHistory_DeadlinesGetStudentData []Data_GetExperienceHistory_DeadlinesGetStudentData `json:"getExperienceHistory"`
	Typename             string                 `json:"__typename"`
}

type Data_GetExperience_DeadlinesGetStudentData struct {
	ID       string `json:"id"`
	Value    int64  `json:"value"`
	Data_Level_DeadlinesGetStudentData    Data_Level_DeadlinesGetStudentData  `json:"level"`
	Typename string `json:"__typename"`
}

type Data_Level_DeadlinesGetStudentData struct {
	ID       int64  `json:"id"`
	Data_Range_DeadlinesGetStudentData    Data_Range_DeadlinesGetStudentData  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_DeadlinesGetStudentData struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	RightBorder int64  `json:"rightBorder"`
	LeftBorder  int64  `json:"leftBorder"`
	Typename    string `json:"__typename"`
}

type Data_GetExperienceHistory_DeadlinesGetStudentData struct {
	ID                 string `json:"id"`
	AwardDate          string `json:"awardDate"`
	ExperienceReceived int64  `json:"experienceReceived"`
	Typename           string `json:"__typename"`
}

type Data_GetStudentProfile_DeadlinesGetStudentData struct {
	Data_User_DeadlinesGetStudentData     Data_User_DeadlinesGetStudentData   `json:"user"`
	Typename string `json:"__typename"`
}

type Data_User_DeadlinesGetStudentData struct {
	ID         string `json:"id"`
	Login      string `json:"login"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Typename   string `json:"__typename"`
}


func (ctx *RequestContext) DeadlinesGetStudentData(variables Variables_DeadlinesGetStudentData) (Data_DeadlinesGetStudentData, error) {
	request := gql.NewQueryRequest[Variables_DeadlinesGetStudentData](
		"query deadlinesGetStudentData {\n  student {\n    getStudentProfile {\n      user {\n        id\n        login\n        firstName\n        middleName\n        lastName\n        __typename\n      }\n      __typename\n    }\n    getExperience {\n      id\n      value\n      level {\n        id\n        range {\n          id\n          levelCode\n          rightBorder\n          leftBorder\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    getExperienceHistory {\n      id\n      awardDate\n      experienceReceived\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_DeadlinesGetStudentData](ctx, request)
}