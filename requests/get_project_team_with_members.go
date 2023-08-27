package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetProjectTeamWithMembers struct {
	GoalID string `json:"goalId"`
}


type Data_GetProjectTeamWithMembers struct {
	Student Data_DataStudent_GetProjectTeamWithMembers `json:"student"`
}

type Data_DataStudent_GetProjectTeamWithMembers struct {
	Data_GetProjectTeamWithMembers_GetProjectTeamWithMembers Data_GetProjectTeamWithMembers_GetProjectTeamWithMembers `json:"getProjectTeamWithMembers"`
	Typename                  string                    `json:"__typename"`
}

type Data_GetProjectTeamWithMembers_GetProjectTeamWithMembers struct {
	Data_TeamWithMembers_GetProjectTeamWithMembers Data_TeamWithMembers_GetProjectTeamWithMembers  `json:"teamWithMembers"`
	InvitedStudents []Data_InvitedStudent_GetProjectTeamWithMembers `json:"invitedStudents"`
	Typename        string           `json:"__typename"`
}

type Data_InvitedStudent_GetProjectTeamWithMembers struct {
	Student  Data_InvitedStudentStudent_GetProjectTeamWithMembers `json:"student"`
	Typename string                `json:"__typename"`
}

type Data_InvitedStudentStudent_GetProjectTeamWithMembers struct {
	Data_User_GetProjectTeamWithMembers     Data_User_GetProjectTeamWithMembers   `json:"user"`
	Typename string `json:"__typename"`
}

type Data_User_GetProjectTeamWithMembers struct {
	ID             string         `json:"id"`
	AvatarURL      string         `json:"avatarUrl"`
	Login          string         `json:"login"`
	Data_UserExperience_GetProjectTeamWithMembers Data_UserExperience_GetProjectTeamWithMembers `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Data_UserExperience_GetProjectTeamWithMembers struct {
	Data_Level_GetProjectTeamWithMembers            Data_Level_GetProjectTeamWithMembers  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Data_Level_GetProjectTeamWithMembers struct {
	ID       int64  `json:"id"`
	Data_Range_GetProjectTeamWithMembers    Data_Range_GetProjectTeamWithMembers  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_GetProjectTeamWithMembers struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type Data_TeamWithMembers_GetProjectTeamWithMembers struct {
	Data_Team_GetProjectTeamWithMembers     Data_Team_GetProjectTeamWithMembers     `json:"team"`
	Members  []Data_Member_GetProjectTeamWithMembers `json:"members"`
	Typename string   `json:"__typename"`
}

type Data_Member_GetProjectTeamWithMembers struct {
	Role     string `json:"role"`
	Data_User_GetProjectTeamWithMembers     Data_User_GetProjectTeamWithMembers   `json:"user"`
	Typename string `json:"__typename"`
}

type Data_Team_GetProjectTeamWithMembers struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	MinTeamMemberCount int64  `json:"minTeamMemberCount"`
	MaxTeamMemberCount int64  `json:"maxTeamMemberCount"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) GetProjectTeamWithMembers(variables Variables_GetProjectTeamWithMembers) (Data_GetProjectTeamWithMembers, error) {
	request := gql.NewQueryRequest[Variables_GetProjectTeamWithMembers](
		"query getProjectTeamWithMembers($goalId: ID!) {\n  student {\n    getProjectTeamWithMembers(goalId: $goalId) {\n      ...TeamWithMembers\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment TeamWithMembers on ProjectTeamWithMembers {\n  teamWithMembers {\n    team {\n      id\n      name\n      status\n      minTeamMemberCount\n      maxTeamMemberCount\n      __typename\n    }\n    members {\n      ...TeamMemberWithRole\n      __typename\n    }\n    __typename\n  }\n  invitedStudents {\n    student {\n      user {\n        ...ProjectTeamMember\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment TeamMemberWithRole on TeamMember {\n  role\n  user {\n    ...ProjectTeamMember\n    __typename\n  }\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_GetProjectTeamWithMembers](ctx, request)
}