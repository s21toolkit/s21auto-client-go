package requests

import "s21client/gql"

type Request_Variables_UnsubscribeFromEvent struct {
	EventID string `json:"eventId"`
}


type Response_Data_UnsubscribeFromEvent struct {
	Response_Student_UnsubscribeFromEvent Response_Student_UnsubscribeFromEvent `json:"student"`
}

type Response_Student_UnsubscribeFromEvent struct {
	Response_UnsubscribeFromEvent_UnsubscribeFromEvent Response_UnsubscribeFromEvent_UnsubscribeFromEvent `json:"unsubscribeFromEvent"`
	Typename             string               `json:"__typename"`
}

type Response_UnsubscribeFromEvent_UnsubscribeFromEvent struct {
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
	Response_Activity_UnsubscribeFromEvent             Response_Activity_UnsubscribeFromEvent      `json:"activity"`
	Penalty              interface{}   `json:"penalty"`
	Typename             string        `json:"__typename"`
}

type Response_Activity_UnsubscribeFromEvent struct {
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
	Organizers           []Response_Organizer_UnsubscribeFromEvent `json:"organizers"`
	Typename             string      `json:"__typename"`
}

type Response_Organizer_UnsubscribeFromEvent struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) UnsubscribeFromEvent(variables Request_Variables_UnsubscribeFromEvent) (Response_Data_UnsubscribeFromEvent, error) {
	request := gql.NewQueryRequest[Request_Variables_UnsubscribeFromEvent](
		"mutation unsubscribeFromEvent($eventId: ID!) {   student {     unsubscribeFromEvent(eventId: $eventId) {       ...UpcomingEvent       __typename     }     __typename   } }  fragment UpcomingEvent on CalendarEvent {   id   start   end   bookings {     id     task {       id       goalName       __typename     }     __typename   }   eventSlots {     id     eventId     type     start     end     __typename   }   maxStudentCount   location   ipRange   eventType   eventCode   description   externalId   currentStudentsCount   exam {     examId     eventId     beginDate     endDate     location     ip     maxStudentCount     isVisible     name     goalId     isWaitListActive     isInWaitList     currentStudentsCount     createDate     updateDate     schoolId     stopRegisterDate     isRegistered     goalName     eventType     registrationAccessStatus     __typename   }   studentCodeReview {     studentGoalId     __typename   }   activity {     activityEventId     eventId     beginDate     endDate     location     description     maxStudentCount     isVisible     name     isWaitListActive     isInWaitList     currentStudentsCount     createDate     updateDate     schoolId     stopRegisterDate     isRegistered     activityType     eventType     isMandatory     status     organizers {       id       login       __typename     }     __typename   }   penalty {     ...Penalty     __typename   }   __typename }  fragment Penalty on Penalty {   comment   id   duration   status   startTime   createTime   penaltySlot {     currentStudentsCount     description     duration     startTime     id     endTime     __typename   }   reasonId   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_UnsubscribeFromEvent](ctx, request)
}