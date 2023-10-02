package requests

import "github.com/s21toolkit/s21client/gql"

type CalendarGetEvents_Variables struct {
	From string `json:"from"`
	To   string `json:"to"`
}


type CalendarGetEvents_Data struct {
	Student CalendarGetEvents_Data_Student `json:"student"`
}

type CalendarGetEvents_Data_Student struct {
	GetMyCalendarEvents []CalendarGetEvents_Data_GetMyCalendarEvent `json:"getMyCalendarEvents"`
	Typename            string               `json:"__typename"`
}

type CalendarGetEvents_Data_GetMyCalendarEvent struct {
	ID                string        `json:"id"`
	Start             string        `json:"start"`
	End               string        `json:"end"`
	Description       string        `json:"description"`
	EventType         string        `json:"eventType"`
	EventCode         string        `json:"eventCode"`
	EventSlots        []CalendarGetEvents_Data_EventSlot   `json:"eventSlots"`
	Bookings          []CalendarGetEvents_Data_Booking     `json:"bookings"`
	Exam              interface{}   `json:"exam"`
	StudentCodeReview interface{}   `json:"studentCodeReview"`
	Activity          *CalendarGetEvents_Data_Activity     `json:"activity"`
	Goals             []interface{} `json:"goals"`
	Penalty           interface{}   `json:"penalty"`
	Typename          string        `json:"__typename"`
}

type CalendarGetEvents_Data_Activity struct {
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
	StudentFeedback      CalendarGetEvents_Data_StudentFeedback `json:"studentFeedback"`
	Status               string          `json:"status"`
	ActivityType         string          `json:"activityType"`
	IsMandatory          bool            `json:"isMandatory"`
	IsVisible            bool            `json:"isVisible"`
	Comments             []interface{}   `json:"comments"`
	Organizers           []CalendarGetEvents_Data_VerifierUser  `json:"organizers"`
}

type CalendarGetEvents_Data_VerifierUser struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Typename string `json:"__typename"`
}

type CalendarGetEvents_Data_StudentFeedback struct {
	ID       string `json:"id"`
	Rating   int64  `json:"rating"`
	Comment  string `json:"comment"`
	Typename string `json:"__typename"`
}

type CalendarGetEvents_Data_Booking struct {
	ID                string             `json:"id"`
	AnswerID          *string            `json:"answerId"`
	EventSlotID       string             `json:"eventSlotId"`
	Task              *CalendarGetEvents_Data_Task              `json:"task"`
	EventSlot         CalendarGetEvents_Data_EventSlot          `json:"eventSlot"`
	VerifierUser      CalendarGetEvents_Data_VerifierUser       `json:"verifierUser"`
	VerifiableStudent *CalendarGetEvents_Data_VerifiableStudent `json:"verifiableStudent"`
	BookingStatus     string             `json:"bookingStatus"`
	Team              *CalendarGetEvents_Data_Team              `json:"team"`
	IsOnline          bool               `json:"isOnline"`
	VcLinkURL         interface{}        `json:"vcLinkUrl"`
	Typename          string             `json:"__typename"`
}

type CalendarGetEvents_Data_EventSlot struct {
	ID       string  `json:"id"`
	Start    string  `json:"start"`
	End      string  `json:"end"`
	Event    *CalendarGetEvents_Data_Event  `json:"event,omitempty"`
	Typename string  `json:"__typename"`
	Type     *string `json:"type,omitempty"`
}

type CalendarGetEvents_Data_Event struct {
	EventUserRole string `json:"eventUserRole"`
	Typename      string `json:"__typename"`
}

type CalendarGetEvents_Data_Task struct {
	ID                              string                          `json:"id"`
	GoalID                          string                          `json:"goalId"`
	GoalName                        string                          `json:"goalName"`
	StudentTaskAdditionalAttributes CalendarGetEvents_Data_StudentTaskAdditionalAttributes `json:"studentTaskAdditionalAttributes"`
	AssignmentType                  string                          `json:"assignmentType"`
	Typename                        string                          `json:"__typename"`
}

type CalendarGetEvents_Data_StudentTaskAdditionalAttributes struct {
	CookiesCount int64  `json:"cookiesCount"`
	Typename     string `json:"__typename"`
}

type CalendarGetEvents_Data_Team struct {
	ID                 string        `json:"id"`
	TeamLead           CalendarGetEvents_Data_TeamLead      `json:"teamLead"`
	Members            []CalendarGetEvents_Data_TeamLead    `json:"members"`
	InvitedUsers       []interface{} `json:"invitedUsers"`
	TeamName           string        `json:"teamName"`
	TeamStatus         string        `json:"teamStatus"`
	MinTeamMemberCount int64         `json:"minTeamMemberCount"`
	MaxTeamMemberCount int64         `json:"maxTeamMemberCount"`
	Typename           string        `json:"__typename"`
}

