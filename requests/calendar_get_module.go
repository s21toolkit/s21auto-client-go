package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_CalendarGetModule struct {
	ModuleID string `json:"moduleId"`
}


type Response_Data_CalendarGetModule struct {
	Response_Student_CalendarGetModule Response_Student_CalendarGetModule `json:"student"`
}

type Response_Student_CalendarGetModule struct {
	Response_GetModuleByID_CalendarGetModule Response_GetModuleByID_CalendarGetModule `json:"getModuleById"`
	Typename      string        `json:"__typename"`
}

type Response_GetModuleByID_CalendarGetModule struct {
	ID                string      `json:"id"`
	ModuleTitle       string      `json:"moduleTitle"`
	SubjectTitle      string      `json:"subjectTitle"`
	GoalExecutionType string      `json:"goalExecutionType"`
	Response_Trajectory_CalendarGetModule        Response_Trajectory_CalendarGetModule  `json:"trajectory"`
	Response_CurrentTask_CalendarGetModule       Response_CurrentTask_CalendarGetModule `json:"currentTask"`
	Typename          string      `json:"__typename"`
}

type Response_CurrentTask_CalendarGetModule struct {
	ID       string          `json:"id"`
	Task     Response_CurrentTaskTask_CalendarGetModule `json:"task"`
	Typename string          `json:"__typename"`
}

type Response_CurrentTaskTask_CalendarGetModule struct {
	ID             string `json:"id"`
	AssignmentType string `json:"assignmentType"`
	Typename       string `json:"__typename"`
}

type Response_Trajectory_CalendarGetModule struct {
	ID       string  `json:"id"`
	Levels   []Response_Level_CalendarGetModule `json:"levels"`
	Typename string  `json:"__typename"`
}

type Response_Level_CalendarGetModule struct {
	ID           string        `json:"id"`
	GoalElements []Response_GoalElement_CalendarGetModule `json:"goalElements"`
	Typename     string        `json:"__typename"`
}

type Response_GoalElement_CalendarGetModule struct {
	ID       string  `json:"id"`
	Points   []Response_Point_CalendarGetModule `json:"points"`
	Typename string  `json:"__typename"`
}

type Response_Point_CalendarGetModule struct {
	ID          string      `json:"id"`
	Response_StudentTask_CalendarGetModule Response_StudentTask_CalendarGetModule `json:"studentTask"`
	Typename    string      `json:"__typename"`
}

type Response_StudentTask_CalendarGetModule struct {
	ID         string          `json:"id"`
	TaskID     string          `json:"taskId"`
	Task       Response_StudentTaskTask_CalendarGetModule `json:"task"`
	Response_LastAnswer_CalendarGetModule Response_LastAnswer_CalendarGetModule      `json:"lastAnswer"`
	Typename   string          `json:"__typename"`
}

type Response_LastAnswer_CalendarGetModule struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type Response_StudentTaskTask_CalendarGetModule struct {
	ID                              string                          `json:"id"`
	Response_StudentTaskAdditionalAttributes_CalendarGetModule Response_StudentTaskAdditionalAttributes_CalendarGetModule `json:"studentTaskAdditionalAttributes"`
	Typename                        string                          `json:"__typename"`
}

type Response_StudentTaskAdditionalAttributes_CalendarGetModule struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetModule(variables Request_Variables_CalendarGetModule) (Response_Data_CalendarGetModule, error) {
	request := gql.NewQueryRequest[Request_Variables_CalendarGetModule](
		"query calendarGetModule($moduleId: ID!) {\n  student {\n    getModuleById(goalId: $moduleId) {\n      id\n      moduleTitle\n      subjectTitle\n      goalExecutionType\n      trajectory {\n        ...CalendarModuleTrajectory\n        __typename\n      }\n      currentTask {\n        id\n        task {\n          id\n          assignmentType\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarModuleTrajectory on PersonalTrajectory {\n  id\n  levels {\n    id\n    goalElements {\n      id\n      points {\n        id\n        studentTask {\n          ...CalendarStudentTask\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarStudentTask on StudentTask {\n  id\n  taskId\n  task {\n    id\n    studentTaskAdditionalAttributes {\n      ...CalendarStudentTaskAdditionalAttributes\n      __typename\n    }\n    __typename\n  }\n  lastAnswer {\n    id\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarStudentTaskAdditionalAttributes on StudentTaskAdditionalAttributes {\n  cookiesCount\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_CalendarGetModule](ctx, request)
}