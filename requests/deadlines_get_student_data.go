package requests

import "github.com/s21toolkit/s21client/gql"

type DeadlinesGetStudentData_Variables struct {
}


type DeadlinesGetStudentData_Data struct {
	Student DeadlinesGetStudentData_Data_Student `json:"student"`
}

type DeadlinesGetStudentData_Data_Student struct {
	GetStudentProfile    DeadlinesGetStudentData_Data_GetStudentProfile      `json:"getStudentProfile"`
	GetExperience        DeadlinesGetStudentData_Data_GetExperience          `json:"getExperience"`
	GetExperienceHistory []DeadlinesGetStudentData_Data_GetExperienceHistory `json:"getExperienceHistory"`
	Typename             string                 `json:"__typename"`
}

type DeadlinesGetStudentData_Data_GetExperience struct {
	ID       string `json:"id"`
	Value    int64  `json:"value"`
	Level    DeadlinesGetStudentData_Data_Level  `json:"level"`
	Typename string `json:"__typename"`
}

type DeadlinesGetStudentData_Data_Level struct {
	ID       int64  `json:"id"`
	Range    DeadlinesGetStudentData_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type DeadlinesGetStudentData_Data_Range struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	RightBorder int64  `json:"rightBorder"`
	LeftBorder  int64  `json:"leftBorder"`
	Typename    string `json:"__typename"`
}

type DeadlinesGetStudentData_Data_GetExperienceHistory struct {
	ID                 string `json:"id"`
	AwardDate          string `json:"awardDate"`
	ExperienceReceived int64  `json:"experienceReceived"`
	Typename           string `json:"__typename"`
}

type DeadlinesGetStudentData_Data_GetStudentProfile struct {
	User     DeadlinesGetStudentData_Data_User   `json:"user"`
	Typename string `json:"__typename"`
}

type DeadlinesGetStudentData_Data_User struct {
	ID         string `json:"id"`
	Login      string `json:"login"`
	FirstName  string `json:"firstName"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Typename   string `json:"__typename"`
}


func (ctx *RequestContext) DeadlinesGetStudentData(variables DeadlinesGetStudentData_Variables) (DeadlinesGetStudentData_Data, error) {
	request := gql.NewQueryRequest[DeadlinesGetStudentData_Variables](
		"query deadlinesGetStudentData {\n  student {\n    getStudentProfile {\n      user {\n        id\n        login\n        firstName\n        middleName\n        lastName\n        __typename\n      }\n      __typename\n    }\n    getExperience {\n      id\n      value\n      level {\n        id\n        range {\n          id\n          levelCode\n          rightBorder\n          leftBorder\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    getExperienceHistory {\n      id\n      awardDate\n      experienceReceived\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[DeadlinesGetStudentData_Data](ctx, request)
}