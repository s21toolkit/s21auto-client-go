package requests

import "s21client/gql"

type Variables_CalendarAddEvent struct {
	Start string `json:"start"`
	End   string `json:"end"`
}


type Data_CalendarAddEvent struct {
	Student_CalendarAddEvent Student_CalendarAddEvent `json:"student"`
}

type Student_CalendarAddEvent struct {
	AddEventToTimetable_CalendarAddEvent []AddEventToTimetable_CalendarAddEvent `json:"addEventToTimetable"`
	Typename            string                `json:"__typename"`
}

type AddEventToTimetable_CalendarAddEvent struct {
	ID                string        `json:"id"`
	Start             string        `json:"start"`
	End               string        `json:"end"`
	Description       string        `json:"description"`
	EventType         string        `json:"eventType"`
	EventCode         interface{}   `json:"eventCode"`
	EventSlots        []EventSlot_CalendarAddEvent   `json:"eventSlots"`
	Bookings          []interface{} `json:"bookings"`
	Exam              interface{}   `json:"exam"`
	StudentCodeReview interface{}   `json:"studentCodeReview"`
	Activity          interface{}   `json:"activity"`
	Goals             []interface{} `json:"goals"`
	Penalty           interface{}   `json:"penalty"`
	Typename          string        `json:"__typename"`
}

type EventSlot_CalendarAddEvent struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) CalendarAddEvent(variables Variables_CalendarAddEvent) (Data_CalendarAddEvent, error) {
	request := gql.NewQueryRequest[Variables_CalendarAddEvent](
		"mutation calendarAddEvent($start: DateTime!, $end: DateTime!) {   student {     addEventToTimetable(start: $start, end: $end) {       ...CalendarEvent       __typename     }     __typename   } }  fragment CalendarEvent on CalendarEvent {   id   start   end   description   eventType   eventCode   eventSlots {     id     type     start     end     __typename   }   bookings {     ...CalendarReviewBooking     __typename   }   exam {     ...CalendarEventExam     __typename   }   studentCodeReview {     studentGoalId     __typename   }   activity {     ...CalendarEventActivity     studentFeedback {       id       rating       comment       __typename     }     status     activityType     isMandatory     isWaitListActive     isVisible     comments {       type       createTs       comment       __typename     }     organizers {       id       login       __typename     }     __typename   }   goals {     goalId     goalName     __typename   }   penalty {     ...Penalty     __typename   }   __typename }  fragment CalendarReviewBooking on CalendarBooking {   id   answerId   eventSlotId   task {     id     goalId     goalName     studentTaskAdditionalAttributes {       cookiesCount       __typename     }     assignmentType     __typename   }   eventSlot {     id     start     end     event {       eventUserRole       __typename     }     __typename   }   verifierUser {     ...CalendarReviewUser     __typename   }   verifiableStudent {     id     user {       ...CalendarReviewUser       __typename     }     __typename   }   bookingStatus   team {     ...ProjectTeamMembers     __typename   }   isOnline   vcLinkUrl   __typename }  fragment CalendarReviewUser on User {   id   login   __typename }  fragment ProjectTeamMembers on ProjectTeamMembers {   id   teamLead {     ...ProjectTeamMember     __typename   }   members {     ...ProjectTeamMember     __typename   }   invitedUsers {     ...ProjectTeamMember     __typename   }   teamName   teamStatus   minTeamMemberCount   maxTeamMemberCount   __typename }  fragment ProjectTeamMember on User {   id   avatarUrl   login   userExperience {     level {       id       range {         levelCode         __typename       }       __typename     }     cookiesCount     codeReviewPoints     __typename   }   __typename }  fragment CalendarEventExam on Exam {   examId   eventId   beginDate   endDate   name   location   currentStudentsCount   maxStudentCount   updateDate   goalId   goalName   isWaitListActive   isInWaitList   stopRegisterDate   __typename }  fragment CalendarEventActivity on ActivityEvent {   activityEventId   eventId   name   beginDate   endDate   isRegistered   description   currentStudentsCount   maxStudentCount   location   updateDate   isWaitListActive   isInWaitList   stopRegisterDate   __typename }  fragment Penalty on Penalty {   comment   id   duration   status   startTime   createTime   penaltySlot {     currentStudentsCount     description     duration     startTime     id     endTime     __typename   }   reasonId   __typename } ",
		variables,
	)

	return GqlRequest[Data_CalendarAddEvent](ctx, request)
}