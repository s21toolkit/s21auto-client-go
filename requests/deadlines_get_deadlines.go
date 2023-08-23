package requests

import "s21client/gql"

type Variables_DeadlinesGetDeadlines struct {
	Page_DeadlinesGetDeadlines             Page_DeadlinesGetDeadlines     `json:"page"`
	DeadlineStatuses []string `json:"deadlineStatuses"`
	Sorting_DeadlinesGetDeadlines          Sorting_DeadlinesGetDeadlines  `json:"sorting"`
}

type Page_DeadlinesGetDeadlines struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}

type Sorting_DeadlinesGetDeadlines struct {
	Name string `json:"name"`
	Asc  bool   `json:"asc"`
}


type Data_DeadlinesGetDeadlines struct {
	Student_DeadlinesGetDeadlines Student_DeadlinesGetDeadlines `json:"student"`
}

type Student_DeadlinesGetDeadlines struct {
	GetDeadlines []GetDeadline_DeadlinesGetDeadlines `json:"getDeadlines"`
	Typename     string        `json:"__typename"`
}

type GetDeadline_DeadlinesGetDeadlines struct {
	Deadline_DeadlinesGetDeadlines      Deadline_DeadlinesGetDeadlines      `json:"deadline"`
	ShiftRequests []interface{} `json:"shiftRequests"`
	DeadlineGoal_DeadlinesGetDeadlines  DeadlineGoal_DeadlinesGetDeadlines  `json:"deadlineGoal"`
	ShiftCount    int64         `json:"shiftCount"`
	Typename      string        `json:"__typename"`
}

type Deadline_DeadlinesGetDeadlines struct {
	DeadlineID          string  `json:"deadlineId"`
	Description         string  `json:"description"`
	Comment             string  `json:"comment"`
	DeadlineToDaysArray []int64 `json:"deadlineToDaysArray"`
	DeadlineTs          string  `json:"deadlineTs"`
	CreateTs            string  `json:"createTs"`
	UpdateTs            *string `json:"updateTs"`
	Status              string  `json:"status"`
	Rules               []Rule_DeadlinesGetDeadlines  `json:"rules"`
	Typename            string  `json:"__typename"`
}

type Rule_DeadlinesGetDeadlines struct {
	LogicalOperatorID interface{}    `json:"logicalOperatorId"`
	RulesInGroup_DeadlinesGetDeadlines      []RulesInGroup_DeadlinesGetDeadlines `json:"rulesInGroup"`
	Typename          string         `json:"__typename"`
}

type RulesInGroup_DeadlinesGetDeadlines struct {
	LogicalOperatorID interface{} `json:"logicalOperatorId"`
	Value_DeadlinesGetDeadlines             Value_DeadlinesGetDeadlines       `json:"value"`
	Typename          string      `json:"__typename"`
}

type Value_DeadlinesGetDeadlines struct {
	FieldID     string      `json:"fieldId"`
	SubFieldKey interface{} `json:"subFieldKey"`
	Operator    string      `json:"operator"`
	Value_DeadlinesGetDeadlines       []string    `json:"value"`
	Typename    string      `json:"__typename"`
}

type DeadlineGoal_DeadlinesGetDeadlines struct {
	GoalProjects []interface{} `json:"goalProjects"`
	GoalCourses  []interface{} `json:"goalCourses"`
	Levels       []Level_DeadlinesGetDeadlines       `json:"levels"`
	Typename     string        `json:"__typename"`
}

type Level_DeadlinesGetDeadlines struct {
	ID          string `json:"id"`
	Level_DeadlinesGetDeadlines       int64  `json:"level"`
	LevelCode   int64  `json:"levelCode"`
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}

func (ctx *RequestContext) DeadlinesGetDeadlines(variables Variables_DeadlinesGetDeadlines) (Data_DeadlinesGetDeadlines, error) {
	request := gql.NewQueryRequest[Variables_DeadlinesGetDeadlines](
		"query deadlinesGetDeadlines($deadlineStatuses: [DeadlineStatus!]!, $page: PagingInput!, $deadlinesFrom: DateTime, $deadlinesTo: DateTime, $sorting: [SortingField]) {   student {     getDeadlines(       deadlineStatuses: $deadlineStatuses       page: $page       deadlineFrom: $deadlinesFrom       deadlineTo: $deadlinesTo       sorting: $sorting     ) {       deadline {         ...DeadlineData         __typename       }       shiftRequests {         deadlineShiftRequestId         status         daysToShift         createTs         __typename       }       deadlineGoal {         ...DeadlineGoalData         __typename       }       shiftCount       __typename     }     __typename   } }  fragment DeadlineData on Deadline {   deadlineId   description   comment   deadlineToDaysArray   deadlineTs   createTs   updateTs   status   rules {     logicalOperatorId     rulesInGroup {       logicalOperatorId       value {         fieldId         subFieldKey         operator         value         __typename       }       __typename     }     __typename   }   __typename }  fragment DeadlineGoalData on DeadlineGoal {   goalProjects {     studentGoalId     project {       goalName       goalId       __typename     }     status     executionType     finalPercentage     finalPoint     pointTask     __typename   }   goalCourses {     ...GoalCourse     __typename   }   levels {     ...Level     __typename   }   __typename }  fragment GoalCourse on CourseCoverInformation {   localCourseId   courseName   courseType   experienceFact   finalPercentage   displayedCourseStatus   __typename }  fragment Level on ExperienceLevelRange {   id   level   levelCode   leftBorder   rightBorder   __typename } ",
		variables,
	)

	return GqlRequest[Data_DeadlinesGetDeadlines](ctx, request)
}