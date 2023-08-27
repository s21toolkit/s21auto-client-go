package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_DeadlinesGetDeadlines struct {
	Variables_Page_DeadlinesGetDeadlines             Variables_Page_DeadlinesGetDeadlines     `json:"page"`
	DeadlineStatuses []string `json:"deadlineStatuses"`
	Variables_Sorting_DeadlinesGetDeadlines          Variables_Sorting_DeadlinesGetDeadlines  `json:"sorting"`
}

type Variables_Page_DeadlinesGetDeadlines struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type Variables_Sorting_DeadlinesGetDeadlines struct {
	Name string `json:"name"`
	Asc  bool   `json:"asc"`
}


type Data_DeadlinesGetDeadlines struct {
	Data_Student_DeadlinesGetDeadlines Data_Student_DeadlinesGetDeadlines `json:"student"`
}

type Data_Student_DeadlinesGetDeadlines struct {
	GetDeadlines []Data_GetDeadline_DeadlinesGetDeadlines `json:"getDeadlines"`
	Typename     string        `json:"__typename"`
}

type Data_GetDeadline_DeadlinesGetDeadlines struct {
	Data_Deadline_DeadlinesGetDeadlines      Data_Deadline_DeadlinesGetDeadlines      `json:"deadline"`
	ShiftRequests []interface{} `json:"shiftRequests"`
	Data_DeadlineGoal_DeadlinesGetDeadlines  Data_DeadlineGoal_DeadlinesGetDeadlines  `json:"deadlineGoal"`
	ShiftCount    int64         `json:"shiftCount"`
	Typename      string        `json:"__typename"`
}

type Data_Deadline_DeadlinesGetDeadlines struct {
	DeadlineID          string  `json:"deadlineId"`
	Description         string  `json:"description"`
	Comment             string  `json:"comment"`
	DeadlineToDaysArray []int64 `json:"deadlineToDaysArray"`
	DeadlineTs          string  `json:"deadlineTs"`
	CreateTs            string  `json:"createTs"`
	UpdateTs            *string `json:"updateTs"`
	Status              string  `json:"status"`
	Rules               []Data_Rule_DeadlinesGetDeadlines  `json:"rules"`
	Typename            string  `json:"__typename"`
}

type Data_Rule_DeadlinesGetDeadlines struct {
	LogicalOperatorID interface{}    `json:"logicalOperatorId"`
	Data_RulesInGroup_DeadlinesGetDeadlines      []Data_RulesInGroup_DeadlinesGetDeadlines `json:"rulesInGroup"`
	Typename          string         `json:"__typename"`
}

type Data_RulesInGroup_DeadlinesGetDeadlines struct {
	LogicalOperatorID interface{} `json:"logicalOperatorId"`
	Data_Value_DeadlinesGetDeadlines             Data_Value_DeadlinesGetDeadlines       `json:"value"`
	Typename          string      `json:"__typename"`
}

type Data_Value_DeadlinesGetDeadlines struct {
	FieldID     string      `json:"fieldId"`
	SubFieldKey interface{} `json:"subFieldKey"`
	Operator    string      `json:"operator"`
	Data_Value_DeadlinesGetDeadlines       []string    `json:"value"`
	Typename    string      `json:"__typename"`
}

type Data_DeadlineGoal_DeadlinesGetDeadlines struct {
	GoalProjects []interface{} `json:"goalProjects"`
	GoalCourses  []interface{} `json:"goalCourses"`
	Levels       []Data_Level_DeadlinesGetDeadlines       `json:"levels"`
	Typename     string        `json:"__typename"`
}

type Data_Level_DeadlinesGetDeadlines struct {
	ID          string `json:"id"`
	Data_Level_DeadlinesGetDeadlines       int64  `json:"level"`
	LevelCode   int64  `json:"levelCode"`
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}


func (ctx *RequestContext) DeadlinesGetDeadlines(variables Variables_DeadlinesGetDeadlines) (Data_DeadlinesGetDeadlines, error) {
	request := gql.NewQueryRequest[Variables_DeadlinesGetDeadlines](
		"query deadlinesGetDeadlines($deadlineStatuses: [DeadlineStatus!]!, $page: PagingInput!, $deadlinesFrom: DateTime, $deadlinesTo: DateTime, $sorting: [SortingField]) {\n  student {\n    getDeadlines(\n      deadlineStatuses: $deadlineStatuses\n      page: $page\n      deadlineFrom: $deadlinesFrom\n      deadlineTo: $deadlinesTo\n      sorting: $sorting\n    ) {\n      deadline {\n        ...DeadlineData\n        __typename\n      }\n      shiftRequests {\n        deadlineShiftRequestId\n        status\n        daysToShift\n        createTs\n        __typename\n      }\n      deadlineGoal {\n        ...DeadlineGoalData\n        __typename\n      }\n      shiftCount\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DeadlineData on Deadline {\n  deadlineId\n  description\n  comment\n  deadlineToDaysArray\n  deadlineTs\n  createTs\n  updateTs\n  status\n  rules {\n    logicalOperatorId\n    rulesInGroup {\n      logicalOperatorId\n      value {\n        fieldId\n        subFieldKey\n        operator\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment DeadlineGoalData on DeadlineGoal {\n  goalProjects {\n    studentGoalId\n    project {\n      goalName\n      goalId\n      __typename\n    }\n    status\n    executionType\n    finalPercentage\n    finalPoint\n    pointTask\n    __typename\n  }\n  goalCourses {\n    ...GoalCourse\n    __typename\n  }\n  levels {\n    ...Level\n    __typename\n  }\n  __typename\n}\n\nfragment GoalCourse on CourseCoverInformation {\n  localCourseId\n  courseName\n  courseType\n  experienceFact\n  finalPercentage\n  displayedCourseStatus\n  __typename\n}\n\nfragment Level on ExperienceLevelRange {\n  id\n  level\n  levelCode\n  leftBorder\n  rightBorder\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_DeadlinesGetDeadlines](ctx, request)
}