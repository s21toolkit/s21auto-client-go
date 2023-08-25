package requests

import "s21client/gql"

type Request_Variables_GetUpcomingEvents struct {
	EventCodes                     []string `json:"eventCodes"`
	RegistrationAccessStatusFilter *string  `json:"registrationAccessStatusFilter,omitempty"`
	Request_Page_GetUpcomingEvents                           *Request_Page_GetUpcomingEvents    `json:"page,omitempty"`
}

type Request_Page_GetUpcomingEvents struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}


type Response_Data_GetUpcomingEvents struct {
	Response_Student_GetUpcomingEvents Response_Student_GetUpcomingEvents `json:"student"`
}

type Response_Student_GetUpcomingEvents struct {
	Response_GetUpcomingEventsForRegistration_GetUpcomingEvents []Response_GetUpcomingEventsForRegistration_GetUpcomingEvents `json:"getUpcomingEventsForRegistration"`
	Typename                         string                             `json:"__typename"`
}

type Response_GetUpcomingEventsForRegistration_GetUpcomingEvents struct {
	ID                   string        `json:"id"`
	Start                string        `json:"start"`
	End                  string        `json:"end"`
	Bookings             []interface{} `json:"bookings"`
	EventSlots           []interface{} `json:"eventSlots"`
	MaxStudentCount      int64         `json:"maxStudentCount"`
	Location             string        `json:"location"`
	IPRange              string        `json:"ipRange"`
	EventType            string        `json:"eventType"`
	EventCode            string        `json:"eventCode"`
	Description          string        `json:"description"`
	ExternalID           int64         `json:"externalId"`
	CurrentStudentsCount int64         `json:"currentStudentsCount"`
	Exam                 interface{}   `json:"exam"`
	StudentCodeReview    interface{}   `json:"studentCodeReview"`
	Response_Activity_GetUpcomingEvents             Response_Activity_GetUpcomingEvents      `json:"activity"`
	Penalty              interface{}   `json:"penalty"`
	Typename             string        `json:"__typename"`
}

type Response_Activity_GetUpcomingEvents struct {
	ActivityEventID      string      `json:"activityEventId"`
	EventID              string      `json:"eventId"`
	BeginDate            string      `json:"beginDate"`
	EndDate              string      `json:"endDate"`
	Location             string      `json:"location"`
	Description          string      `json:"description"`
	MaxStudentCount      int64       `json:"maxStudentCount"`
	IsVisible            bool        `json:"isVisible"`
	Name                 string      `json:"name"`
	IsWaitListActive     bool        `json:"isWaitListActive"`
	IsInWaitList         bool        `json:"isInWaitList"`
	CurrentStudentsCount int64       `json:"currentStudentsCount"`
	CreateDate           string      `json:"createDate"`
	UpdateDate           string      `json:"updateDate"`
	SchoolID             string      `json:"schoolId"`
	StopRegisterDate     string      `json:"stopRegisterDate"`
	IsRegistered         bool        `json:"isRegistered"`
	ActivityType         string      `json:"activityType"`
	EventType            string      `json:"eventType"`
	IsMandatory          bool        `json:"isMandatory"`
	Status               string      `json:"status"`
	Organizers           []Response_Organizer_GetUpcomingEvents `json:"organizers"`
	Typename             string      `json:"__typename"`
}

type Response_Organizer_GetUpcomingEvents struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetUpcomingEvents(variables Request_Variables_GetUpcomingEvents) (Response_Data_GetUpcomingEvents, error) {
	request := gql.NewQueryRequest[Request_Variables_GetUpcomingEvents](
		"query getUpcomingEvents($eventCodes: [String!]!, $registrationAccessStatusFilter: RegistartionStatusEnum, $page: PagingInput) {   student {     getUpcomingEventsForRegistration(       eventCodes: $eventCodes       registrationAccessStatusFilter: $registrationAccessStatusFilter       page: $page     ) {       ...UpcomingEvent       __typename     }     __typename   } }  fragment UpcomingEvent on CalendarEvent {   id   start   end   bookings {     id     task {       id       goalName       __typename     }     __typename   }   eventSlots {     id     eventId     type     start     end     __typename   }   maxStudentCount   location   ipRange   eventType   eventCode   description   externalId   currentStudentsCount   exam {     examId     eventId     beginDate     endDate     location     ip     maxStudentCount     isVisible     name     goalId     isWaitListActive     isInWaitList     currentStudentsCount     createDate     updateDate     schoolId     stopRegisterDate     isRegistered     goalName     eventType     registrationAccessStatus     __typename   }   studentCodeReview {     studentGoalId     __typename   }   activity {     activityEventId     eventId     beginDate     endDate     location     description     maxStudentCount     isVisible     name     isWaitListActive     isInWaitList     currentStudentsCount     createDate     updateDate     schoolId     stopRegisterDate     isRegistered     activityType     eventType     isMandatory     status     organizers {       id       login       __typename     }     __typename   }   penalty {     ...Penalty     __typename   }   __typename }  fragment Penalty on Penalty {   comment   id   duration   status   startTime   createTime   penaltySlot {     currentStudentsCount     description     duration     startTime     id     endTime     __typename   }   reasonId   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_GetUpcomingEvents](ctx, request)
}