package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_CalendarGetEvents struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type Data_CalendarGetEvents struct {
	Data_Student_CalendarGetEvents Data_Student_CalendarGetEvents `json:"student"`
}

type Data_Student_CalendarGetEvents struct {
	GetMyCalendarEvents []Data_GetMyCalendarEvent_CalendarGetEvents `json:"getMyCalendarEvents"`
	Typename            string               `json:"__typename"`
}

type Data_GetMyCalendarEvent_CalendarGetEvents struct {
	ID                string        `json:"id"`
	Start             string        `json:"start"`
	End               string        `json:"end"`
	Description       string        `json:"description"`
	EventType         string        `json:"eventType"`
	EventCode         string        `json:"eventCode"`
	EventSlots        []Data_EventSlot_CalendarGetEvents   `json:"eventSlots"`
	Bookings          []Data_Booking_CalendarGetEvents     `json:"bookings"`
	Exam              interface{}   `json:"exam"`
	StudentCodeReview interface{}   `json:"studentCodeReview"`
	Data_Activity_CalendarGetEvents          *Data_Activity_CalendarGetEvents     `json:"activity"`
	Goals             []interface{} `json:"goals"`
	Penalty           interface{}   `json:"penalty"`
	Typename          string        `json:"__typename"`
}

type Data_Activity_CalendarGetEvents struct {
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
	Data_StudentFeedback_CalendarGetEvents      Data_StudentFeedback_CalendarGetEvents `json:"studentFeedback"`
	Status               string          `json:"status"`
	ActivityType         string          `json:"activityType"`
	IsMandatory          bool            `json:"isMandatory"`
	IsVisible            bool            `json:"isVisible"`
	Comments             []interface{}   `json:"comments"`
	Organizers           []Data_VerifierUser_CalendarGetEvents  `json:"organizers"`
}

type Data_VerifierUser_CalendarGetEvents struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}

type Data_StudentFeedback_CalendarGetEvents struct {
	ID       string `json:"id"`
	Rating   int64  `json:"rating"`
	Comment  string `json:"comment"`
	Typename string `json:"__typename"`
}

type Data_Booking_CalendarGetEvents struct {
	ID                string             `json:"id"`
	AnswerID          *string            `json:"answerId"`
	EventSlotID       string             `json:"eventSlotId"`
	Data_Task_CalendarGetEvents              *Data_Task_CalendarGetEvents              `json:"task"`
	Data_EventSlot_CalendarGetEvents         Data_EventSlot_CalendarGetEvents          `json:"eventSlot"`
	Data_VerifierUser_CalendarGetEvents      Data_VerifierUser_CalendarGetEvents       `json:"verifierUser"`
	Data_VerifiableStudent_CalendarGetEvents *Data_VerifiableStudent_CalendarGetEvents `json:"verifiableStudent"`
	BookingStatus     string             `json:"bookingStatus"`
	Data_Team_CalendarGetEvents              *Data_Team_CalendarGetEvents              `json:"team"`
	IsOnline          bool               `json:"isOnline"`
	VcLinkURL         interface{}        `json:"vcLinkUrl"`
	Typename          string             `json:"__typename"`
}

type Data_EventSlot_CalendarGetEvents struct {
	ID       string  `json:"id"`
	Start    string  `json:"start"`
	End      string  `json:"end"`
	Data_Event_CalendarGetEvents    *Data_Event_CalendarGetEvents  `json:"event,omitempty"`
	Typename string  `json:"__typename"`
	Type     *string `json:"type,omitempty"`
}

type Data_Event_CalendarGetEvents struct {
	EventUserRole string `json:"eventUserRole"`
	Typename      string `json:"__typename"`
}

type Data_Task_CalendarGetEvents struct {
	ID                              string                          `json:"id"`
	GoalID                          string                          `json:"goalId"`
	GoalName                        string                          `json:"goalName"`
	Data_StudentTaskAdditionalAttributes_CalendarGetEvents Data_StudentTaskAdditionalAttributes_CalendarGetEvents `json:"studentTaskAdditionalAttributes"`
	AssignmentType                  string                          `json:"assignmentType"`
	Typename                        string                          `json:"__typename"`
}

