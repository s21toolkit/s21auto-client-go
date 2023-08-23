package requests

import "s21client/gql"

type Variables_CalendarGetEvents struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Data_CalendarGetEvents struct {
	Student_CalendarGetEvents Student_CalendarGetEvents `json:"student"`
}

type Student_CalendarGetEvents struct {
	GetMyCalendarEvents []GetMyCalendarEvent_CalendarGetEvents `json:"getMyCalendarEvents"`
	Typename            string               `json:"__typename"`
}

type GetMyCalendarEvent_CalendarGetEvents struct {
	ID                string        `json:"id"`
	Start             string        `json:"start"`
	End               string        `json:"end"`
	Description       string        `json:"description"`
	EventType         string        `json:"eventType"`
	EventCode         string        `json:"eventCode"`
	EventSlots        []EventSlot_CalendarGetEvents   `json:"eventSlots"`
	Bookings          []interface{} `json:"bookings"`
	Exam              interface{}   `json:"exam"`
	StudentCodeReview interface{}   `json:"studentCodeReview"`
	Activity_CalendarGetEvents          *Activity_CalendarGetEvents     `json:"activity"`
	Goals             []interface{} `json:"goals"`
	Penalty           interface{}   `json:"penalty"`
	Typename          string        `json:"__typename"`
}

type Activity_CalendarGetEvents struct {
	ActivityEventID      string          `json:"activityEventId"`
	EventID              string          `json:"eventId"`
	Name                 string          `json:"name"`
	BeginDate            string          `json:"beginDate"`
	EndDate              string          `json:"endDate"`
	IsRegistered         bool            `json:"isRegistered"`
	Description          string          `json:"description"`
	CurrentStudentsCount int64           `json:"currentStudentsCount"`
	MaxStudentCount      int64           `json:"maxStudentCount"`
	Location             string          `json:"location"`
	UpdateDate           string          `json:"updateDate"`
	IsWaitListActive     bool            `json:"isWaitListActive"`
	IsInWaitList         bool            `json:"isInWaitList"`
	StopRegisterDate     string          `json:"stopRegisterDate"`
	Typename             string          `json:"__typename"`
	StudentFeedback_CalendarGetEvents      StudentFeedback_CalendarGetEvents `json:"studentFeedback"`
	Status               string          `json:"status"`
	ActivityType         string          `json:"activityType"`
	IsMandatory          bool            `json:"isMandatory"`
	IsVisible            bool            `json:"isVisible"`
	Comments             []interface{}   `json:"comments"`
	Organizers           []Organizer_CalendarGetEvents     `json:"organizers"`
}

type Organizer_CalendarGetEvents struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}

type StudentFeedback_CalendarGetEvents struct {
	ID       string `json:"id"`
	Rating   int64  `json:"rating"`
	Comment  string `json:"comment"`
	Typename string `json:"__typename"`
}

type EventSlot_CalendarGetEvents struct {
	ID       string `json:"id"`
	Type     string `json:"type"`
	Start    string `json:"start"`
	End      string `json:"end"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetEvents(variables Variables_CalendarGetEvents) (Data_CalendarGetEvents, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetEvents](
		"query calendarGetEvents($from: DateTime!, $to: DateTime!) {   student {     getMyCalendarEvents(from: $from, to: $to) {       ...CalendarEvent       __typename     }     __typename   } }  fragment CalendarEvent on CalendarEvent {   id   start   end   description   eventType   eventCode   eventSlots {     id     type     start     end     __typename   }   bookings {     ...CalendarReviewBooking     __typename   }   exam {     ...CalendarEventExam     __typename   }   studentCodeReview {     studentGoalId     __typename   }   activity {     ...CalendarEventActivity     studentFeedback {       id       rating       comment       __typename     }     status     activityType     isMandatory     isWaitListActive     isVisible     comments {       type       createTs       comment       __typename     }     organizers {       id       login       __typename     }     __typename   }   goals {     goalId     goalName     __typename   }   penalty {     ...Penalty     __typename   }   __typename }  fragment CalendarReviewBooking on CalendarBooking {   id   answerId   eventSlotId   task {     id     goalId     goalName     studentTaskAdditionalAttributes {       cookiesCount       __typename     }     assignmentType     __typename   }   eventSlot {     id     start     end     event {       eventUserRole       __typename     }     __typename   }   verifierUser {     ...CalendarReviewUser     __typename   }   verifiableStudent {     id     user {       ...CalendarReviewUser       __typename     }     __typename   }   bookingStatus   team {     ...ProjectTeamMembers     __typename   }   isOnline   vcLinkUrl   __typename }  fragment CalendarReviewUser on User {   id   login   __typename }  fragment ProjectTeamMembers on ProjectTeamMembers {   id   teamLead {     ...ProjectTeamMember     __typename   }   members {     ...ProjectTeamMember     __typename   }   invitedUsers {     ...ProjectTeamMember     __typename   }   teamName   teamStatus   minTeamMemberCount   maxTeamMemberCount   __typename }  fragment ProjectTeamMember on User {   id   avatarUrl   login   userExperience {     level {       id       range {         levelCode         __typename       }       __typename     }     cookiesCount     codeReviewPoints     __typename   }   __typename }  fragment CalendarEventExam on Exam {   examId   eventId   beginDate   endDate   name   location   currentStudentsCount   maxStudentCount   updateDate   goalId   goalName   isWaitListActive   isInWaitList   stopRegisterDate   __typename }  fragment CalendarEventActivity on ActivityEvent {   activityEventId   eventId   name   beginDate   endDate   isRegistered   description   currentStudentsCount   maxStudentCount   location   updateDate   isWaitListActive   isInWaitList   stopRegisterDate   __typename }  fragment Penalty on Penalty {   comment   id   duration   status   startTime   createTime   penaltySlot {     currentStudentsCount     description     duration     startTime     id     endTime     __typename   }   reasonId   __typename } ",
		variables,
	)

	return GqlRequest[Data_CalendarGetEvents](ctx, request)
}