package requests

import "s21client/gql"

type Request_Variables_CalendarGetEvents struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Response_Data_CalendarGetEvents struct {
	Response_Student_CalendarGetEvents Response_Student_CalendarGetEvents `json:"student"`
}

type Response_Student_CalendarGetEvents struct {
	GetMyCalendarEvents []Response_GetMyCalendarEvent_CalendarGetEvents `json:"getMyCalendarEvents"`
	Typename            string               `json:"__typename"`
}

type Response_GetMyCalendarEvent_CalendarGetEvents struct {
	ID                string        `json:"id"`
	Start             string        `json:"start"`
	End               string        `json:"end"`
	Description       string        `json:"description"`
	EventType         string        `json:"eventType"`
	EventCode         string        `json:"eventCode"`
	EventSlots        []Response_EventSlot_CalendarGetEvents   `json:"eventSlots"`
	Bookings          []interface{} `json:"bookings"`
	Exam              interface{}   `json:"exam"`
	StudentCodeReview interface{}   `json:"studentCodeReview"`
	Activity          interface{}   `json:"activity"`
	Goals             []interface{} `json:"goals"`
	Penalty           interface{}   `json:"penalty"`
	Typename          string        `json:"__typename"`
}

type Response_EventSlot_CalendarGetEvents struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetEvents(variables Request_Variables_CalendarGetEvents) (Response_Data_CalendarGetEvents, error) {
	request := gql.NewQueryRequest[Request_Variables_CalendarGetEvents](
		"query calendarGetEvents($from: DateTime!, $to: DateTime!) {   student {     getMyCalendarEvents(from: $from, to: $to) {       ...CalendarEvent       __typename     }     __typename   } }  fragment CalendarEvent on CalendarEvent {   id   start   end   description   eventType   eventCode   eventSlots {     id     type     start     end     __typename   }   bookings {     ...CalendarReviewBooking     __typename   }   exam {     ...CalendarEventExam     __typename   }   studentCodeReview {     studentGoalId     __typename   }   activity {     ...CalendarEventActivity     studentFeedback {       id       rating       comment       __typename     }     status     activityType     isMandatory     isWaitListActive     isVisible     comments {       type       createTs       comment       __typename     }     organizers {       id       login       __typename     }     __typename   }   goals {     goalId     goalName     __typename   }   penalty {     ...Penalty     __typename   }   __typename }  fragment CalendarReviewBooking on CalendarBooking {   id   answerId   eventSlotId   task {     id     goalId     goalName     studentTaskAdditionalAttributes {       cookiesCount       __typename     }     assignmentType     __typename   }   eventSlot {     id     start     end     event {       eventUserRole       __typename     }     __typename   }   verifierUser {     ...CalendarReviewUser     __typename   }   verifiableStudent {     id     user {       ...CalendarReviewUser       __typename     }     __typename   }   bookingStatus   team {     ...ProjectTeamMembers     __typename   }   isOnline   vcLinkUrl   __typename }  fragment CalendarReviewUser on User {   id   login   __typename }  fragment ProjectTeamMembers on ProjectTeamMembers {   id   teamLead {     ...ProjectTeamMember     __typename   }   members {     ...ProjectTeamMember     __typename   }   invitedUsers {     ...ProjectTeamMember     __typename   }   teamName   teamStatus   minTeamMemberCount   maxTeamMemberCount   __typename }  fragment ProjectTeamMember on User {   id   avatarUrl   login   userExperience {     level {       id       range {         levelCode         __typename       }       __typename     }     cookiesCount     codeReviewPoints     __typename   }   __typename }  fragment CalendarEventExam on Exam {   examId   eventId   beginDate   endDate   name   location   currentStudentsCount   maxStudentCount   updateDate   goalId   goalName   isWaitListActive   isInWaitList   stopRegisterDate   __typename }  fragment CalendarEventActivity on ActivityEvent {   activityEventId   eventId   name   beginDate   endDate   isRegistered   description   currentStudentsCount   maxStudentCount   location   updateDate   isWaitListActive   isInWaitList   stopRegisterDate   __typename }  fragment Penalty on Penalty {   comment   id   duration   status   startTime   createTime   penaltySlot {     currentStudentsCount     description     duration     startTime     id     endTime     __typename   }   reasonId   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_CalendarGetEvents](ctx, request)
}