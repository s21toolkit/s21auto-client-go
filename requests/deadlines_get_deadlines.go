package requests

import "github.com/s21toolkit/s21client/gql"

type DeadlinesGetDeadlines_Variables struct {
	Page             DeadlinesGetDeadlines_Variables_Page     `json:"page"`
	DeadlineStatuses []string `json:"deadlineStatuses"`
	Sorting          DeadlinesGetDeadlines_Variables_Sorting  `json:"sorting"`
}

type DeadlinesGetDeadlines_Variables_Page struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type DeadlinesGetDeadlines_Variables_Sorting struct {
	Name string `json:"name"`
	Asc  bool   `json:"asc"`
}


type DeadlinesGetDeadlines_Data struct {
	Student DeadlinesGetDeadlines_Data_Student `json:"student"`
}

type DeadlinesGetDeadlines_Data_Student struct {
	GetDeadlines []DeadlinesGetDeadlines_Data_GetDeadline `json:"getDeadlines"`
	Typename     string        `json:"__typename"`
}

type DeadlinesGetDeadlines_Data_GetDeadline struct {
	Deadline      DeadlinesGetDeadlines_Data_Deadline      `json:"deadline"`
	ShiftRequests []interface{} `json:"shiftRequests"`
	DeadlineGoal  DeadlinesGetDeadlines_Data_DeadlineGoal  `json:"deadlineGoal"`
	ShiftCount    int64         `json:"shiftCount"`
	Typename      string        `json:"__typename"`
}

type DeadlinesGetDeadlines_Data_Deadline struct {
	DeadlineID          string  `json:"deadlineId"`
	Description         string  `json:"description"`
	Comment             string  `json:"comment"`
	DeadlineToDaysArray []int64 `json:"deadlineToDaysArray"`
	DeadlineTs          string  `json:"deadlineTs"`
	CreateTs            string  `json:"createTs"`
	UpdateTs            *string `json:"updateTs"`
	Status              string  `json:"status"`
	Rules               []DeadlinesGetDeadlines_Data_Rule  `json:"rules"`
	Typename            string  `json:"__typename"`
}

type DeadlinesGetDeadlines_Data_Rule struct {
	LogicalOperatorID interface{}    `json:"logicalOperatorId"`
	RulesInGroup      []DeadlinesGetDeadlines_Data_RulesInGroup `json:"rulesInGroup"`
	Typename          string         `json:"__typename"`
}

type DeadlinesGetDeadlines_Data_RulesInGroup struct {
	LogicalOperatorID interface{} `json:"logicalOperatorId"`
	Value             DeadlinesGetDeadlines_Data_Value       `json:"value"`
	Typename          string      `json:"__typename"`
}

type DeadlinesGetDeadlines_Data_Value struct {
	FieldID     string      `json:"fieldId"`
	SubFieldKey interface{} `json:"subFieldKey"`
	Operator    string      `json:"operator"`
	Value       []string    `json:"value"`
	Typename    string      `json:"__typename"`
}

type DeadlinesGetDeadlines_Data_DeadlineGoal struct {
	GoalProjects []interface{} `json:"goalProjects"`
	GoalCourses  []interface{} `json:"goalCourses"`
	Levels       []DeadlinesGetDeadlines_Data_Level       `json:"levels"`
	Typename     string        `json:"__typename"`
}

type DeadlinesGetDeadlines_Data_Level struct {
	ID          string `json:"id"`
	Level       int64  `json:"level"`
	LevelCode   int64  `json:"levelCode"`
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}


func (ctx *RequestContext) DeadlinesGetDeadlines(variables DeadlinesGetDeadlines_Variables) (DeadlinesGetDeadlines_Data, error) {
	request := gql.NewQueryRequest[DeadlinesGetDeadlines_Variables](
		"query deadlinesGetDeadlines($deadlineStatuses: [DeadlineStatus!]!, $page: PagingInput!, $deadlinesFrom: DateTime, $deadlinesTo: DateTime, $sorting: [SortingField]) {\n  student {\n    getDeadlines(\n      deadlineStatuses: $deadlineStatuses\n      page: $page\n      deadlineFrom: $deadlinesFrom\n      deadlineTo: $deadlinesTo\n      sorting: $sorting\n    ) {\n      deadline {\n        ...DeadlineData\n        __typename\n      }\n      shiftRequests {\n        deadlineShiftRequestId\n        status\n        daysToShift\n        createTs\n        __typename\n      }\n      deadlineGoal {\n        ...DeadlineGoalData\n        __typename\n      }\n      shiftCount\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DeadlineData on Deadline {\n  deadlineId\n  description\n  comment\n  deadlineToDaysArray\n  deadlineTs\n  createTs\n  updateTs\n  status\n  rules {\n    logicalOperatorId\n    rulesInGroup {\n      logicalOperatorId\n      value {\n        fieldId\n        subFieldKey\n        operator\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment DeadlineGoalData on DeadlineGoal {\n  goalProjects {\n    studentGoalId\n    project {\n      goalName\n      goalId\n      __typename\n    }\n    status\n    executionType\n    finalPercentage\n    finalPoint\n    pointTask\n    __typename\n  }\n  goalCourses {\n    ...GoalCourse\n    __typename\n  }\n  levels {\n    ...Level\n    __typename\n  }\n  __typename\n}\n\nfragment GoalCourse on CourseCoverInformation {\n  localCourseId\n  courseName\n  courseType\n  experienceFact\n  finalPercentage\n  displayedCourseStatus\n  __typename\n}\n\nfragment Level on ExperienceLevelRange {\n  id\n  level\n  levelCode\n  leftBorder\n  rightBorder\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[DeadlinesGetDeadlines_Data](ctx, request)
}