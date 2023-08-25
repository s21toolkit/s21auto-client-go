package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetProjectTeamWithMembers struct {
	GoalID string `json:"goalId"`
}


type Response_Data_GetProjectTeamWithMembers struct {
	Student Response_DataStudent_GetProjectTeamWithMembers `json:"student"`
}

type Response_DataStudent_GetProjectTeamWithMembers struct {
	Response_GetProjectTeamWithMembers_GetProjectTeamWithMembers Response_GetProjectTeamWithMembers_GetProjectTeamWithMembers `json:"getProjectTeamWithMembers"`
	Typename                  string                    `json:"__typename"`
}

type Response_GetProjectTeamWithMembers_GetProjectTeamWithMembers struct {
	Response_TeamWithMembers_GetProjectTeamWithMembers Response_TeamWithMembers_GetProjectTeamWithMembers  `json:"teamWithMembers"`
	InvitedStudents []Response_InvitedStudent_GetProjectTeamWithMembers `json:"invitedStudents"`
	Typename        string           `json:"__typename"`
}

type Response_InvitedStudent_GetProjectTeamWithMembers struct {
	Student  Response_InvitedStudentStudent_GetProjectTeamWithMembers `json:"student"`
	Typename string                `json:"__typename"`
}

type Response_InvitedStudentStudent_GetProjectTeamWithMembers struct {
	Response_User_GetProjectTeamWithMembers     Response_User_GetProjectTeamWithMembers   `json:"user"`
	Typename string `json:"__typename"`
}

type Response_User_GetProjectTeamWithMembers struct {
	ID             string         `json:"id"`
	AvatarURL      string         `json:"avatarUrl"`
	Login          string         `json:"login"`
	Response_UserExperience_GetProjectTeamWithMembers Response_UserExperience_GetProjectTeamWithMembers `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Response_UserExperience_GetProjectTeamWithMembers struct {
	Response_Level_GetProjectTeamWithMembers            Response_Level_GetProjectTeamWithMembers  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Response_Level_GetProjectTeamWithMembers struct {
	ID       int64  `json:"id"`
	Response_Range_GetProjectTeamWithMembers    Response_Range_GetProjectTeamWithMembers  `json:"range"`
	Typename string `json:"__typename"`
}

type Response_Range_GetProjectTeamWithMembers struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type Response_TeamWithMembers_GetProjectTeamWithMembers struct {
	Response_Team_GetProjectTeamWithMembers     Response_Team_GetProjectTeamWithMembers     `json:"team"`
	Members  []Response_Member_GetProjectTeamWithMembers `json:"members"`
	Typename string   `json:"__typename"`
}

type Response_Member_GetProjectTeamWithMembers struct {
	Role     string `json:"role"`
	Response_User_GetProjectTeamWithMembers     Response_User_GetProjectTeamWithMembers   `json:"user"`
	Typename string `json:"__typename"`
}

type Response_Team_GetProjectTeamWithMembers struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	MinTeamMemberCount int64  `json:"minTeamMemberCount"`
	MaxTeamMemberCount int64  `json:"maxTeamMemberCount"`
	Typename           string `json:"__typename"`
}

func (ctx *RequestContext) GetProjectTeamWithMembers(variables Request_Variables_GetProjectTeamWithMembers) (Response_Data_GetProjectTeamWithMembers, error) {
	request := gql.NewQueryRequest[Request_Variables_GetProjectTeamWithMembers](
		"query getProjectTeamWithMembers($goalId: ID!) {\n  student {\n    getProjectTeamWithMembers(goalId: $goalId) {\n      ...TeamWithMembers\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment TeamWithMembers on ProjectTeamWithMembers {\n  teamWithMembers {\n    team {\n      id\n      name\n      status\n      minTeamMemberCount\n      maxTeamMemberCount\n      __typename\n    }\n    members {\n      ...TeamMemberWithRole\n      __typename\n    }\n    __typename\n  }\n  invitedStudents {\n    student {\n      user {\n        ...ProjectTeamMember\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment TeamMemberWithRole on TeamMember {\n  role\n  user {\n    ...ProjectTeamMember\n    __typename\n  }\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_GetProjectTeamWithMembers](ctx, request)
}