type Data_StudentTaskAdditionalAttributes_CalendarGetEvents struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}

type Data_Team_CalendarGetEvents struct {
	ID                 string        `json:"id"`
	Data_TeamLead_CalendarGetEvents           Data_TeamLead_CalendarGetEvents      `json:"teamLead"`
	Members            []Data_TeamLead_CalendarGetEvents    `json:"members"`
	InvitedUsers       []interface{} `json:"invitedUsers"`
	TeamName           string        `json:"teamName"`
	TeamStatus         string        `json:"teamStatus"`
	MinTeamMemberCount int64         `json:"minTeamMemberCount"`
	MaxTeamMemberCount int64         `json:"maxTeamMemberCount"`
	Typename           string        `json:"__typename"`
}

type Data_TeamLead_CalendarGetEvents struct {
	ID             string         `json:"id"`
	AvatarURL      string         `json:"avatarUrl"`
	Login          string         `json:"login"`
	Data_UserExperience_CalendarGetEvents Data_UserExperience_CalendarGetEvents `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Data_UserExperience_CalendarGetEvents struct {
	Data_Level_CalendarGetEvents            Data_Level_CalendarGetEvents  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Data_Level_CalendarGetEvents struct {
	ID       int64  `json:"id"`
	Data_Range_CalendarGetEvents    Data_Range_CalendarGetEvents  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_CalendarGetEvents struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type Data_VerifiableStudent_CalendarGetEvents struct {
	ID       string       `json:"id"`
	User     Data_VerifierUser_CalendarGetEvents `json:"user"`
	Typename string       `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetEvents(variables Variables_CalendarGetEvents) (Data_CalendarGetEvents, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetEvents](
		"query calendarGetEvents($from: DateTime!, $to: DateTime!) {\n  student {\n    getMyCalendarEvents(from: $from, to: $to) {\n      ...CalendarEvent\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarEvent on CalendarEvent {\n  id\n  start\n  end\n  description\n  eventType\n  eventCode\n  eventSlots {\n    id\n    type\n    start\n    end\n    __typename\n  }\n  bookings {\n    ...CalendarReviewBooking\n    __typename\n  }\n  exam {\n    ...CalendarEventExam\n    __typename\n  }\n  studentCodeReview {\n    studentGoalId\n    __typename\n  }\n  activity {\n    ...CalendarEventActivity\n    studentFeedback {\n      id\n      rating\n      comment\n      __typename\n    }\n    status\n    activityType\n    isMandatory\n    isWaitListActive\n    isVisible\n    comments {\n      type\n      createTs\n      comment\n      __typename\n    }\n    organizers {\n      id\n      login\n      __typename\n    }\n    __typename\n  }\n  goals {\n    goalId\n    goalName\n    __typename\n  }\n  penalty {\n    ...Penalty\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarReviewBooking on CalendarBooking {\n  id\n  answerId\n  eventSlotId\n  task {\n    id\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    assignmentType\n    __typename\n  }\n  eventSlot {\n    id\n    start\n    end\n    event {\n      eventUserRole\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...CalendarReviewUser\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...CalendarReviewUser\n      __typename\n    }\n    __typename\n  }\n  bookingStatus\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment CalendarReviewUser on User {\n  id\n  login\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarEventExam on Exam {\n  examId\n  eventId\n  beginDate\n  endDate\n  name\n  location\n  currentStudentsCount\n  maxStudentCount\n  updateDate\n  goalId\n  goalName\n  isWaitListActive\n  isInWaitList\n  stopRegisterDate\n  __typename\n}\n\nfragment CalendarEventActivity on ActivityEvent {\n  activityEventId\n  eventId\n  name\n  beginDate\n  endDate\n  isRegistered\n  description\n  currentStudentsCount\n  maxStudentCount\n  location\n  updateDate\n  isWaitListActive\n  isInWaitList\n  stopRegisterDate\n  __typename\n}\n\nfragment Penalty on Penalty {\n  comment\n  id\n  duration\n  status\n  startTime\n  createTime\n  penaltySlot {\n    currentStudentsCount\n    description\n    duration\n    startTime\n    id\n    endTime\n    __typename\n  }\n  reasonId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_CalendarGetEvents](ctx, request)
}