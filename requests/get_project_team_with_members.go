package requests

import "s21client/gql"

type Variables_GetProjectTeamWithMembers struct {
	GoalID string `json:"goalId"`
}


type Data_GetProjectTeamWithMembers struct {
	Student_GetProjectTeamWithMembers Student_GetProjectTeamWithMembers `json:"student"`
}

type Student_GetProjectTeamWithMembers struct {
	GetProjectTeamWithMembers_GetProjectTeamWithMembers GetProjectTeamWithMembers_GetProjectTeamWithMembers `json:"getProjectTeamWithMembers"`
	Typename                  string                    `json:"__typename"`
}

type GetProjectTeamWithMembers_GetProjectTeamWithMembers struct {
	TeamWithMembers_GetProjectTeamWithMembers TeamWithMembers_GetProjectTeamWithMembers `json:"teamWithMembers"`
	InvitedStudents []interface{}   `json:"invitedStudents"`
	Typename        string          `json:"__typename"`
}

type TeamWithMembers_GetProjectTeamWithMembers struct {
	Team_GetProjectTeamWithMembers     Team_GetProjectTeamWithMembers     `json:"team"`
	Members  []Member_GetProjectTeamWithMembers `json:"members"`
	Typename string   `json:"__typename"`
}

type Member_GetProjectTeamWithMembers struct {
	Role     string `json:"role"`
	User_GetProjectTeamWithMembers     User_GetProjectTeamWithMembers   `json:"user"`
	Typename string `json:"__typename"`
}

type User_GetProjectTeamWithMembers struct {
	ID             string         `json:"id"`
	AvatarURL      string         `json:"avatarUrl"`
	Login          string         `json:"login"`
	UserExperience_GetProjectTeamWithMembers UserExperience_GetProjectTeamWithMembers `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type UserExperience_GetProjectTeamWithMembers struct {
	Level_GetProjectTeamWithMembers            Level_GetProjectTeamWithMembers  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type Level_GetProjectTeamWithMembers struct {
	ID       int64  `json:"id"`
	Range_GetProjectTeamWithMembers    Range_GetProjectTeamWithMembers  `json:"range"`
	Typename string `json:"__typename"`
}

type Range_GetProjectTeamWithMembers struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type Team_GetProjectTeamWithMembers struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Status             string `json:"status"`
	MinTeamMemberCount int64  `json:"minTeamMemberCount"`
	MaxTeamMemberCount int64  `json:"maxTeamMemberCount"`
	Typename           string `json:"__typename"`
}

func (ctx *RequestContext) GetProjectTeamWithMembers(variables Variables_GetProjectTeamWithMembers) (Data_GetProjectTeamWithMembers, error) {
	request := gql.NewQueryRequest[Variables_GetProjectTeamWithMembers](
		"query getProjectTeamWithMembers($goalId: ID!) {   student {     getProjectTeamWithMembers(goalId: $goalId) {       ...TeamWithMembers       __typename     }     __typename   } }  fragment TeamWithMembers on ProjectTeamWithMembers {   teamWithMembers {     team {       id       name       status       minTeamMemberCount       maxTeamMemberCount       __typename     }     members {       ...TeamMemberWithRole       __typename     }     __typename   }   invitedStudents {     student {       user {         ...ProjectTeamMember         __typename       }       __typename     }     __typename   }   __typename }  fragment TeamMemberWithRole on TeamMember {   role   user {     ...ProjectTeamMember     __typename   }   __typename }  fragment ProjectTeamMember on User {   id   avatarUrl   login   userExperience {     level {       id       range {         levelCode         __typename       }       __typename     }     cookiesCount     codeReviewPoints     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetProjectTeamWithMembers](ctx, request)
}