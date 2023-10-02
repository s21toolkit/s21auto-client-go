package requests

import "github.com/s21toolkit/s21client/gql"

type DeadlineReminderGetClosestDeadlinePopup_Variables struct {
}


type DeadlineReminderGetClosestDeadlinePopup_Data struct {
	Student DeadlineReminderGetClosestDeadlinePopup_Data_Student `json:"student"`
}

type DeadlineReminderGetClosestDeadlinePopup_Data_Student struct {
	GetClosestDeadlinePopup interface{} `json:"getClosestDeadlinePopup"`
	Typename                string      `json:"__typename"`
}


func (ctx *RequestContext) DeadlineReminderGetClosestDeadlinePopup(variables DeadlineReminderGetClosestDeadlinePopup_Variables) (DeadlineReminderGetClosestDeadlinePopup_Data, error) {
	request := gql.NewQueryRequest[DeadlineReminderGetClosestDeadlinePopup_Variables](
		"query deadlineReminderGetClosestDeadlinePopup {\n  student {\n    getClosestDeadlinePopup {\n      deadline {\n        ...DeadlineData\n        __typename\n      }\n      deadlineGoal {\n        ...DeadlineGoalData\n        __typename\n      }\n      shiftCount\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DeadlineData on Deadline {\n  deadlineId\n  description\n  comment\n  deadlineToDaysArray\n  deadlineTs\n  createTs\n  updateTs\n  status\n  rules {\n    logicalOperatorId\n    rulesInGroup {\n      logicalOperatorId\n      value {\n        fieldId\n        subFieldKey\n        operator\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment DeadlineGoalData on DeadlineGoal {\n  goalProjects {\n    studentGoalId\n    project {\n      goalName\n      goalId\n      __typename\n    }\n    status\n    executionType\n    finalPercentage\n    finalPoint\n    pointTask\n    __typename\n  }\n  goalCourses {\n    ...GoalCourse\n    __typename\n  }\n  levels {\n    ...Level\n    __typename\n  }\n  __typename\n}\n\nfragment GoalCourse on CourseCoverInformation {\n  localCourseId\n  courseName\n  courseType\n  experienceFact\n  finalPercentage\n  displayedCourseStatus\n  __typename\n}\n\nfragment Level on ExperienceLevelRange {\n  id\n  level\n  levelCode\n  leftBorder\n  rightBorder\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[DeadlineReminderGetClosestDeadlinePopup_Data](ctx, request)
}