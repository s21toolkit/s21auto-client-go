package requests

import "s21client/gql"

type Variables_DeadlineReminderGetClosestDeadlinePopup struct {
}


type Data_DeadlineReminderGetClosestDeadlinePopup struct {
	Student_DeadlineReminderGetClosestDeadlinePopup Student_DeadlineReminderGetClosestDeadlinePopup `json:"student"`
}

type Student_DeadlineReminderGetClosestDeadlinePopup struct {
	GetClosestDeadlinePopup interface{} `json:"getClosestDeadlinePopup"`
	Typename                string      `json:"__typename"`
}

func (ctx *RequestContext) DeadlineReminderGetClosestDeadlinePopup(variables Variables_DeadlineReminderGetClosestDeadlinePopup) (Data_DeadlineReminderGetClosestDeadlinePopup, error) {
	request := gql.NewQueryRequest[Variables_DeadlineReminderGetClosestDeadlinePopup](
		"query deadlineReminderGetClosestDeadlinePopup {   student {     getClosestDeadlinePopup {       deadline {         ...DeadlineData         __typename       }       deadlineGoal {         ...DeadlineGoalData         __typename       }       shiftCount       __typename     }     __typename   } }  fragment DeadlineData on Deadline {   deadlineId   description   comment   deadlineToDaysArray   deadlineTs   createTs   updateTs   status   rules {     logicalOperatorId     rulesInGroup {       logicalOperatorId       value {         fieldId         subFieldKey         operator         value         __typename       }       __typename     }     __typename   }   __typename }  fragment DeadlineGoalData on DeadlineGoal {   goalProjects {     studentGoalId     project {       goalName       goalId       __typename     }     status     executionType     finalPercentage     finalPoint     pointTask     __typename   }   goalCourses {     ...GoalCourse     __typename   }   levels {     ...Level     __typename   }   __typename }  fragment GoalCourse on CourseCoverInformation {   localCourseId   courseName   courseType   experienceFact   finalPercentage   displayedCourseStatus   __typename }  fragment Level on ExperienceLevelRange {   id   level   levelCode   leftBorder   rightBorder   __typename } ",
		variables,
	)

	return GqlRequest[Data_DeadlineReminderGetClosestDeadlinePopup](ctx, request)
}