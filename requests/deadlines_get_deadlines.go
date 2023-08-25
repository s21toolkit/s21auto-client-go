package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_DeadlinesGetDeadlines struct {
	Request_Page_DeadlinesGetDeadlines             Request_Page_DeadlinesGetDeadlines     `json:"page"`
	DeadlineStatuses []string `json:"deadlineStatuses"`
	Request_Sorting_DeadlinesGetDeadlines          Request_Sorting_DeadlinesGetDeadlines  `json:"sorting"`
}

type Request_Page_DeadlinesGetDeadlines struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type Request_Sorting_DeadlinesGetDeadlines struct {
	Name string `json:"name"`
	Asc  bool   `json:"asc"`
}


type Response_Data_DeadlinesGetDeadlines struct {
	Response_Student_DeadlinesGetDeadlines Response_Student_DeadlinesGetDeadlines `json:"student"`
}

type Response_Student_DeadlinesGetDeadlines struct {
	GetDeadlines []Response_GetDeadline_DeadlinesGetDeadlines `json:"getDeadlines"`
	Typename     string        `json:"__typename"`
}

type Response_GetDeadline_DeadlinesGetDeadlines struct {
	Response_Deadline_DeadlinesGetDeadlines      Response_Deadline_DeadlinesGetDeadlines      `json:"deadline"`
	ShiftRequests []interface{} `json:"shiftRequests"`
	Response_DeadlineGoal_DeadlinesGetDeadlines  Response_DeadlineGoal_DeadlinesGetDeadlines  `json:"deadlineGoal"`
	ShiftCount    int64         `json:"shiftCount"`
	Typename      string        `json:"__typename"`
}

type Response_Deadline_DeadlinesGetDeadlines struct {
	DeadlineID          string  `json:"deadlineId"`
	Description         string  `json:"description"`
	Comment             string  `json:"comment"`
	DeadlineToDaysArray []int64 `json:"deadlineToDaysArray"`
	DeadlineTs          string  `json:"deadlineTs"`
	CreateTs            string  `json:"createTs"`
	UpdateTs            *string `json:"updateTs"`
	Status              string  `json:"status"`
	Rules               []Response_Rule_DeadlinesGetDeadlines  `json:"rules"`
	Typename            string  `json:"__typename"`
}

type Response_Rule_DeadlinesGetDeadlines struct {
	LogicalOperatorID interface{}    `json:"logicalOperatorId"`
	Response_RulesInGroup_DeadlinesGetDeadlines      []Response_RulesInGroup_DeadlinesGetDeadlines `json:"rulesInGroup"`
	Typename          string         `json:"__typename"`
}

type Response_RulesInGroup_DeadlinesGetDeadlines struct {
	LogicalOperatorID interface{} `json:"logicalOperatorId"`
	Response_Value_DeadlinesGetDeadlines             Response_Value_DeadlinesGetDeadlines       `json:"value"`
	Typename          string      `json:"__typename"`
}

type Response_Value_DeadlinesGetDeadlines struct {
	FieldID     string      `json:"fieldId"`
	SubFieldKey interface{} `json:"subFieldKey"`
	Operator    string      `json:"operator"`
	Response_Value_DeadlinesGetDeadlines       []string    `json:"value"`
	Typename    string      `json:"__typename"`
}

type Response_DeadlineGoal_DeadlinesGetDeadlines struct {
	GoalProjects []interface{} `json:"goalProjects"`
	GoalCourses  []interface{} `json:"goalCourses"`
	Levels       []Response_Level_DeadlinesGetDeadlines       `json:"levels"`
	Typename     string        `json:"__typename"`
}

type Response_Level_DeadlinesGetDeadlines struct {
	ID          string `json:"id"`
	Response_Level_DeadlinesGetDeadlines       int64  `json:"level"`
	LevelCode   int64  `json:"levelCode"`
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}

func (ctx *RequestContext) DeadlinesGetDeadlines(variables Request_Variables_DeadlinesGetDeadlines) (Response_Data_DeadlinesGetDeadlines, error) {
	request := gql.NewQueryRequest[Request_Variables_DeadlinesGetDeadlines](
		"query deadlinesGetDeadlines($deadlineStatuses: [DeadlineStatus!]!, $page: PagingInput!, $deadlinesFrom: DateTime, $deadlinesTo: DateTime, $sorting: [SortingField]) {\n  student {\n    getDeadlines(\n      deadlineStatuses: $deadlineStatuses\n      page: $page\n      deadlineFrom: $deadlinesFrom\n      deadlineTo: $deadlinesTo\n      sorting: $sorting\n    ) {\n      deadline {\n        ...DeadlineData\n        __typename\n      }\n      shiftRequests {\n        deadlineShiftRequestId\n        status\n        daysToShift\n        createTs\n        __typename\n      }\n      deadlineGoal {\n        ...DeadlineGoalData\n        __typename\n      }\n      shiftCount\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DeadlineData on Deadline {\n  deadlineId\n  description\n  comment\n  deadlineToDaysArray\n  deadlineTs\n  createTs\n  updateTs\n  status\n  rules {\n    logicalOperatorId\n    rulesInGroup {\n      logicalOperatorId\n      value {\n        fieldId\n        subFieldKey\n        operator\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment DeadlineGoalData on DeadlineGoal {\n  goalProjects {\n    studentGoalId\n    project {\n      goalName\n      goalId\n      __typename\n    }\n    status\n    executionType\n    finalPercentage\n    finalPoint\n    pointTask\n    __typename\n  }\n  goalCourses {\n    ...GoalCourse\n    __typename\n  }\n  levels {\n    ...Level\n    __typename\n  }\n  __typename\n}\n\nfragment GoalCourse on CourseCoverInformation {\n  localCourseId\n  courseName\n  courseType\n  experienceFact\n  finalPercentage\n  displayedCourseStatus\n  __typename\n}\n\nfragment Level on ExperienceLevelRange {\n  id\n  level\n  levelCode\n  leftBorder\n  rightBorder\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_DeadlinesGetDeadlines](ctx, request)
}