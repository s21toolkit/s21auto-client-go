package requests

import "s21client/gql"

type Request_Variables_PublicProfileGetAchievements struct {
	UserID string `json:"userId"`
}


type Response_Data_PublicProfileGetAchievements struct {
	Response_Student_PublicProfileGetAchievements Response_Student_PublicProfileGetAchievements `json:"student"`
}

type Response_Student_PublicProfileGetAchievements struct {
	Response_GetBadgesPublicProfile_PublicProfileGetAchievements []Response_GetBadgesPublicProfile_PublicProfileGetAchievements `json:"getBadgesPublicProfile"`
	Typename               string                   `json:"__typename"`
}

type Response_GetBadgesPublicProfile_PublicProfileGetAchievements struct {
	Points   int64  `json:"points"`
	ID       string `json:"id"`
	Response_Badge_PublicProfileGetAchievements    Response_Badge_PublicProfileGetAchievements  `json:"badge"`
	Typename string `json:"__typename"`
}

type Response_Badge_PublicProfileGetAchievements struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetAchievements(variables Request_Variables_PublicProfileGetAchievements) (Response_Data_PublicProfileGetAchievements, error) {
	request := gql.NewQueryRequest[Request_Variables_PublicProfileGetAchievements](
		"query publicProfileGetAchievements($userId: UUID!) {   student {     getBadgesPublicProfile(userId: $userId) {       points       id       badge {         id         name         avatarUrl         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_PublicProfileGetAchievements](ctx, request)
}