package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_CancelInvitation struct {
	UserID string `json:"userId"`
	TeamID string `json:"teamId"`
}

type Response_Data_CancelInvitation struct {
	Student Response_DataStudent_CancelInvitation `json:"student"`
}

type Response_DataStudent_CancelInvitation struct {
	Response_CancelInvitation_CancelInvitation Response_CancelInvitation_CancelInvitation `json:"cancelInvitation"`
	Typename                                   string                                     `json:"__typename"`
}

type Response_CancelInvitation_CancelInvitation struct {
	Student          Response_CancelInvitationStudent_CancelInvitation `json:"student"`
	InvitationStatus string                                            `json:"invitationStatus"`
	Typename         string                                            `json:"__typename"`
}

type Response_CancelInvitationStudent_CancelInvitation struct {
	ID                             string                         `json:"id"`
	Response_User_CancelInvitation Response_User_CancelInvitation `json:"user"`
	Typename                       string                         `json:"__typename"`
}

type Response_User_CancelInvitation struct {
	ID                                       string                                   `json:"id"`
	Login                                    string                                   `json:"login"`
	AvatarURL                                string                                   `json:"avatarUrl"`
	Response_UserExperience_CancelInvitation Response_UserExperience_CancelInvitation `json:"userExperience"`
	Typename                                 string                                   `json:"__typename"`
}

type Response_UserExperience_CancelInvitation struct {
	ID                              string                          `json:"id"`
	CookiesCount                    int64                           `json:"cookiesCount"`
	CodeReviewPoints                int64                           `json:"codeReviewPoints"`
	CoinsCount                      int64                           `json:"coinsCount"`
	Response_Level_CancelInvitation Response_Level_CancelInvitation `json:"level"`
	Typename                        string                          `json:"__typename"`
}

type Response_Level_CancelInvitation struct {
	ID                              int64                           `json:"id"`
	Response_Range_CancelInvitation Response_Range_CancelInvitation `json:"range"`
	Typename                        string                          `json:"__typename"`
}

type Response_Range_CancelInvitation struct {
	ID        string `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) CancelInvitation(variables Request_Variables_CancelInvitation) (Response_Data_CancelInvitation, error) {
	request := gql.NewQueryRequest[Request_Variables_CancelInvitation](
		"mutation cancelInvitation($teamId: UUID!, $userId: ID!) {   student {     cancelInvitation(teamId: $teamId, userId: $userId) {       ...StudentInvitationInfo       __typename     }     __typename   } }  fragment StudentInvitationInfo on StudentInvitationInfo {   student {     ...AvailableStudentForTeam     __typename   }   invitationStatus   __typename }  fragment AvailableStudentForTeam on Student {   id   user {     id     login     avatarUrl     userExperience {       ...CurrentUserExperience       __typename     }     __typename   }   __typename }  fragment CurrentUserExperience on UserExperience {   id   cookiesCount   codeReviewPoints   coinsCount   level {     id     range {       id       levelCode       __typename     }     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_CancelInvitation](ctx, request)
}
