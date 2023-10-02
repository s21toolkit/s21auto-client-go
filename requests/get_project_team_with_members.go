package requests

import "github.com/s21toolkit/s21client/gql"

type GetProjectTeamWithMembers_Variables struct {
	GoalID string `json:"goalId"`
}


type GetProjectTeamWithMembers_Data struct {
	Student GetProjectTeamWithMembers_Data_DataStudent `json:"student"`
}

type GetProjectTeamWithMembers_Data_DataStudent struct {
	GetProjectTeamWithMembers GetProjectTeamWithMembers_Data_GetProjectTeamWithMembers `json:"getProjectTeamWithMembers"`
	Typename                  string                    `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_GetProjectTeamWithMembers struct {
	TeamWithMembers GetProjectTeamWithMembers_Data_TeamWithMembers  `json:"teamWithMembers"`
	InvitedStudents []GetProjectTeamWithMembers_Data_InvitedStudent `json:"invitedStudents"`
	Typename        string           `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_InvitedStudent struct {
	Student  GetProjectTeamWithMembers_Data_InvitedStudentStudent `json:"student"`
	Typename string                `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_InvitedStudentStudent struct {
	User     GetProjectTeamWithMembers_Data_User   `json:"user"`
	Typename string `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_User struct {
	ID             string         `json:"id"`
	AvatarURL      string         `json:"avatarUrl"`
	Login          string         `json:"login"`
	UserExperience GetProjectTeamWithMembers_Data_UserExperience `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_UserExperience struct {
	Level            GetProjectTeamWithMembers_Data_Level  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_Level struct {
	ID       int64  `json:"id"`
	Range    GetProjectTeamWithMembers_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_Range struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_TeamWithMembers struct {
	Team     GetProjectTeamWithMembers_Data_Team     `json:"team"`
	Members  []GetProjectTeamWithMembers_Data_Member `json:"members"`
	Typename string   `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_Member struct {
	Role     string `json:"role"`
	User     GetProjectTeamWithMembers_Data_User   `json:"user"`
	Typename string `json:"__typename"`
}

type GetProjectTeamWithMembers_Data_Team struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	MinTeamMemberCount int64  `json:"minTeamMemberCount"`
	MaxTeamMemberCount int64  `json:"maxTeamMemberCount"`
	Typename           string `json:"__typename"`
}


func (ctx *RequestContext) GetProjectTeamWithMembers(variables GetProjectTeamWithMembers_Variables) (GetProjectTeamWithMembers_Data, error) {
	request := gql.NewQueryRequest[GetProjectTeamWithMembers_Variables](
		"query getProjectTeamWithMembers($goalId: ID!) {\n  student {\n    getProjectTeamWithMembers(goalId: $goalId) {\n      ...TeamWithMembers\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment TeamWithMembers on ProjectTeamWithMembers {\n  teamWithMembers {\n    team {\n      id\n      name\n      status\n      minTeamMemberCount\n      maxTeamMemberCount\n      __typename\n    }\n    members {\n      ...TeamMemberWithRole\n      __typename\n    }\n    __typename\n  }\n  invitedStudents {\n    student {\n      user {\n        ...ProjectTeamMember\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment TeamMemberWithRole on TeamMember {\n  role\n  user {\n    ...ProjectTeamMember\n    __typename\n  }\n  __typename\n}\n\nfragment ProjectTeamMember on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetProjectTeamWithMembers_Data](ctx, request)
}