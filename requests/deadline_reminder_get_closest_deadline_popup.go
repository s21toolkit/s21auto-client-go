package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_DeadlineReminderGetClosestDeadlinePopup struct {
}


type Data_DeadlineReminderGetClosestDeadlinePopup struct {
	Data_Student_DeadlineReminderGetClosestDeadlinePopup Data_Student_DeadlineReminderGetClosestDeadlinePopup `json:"student"`
}

type Data_Student_DeadlineReminderGetClosestDeadlinePopup struct {
	GetClosestDeadlinePopup interface{} `json:"getClosestDeadlinePopup"`
	Typename                string      `json:"__typename"`
}


func (ctx *RequestContext) DeadlineReminderGetClosestDeadlinePopup(variables Variables_DeadlineReminderGetClosestDeadlinePopup) (Data_DeadlineReminderGetClosestDeadlinePopup, error) {
	request := gql.NewQueryRequest[Variables_DeadlineReminderGetClosestDeadlinePopup](
		"query deadlineReminderGetClosestDeadlinePopup {\n  student {\n    getClosestDeadlinePopup {\n      deadline {\n        ...DeadlineData\n        __typename\n      }\n      deadlineGoal {\n        ...DeadlineGoalData\n        __typename\n      }\n      shiftCount\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment DeadlineData on Deadline {\n  deadlineId\n  description\n  comment\n  deadlineToDaysArray\n  deadlineTs\n  createTs\n  updateTs\n  status\n  rules {\n    logicalOperatorId\n    rulesInGroup {\n      logicalOperatorId\n      value {\n        fieldId\n        subFieldKey\n        operator\n        value\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment DeadlineGoalData on DeadlineGoal {\n  goalProjects {\n    studentGoalId\n    project {\n      goalName\n      goalId\n      __typename\n    }\n    status\n    executionType\n    finalPercentage\n    finalPoint\n    pointTask\n    __typename\n  }\n  goalCourses {\n    ...GoalCourse\n    __typename\n  }\n  levels {\n    ...Level\n    __typename\n  }\n  __typename\n}\n\nfragment GoalCourse on CourseCoverInformation {\n  localCourseId\n  courseName\n  courseType\n  experienceFact\n  finalPercentage\n  displayedCourseStatus\n  __typename\n}\n\nfragment Level on ExperienceLevelRange {\n  id\n  level\n  levelCode\n  leftBorder\n  rightBorder\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_DeadlineReminderGetClosestDeadlinePopup](ctx, request)
}