type CalendarGetEvents_Data_TeamLead struct {
	ID             string         `json:"id"`
	AvatarURL      string         `json:"avatarUrl"`
	Login          string         `json:"login"`
	UserExperience CalendarGetEvents_Data_UserExperience `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type CalendarGetEvents_Data_UserExperience struct {
	Level            CalendarGetEvents_Data_Level  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type CalendarGetEvents_Data_Level struct {
	ID       int64  `json:"id"`
	Range    CalendarGetEvents_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type CalendarGetEvents_Data_Range struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type CalendarGetEvents_Data_VerifiableStudent struct {
	ID       string       `json:"id"`
	User     CalendarGetEvents_Data_VerifierUser `json:"user"`
	Typename string       `json:"__typename"`
}


func (ctx *RequestContext) CalendarGetEvents(variables CalendarGetEvents_Variables) (CalendarGetEvents_Data, error) {
	request := gql.NewQueryRequest[CalendarGetEvents_Variables](
		"query calendarGetEvents($from: DateTime!, $to: DateTime!) {\n  student {\n    getMyCalendarEvents(from: $from, to: $to) {\n      ...CalendarEvent\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment CalendarEvent on CalendarEvent {\n  id\n  start\n  end\n  description\n  eventType\n  eventCode\n  eventSlots {\n    id\n    type\n    start\n    end\n    __typename\n  }\n  bookings {\n    ...CalendarReviewBooking\n    __typename\n  }\n  exam {\n    ...CalendarEventExam\n    __typename\n  }\n  studentCodeReview {\n    studentGoalId\n    __typename\n  }\n  activity {\n    ...CalendarEventActivity\n    studentFeedback {\n      id\n      rating\n      comment\n      __typename\n    }\n    status\n    activityType\n    isMandatory\n    isWaitListActive\n    isVisible\n    comments {\n      type\n      createTs\n      comment\n      __typename\n    }\n    organizers {\n      id\n      login\n      __typename\n    }\n    __typename\n  }\n  goals {\n    goalId\n    goalName\n    __typename\n  }\n  penalty {\n    ...Penalty\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarReviewBooking on CalendarBooking {\n  id\n  answerId\n  eventSlotId\n  task {\n    id\n    goalId\n    goalName\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      __typename\n    }\n    assignmentType\n    __typename\n  }\n  eventSlot {\n    id\n    start\n    end\n    event {\n      eventUserRole\n      __typename\n    }\n    __typename\n  }\n  verifierUser {\n    ...CalendarReviewUser\n    __typename\n  }\n  verifiableStudent {\n    id\n    user {\n      ...CalendarReviewUser\n      __typename\n    }\n    __typename\n  }\n  bookingStatus\n  team {\n    ...ProjectTeamMembers\n    __typename\n  }\n  isOnline\n  vcLinkUrl\n  __typename\n}\n\nfragment CalendarReviewUser on User {\n  id\n  login\n  __typename\n}\n\nfragment ProjectTeamMembers on ProjectTeamMembers {\n  id\n  teamLead {\n    ...ProjectTeamMember\n    __typename\n  }\n  members {\n    ...ProjectTeamMember\n    __typename\n  }\n  invitedUsers {\n    ...ProjectTeamMember\n    __typename\n  }\n  teamName\n  teamStatus\n  minTeamMemberCount\n  maxTeamMemberCount\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n\nfragment CalendarEventExam on Exam {\n  examId\n  eventId\n  beginDate\n  endDate\n  name\n  location\n  currentStudentsCount\n  maxStudentCount\n  updateDate\n  goalId\n  goalName\n  isWaitListActive\n  isInWaitList\n  stopRegisterDate\n  __typename\n}\n\nfragment CalendarEventActivity on ActivityEvent {\n  activityEventId\n  eventId\n  name\n  beginDate\n  endDate\n  isRegistered\n  description\n  currentStudentsCount\n  maxStudentCount\n  location\n  updateDate\n  isWaitListActive\n  isInWaitList\n  stopRegisterDate\n  __typename\n}\n\nfragment Penalty on Penalty {\n  comment\n  id\n  duration\n  status\n  startTime\n  createTime\n  penaltySlot {\n    currentStudentsCount\n    description\n    duration\n    startTime\n    id\n    endTime\n    __typename\n  }\n  reasonId\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CalendarGetEvents_Data](ctx, request)
}