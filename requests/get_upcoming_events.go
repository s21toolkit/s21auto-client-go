package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetUpcomingEvents struct {
	EventCodes                     []string `json:"eventCodes"`
	RegistrationAccessStatusFilter *string  `json:"registrationAccessStatusFilter,omitempty"`
	Variables_Page_GetUpcomingEvents                           *Variables_Page_GetUpcomingEvents    `json:"page,omitempty"`
}

type Variables_Page_GetUpcomingEvents struct {
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}


type Data_GetUpcomingEvents struct {
	Data_Student_GetUpcomingEvents Data_Student_GetUpcomingEvents `json:"student"`
}

type Data_Student_GetUpcomingEvents struct {
	Data_GetUpcomingEventsForRegistration_GetUpcomingEvents []Data_GetUpcomingEventsForRegistration_GetUpcomingEvents `json:"getUpcomingEventsForRegistration"`
	Typename                         string                             `json:"__typename"`
}

type Data_GetUpcomingEventsForRegistration_GetUpcomingEvents struct {
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
	Data_Activity_GetUpcomingEvents             Data_Activity_GetUpcomingEvents      `json:"activity"`
	Penalty              interface{}   `json:"penalty"`
	Typename             string        `json:"__typename"`
}

type Data_Activity_GetUpcomingEvents struct {
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
	Organizers           []Data_Organizer_GetUpcomingEvents `json:"organizers"`
	Typename             string      `json:"__typename"`
}

type Data_Organizer_GetUpcomingEvents struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) GetUpcomingEvents(variables Variables_GetUpcomingEvents) (Data_GetUpcomingEvents, error) {
	request := gql.NewQueryRequest[Variables_GetUpcomingEvents](
		"query getUpcomingEvents($eventCodes: [String!]!, $registrationAccessStatusFilter: RegistartionStatusEnum, $page: PagingInput) {\n  student {\n    getUpcomingEventsForRegistration(\n      eventCodes: $eventCodes\n      registrationAccessStatusFilter: $registrationAccessStatusFilter\n      page: $page\n    ) {\n      ...UpcomingEvent\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment UpcomingEvent on CalendarEvent {\n  id\n  start\n  end\n  bookings {\n    id\n    task {\n      id\n      goalName\n      __typename\n    }\n    __typename\n  }\n  eventSlots {\n    id\n    eventId\n    type\n    start\n    end\n    __typename\n  }\n  maxStudentCount\n  location\n  ipRange\n  eventType\n  eventCode\n  description\n  externalId\n  currentStudentsCount\n  exam {\n    examId\n    eventId\n    beginDate\n    endDate\n    location\n    ip\n    maxStudentCount\n    isVisible\n    name\n    goalId\n    isWaitListActive\n    isInWaitList\n    currentStudentsCount\n    createDate\n    updateDate\n    schoolId\n    stopRegisterDate\n    isRegistered\n    goalName\n    eventType\n    registrationAccessStatus\n    __typename\n  }\n  studentCodeReview {\n    studentGoalId\n    __typename\n  }\n  activity {\n    activityEventId\n    eventId\n    beginDate\n    endDate\n    location\n    description\n    maxStudentCount\n    isVisible\n    name\n    isWaitListActive\n    isInWaitList\n    currentStudentsCount\n    createDate\n    updateDate\n    schoolId\n    stopRegisterDate\n    isRegistered\n    activityType\n    eventType\n    isMandatory\n    status\n    organizers {\n      id\n      login\n      __typename\n    }\n    __typename\n  }\n  penalty {\n    ...Penalty\n    __typename\n  }\n  __typename\n}\n\nfragment Penalty on Penalty {\n  comment\n  id\n  duration\n  status\n  startTime\n  createTime\n  penaltySlot {\n    currentStudentsCount\n    description\n    duration\n    startTime\n    id\n    endTime\n    __typename\n  }\n  reasonId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetUpcomingEvents](ctx, request)
}