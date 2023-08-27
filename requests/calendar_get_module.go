package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarGetModule struct {
	ModuleID string `json:"moduleId"`
}


type Data_CalendarGetModule struct {
	Data_Student_CalendarGetModule Data_Student_CalendarGetModule `json:"student"`
}

type Data_Student_CalendarGetModule struct {
	Data_GetModuleByID_CalendarGetModule Data_GetModuleByID_CalendarGetModule `json:"getModuleById"`
	Typename      string        `json:"__typename"`
}

type Data_GetModuleByID_CalendarGetModule struct {
	ID                string      `json:"id"`
	ModuleTitle       string      `json:"moduleTitle"`
	SubjectTitle      string      `json:"subjectTitle"`
	GoalExecutionType string      `json:"goalExecutionType"`
	Data_Trajectory_CalendarGetModule        Data_Trajectory_CalendarGetModule  `json:"trajectory"`
	Data_CurrentTask_CalendarGetModule       Data_CurrentTask_CalendarGetModule `json:"currentTask"`
	Typename          string      `json:"__typename"`
}

type Data_CurrentTask_CalendarGetModule struct {
	ID       string          `json:"id"`
	Task     Data_CurrentTaskTask_CalendarGetModule `json:"task"`
	Typename string          `json:"__typename"`
}

type Data_CurrentTaskTask_CalendarGetModule struct {
	ID             string `json:"id"`
	AssignmentType string `json:"assignmentType"`
	Typename       string `json:"__typename"`
}

type Data_Trajectory_CalendarGetModule struct {
	ID       string  `json:"id"`
	Levels   []Data_Level_CalendarGetModule `json:"levels"`
	Typename string  `json:"__typename"`
}

type Data_Level_CalendarGetModule struct {
	ID           string        `json:"id"`
	GoalElements []Data_GoalElement_CalendarGetModule `json:"goalElements"`
	Typename     string        `json:"__typename"`
}

type Data_GoalElement_CalendarGetModule struct {
	ID       string  `json:"id"`
	Points   []Data_Point_CalendarGetModule `json:"points"`
	Typename string  `json:"__typename"`
}

type Data_Point_CalendarGetModule struct {
	ID          string      `json:"id"`
	Data_StudentTask_CalendarGetModule Data_StudentTask_CalendarGetModule `json:"studentTask"`
	Typename    string      `json:"__typename"`
}

type Data_StudentTask_CalendarGetModule struct {
	ID         string          `json:"id"`
	TaskID     string          `json:"taskId"`
	Task       Data_StudentTaskTask_CalendarGetModule `json:"task"`
	Data_LastAnswer_CalendarGetModule Data_LastAnswer_CalendarGetModule      `json:"lastAnswer"`
	Typename   string          `json:"__typename"`
}

type Data_LastAnswer_CalendarGetModule struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type Data_StudentTaskTask_CalendarGetModule struct {
	ID                              string                          `json:"id"`
	Data_StudentTaskAdditionalAttributes_CalendarGetModule Data_StudentTaskAdditionalAttributes_CalendarGetModule `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type Data_StudentTaskAdditionalAttributes_CalendarGetModule struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetModule(variables Variables_CalendarGetModule) (Data_CalendarGetModule, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetModule](
		"query calendarGetModule($moduleId: ID!) {\n  student {\n    getModuleById(goalId: $moduleId) {\n      id\n      moduleTitle\n      subjectTitle\n      goalExecutionType\n      trajectory {\n        ...CalendarModuleTrajectory\n        __typename\n      }\n      currentTask {\n        id\n        task {\n          id\n          assignmentType\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarModuleTrajectory on PersonalTrajectory {\n  id\n  levels {\n    id\n    goalElements {\n      id\n      points {\n        id\n        studentTask {\n          ...CalendarStudentTask\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarStudentTask on StudentTask {\n  id\n  taskId\n  task {\n    id\n    studentTaskAdditionalAttributes {\n      ...CalendarStudentTaskAdditionalAttributes\n      __typename\n    }\n    __typename\n  }\n  lastAnswer {\n    id\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarStudentTaskAdditionalAttributes on StudentTaskAdditionalAttributes {\n  cookiesCount\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarGetModule](ctx, request)
}