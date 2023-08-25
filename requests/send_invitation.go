package requests

import "s21client/gql"

type Request_Variables_SendInvitation struct {
	UserID string `json:"userId"`
	TeamID string `json:"teamId"`
}


type Response_Data_SendInvitation struct {
	Student Response_DataStudent_SendInvitation `json:"student"`
}

type Response_DataStudent_SendInvitation struct {
	Response_SendInvitation_SendInvitation Response_SendInvitation_SendInvitation `json:"sendInvitation"`
	Typename       string         `json:"__typename"`
}

type Response_SendInvitation_SendInvitation struct {
	Student          Response_SendInvitationStudent_SendInvitation `json:"student"`
	InvitationStatus string                `json:"invitationStatus"`
	Typename         string                `json:"__typename"`
}

type Response_SendInvitationStudent_SendInvitation struct {
	ID       string `json:"id"`
	Response_User_SendInvitation     Response_User_SendInvitation   `json:"user"`
	Typename string `json:"__typename"`
}

type Response_User_SendInvitation struct {
	ID             string         `json:"id"`
	Login          string         `json:"login"`
	AvatarURL      string         `json:"avatarUrl"`
	Response_UserExperience_SendInvitation Response_UserExperience_SendInvitation `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type Response_UserExperience_SendInvitation struct {
	ID               string `json:"id"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	CoinsCount       int64  `json:"coinsCount"`
	Response_Level_SendInvitation            Response_Level_SendInvitation  `json:"level"`
	Typename         string `json:"__typename"`
}

type Response_Level_SendInvitation struct {
	ID       int64  `json:"id"`
	Response_Range_SendInvitation    Response_Range_SendInvitation  `json:"range"`
	Typename string `json:"__typename"`
}

type Response_Range_SendInvitation struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) SendInvitation(variables Request_Variables_SendInvitation) (Response_Data_SendInvitation, error) {
	request := gql.NewQueryRequest[Request_Variables_SendInvitation](
		"mutation sendInvitation($teamId: UUID!, $userId: ID!) {   student {     sendInvitation(teamId: $teamId, userId: $userId) {       ...StudentInvitationInfo       __typename     }     __typename   } }  fragment StudentInvitationInfo on StudentInvitationInfo {   student {     ...AvailableStudentForTeam     __typename   }   invitationStatus   __typename }  fragment AvailableStudentForTeam on Student {   id   user {     id     login     avatarUrl     userExperience {       ...CurrentUserExperience       __typename     }     __typename   }   __typename }  fragment CurrentUserExperience on UserExperience {   id   cookiesCount   codeReviewPoints   coinsCount   level {     id     range {       id       levelCode       __typename     }     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_SendInvitation](ctx, request)
}