package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarGetModule_Variables struct {
	ModuleID string `json:"moduleId"`
}


type CalendarGetModule_Data struct {
	Student CalendarGetModule_Data_Student `json:"student"`
}

type CalendarGetModule_Data_Student struct {
	GetModuleByID CalendarGetModule_Data_GetModuleByID `json:"getModuleById"`
	Typename      string        `json:"__typename"`
}

type CalendarGetModule_Data_GetModuleByID struct {
	ID                string      `json:"id"`
	ModuleTitle       string      `json:"moduleTitle"`
	SubjectTitle      string      `json:"subjectTitle"`
	GoalExecutionType string      `json:"goalExecutionType"`
	Trajectory        CalendarGetModule_Data_Trajectory  `json:"trajectory"`
	CurrentTask       CalendarGetModule_Data_CurrentTask `json:"currentTask"`
	Typename          string      `json:"__typename"`
}

type CalendarGetModule_Data_CurrentTask struct {
	ID       string          `json:"id"`
	Task     CalendarGetModule_Data_CurrentTaskTask `json:"task"`
	Typename string          `json:"__typename"`
}

type CalendarGetModule_Data_CurrentTaskTask struct {
	ID             string `json:"id"`
	AssignmentType string `json:"assignmentType"`
	Typename       string `json:"__typename"`
}

type CalendarGetModule_Data_Trajectory struct {
	ID       string  `json:"id"`
	Levels   []CalendarGetModule_Data_Level `json:"levels"`
	Typename string  `json:"__typename"`
}

type CalendarGetModule_Data_Level struct {
	ID           string        `json:"id"`
	GoalElements []CalendarGetModule_Data_GoalElement `json:"goalElements"`
	Typename     string        `json:"__typename"`
}

type CalendarGetModule_Data_GoalElement struct {
	ID       string  `json:"id"`
	Points   []CalendarGetModule_Data_Point `json:"points"`
	Typename string  `json:"__typename"`
}

type CalendarGetModule_Data_Point struct {
	ID          string      `json:"id"`
	StudentTask CalendarGetModule_Data_StudentTask `json:"studentTask"`
	Typename    string      `json:"__typename"`
}

type CalendarGetModule_Data_StudentTask struct {
	ID         string          `json:"id"`
	TaskID     string          `json:"taskId"`
	Task       CalendarGetModule_Data_StudentTaskTask `json:"task"`
	LastAnswer CalendarGetModule_Data_LastAnswer      `json:"lastAnswer"`
	Typename   string          `json:"__typename"`
}

type CalendarGetModule_Data_LastAnswer struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type CalendarGetModule_Data_StudentTaskTask struct {
	ID                              string                          `json:"id"`
	StudentTaskAdditionalAttributes CalendarGetModule_Data_StudentTaskAdditionalAttributes `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type CalendarGetModule_Data_StudentTaskAdditionalAttributes struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetModule(variables CalendarGetModule_Variables) (CalendarGetModule_Data, error) {
	request := gql.NewQueryRequest[CalendarGetModule_Variables](
		"query calendarGetModule($moduleId: ID!) {\n  student {\n    getModuleById(goalId: $moduleId) {\n      id\n      moduleTitle\n      subjectTitle\n      goalExecutionType\n      trajectory {\n        ...CalendarModuleTrajectory\n        __typename\n      }\n      currentTask {\n        id\n        task {\n          id\n          assignmentType\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarModuleTrajectory on PersonalTrajectory {\n  id\n  levels {\n    id\n    goalElements {\n      id\n      points {\n        id\n        studentTask {\n          ...CalendarStudentTask\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarStudentTask on StudentTask {\n  id\n  taskId\n  task {\n    id\n    studentTaskAdditionalAttributes {\n      ...CalendarStudentTaskAdditionalAttributes\n      __typename\n    }\n    __typename\n  }\n  lastAnswer {\n    id\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarStudentTaskAdditionalAttributes on StudentTaskAdditionalAttributes {\n  cookiesCount\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CalendarGetModule_Data](ctx, request)
}