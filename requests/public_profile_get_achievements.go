package requests

import "s21client/gql"

type Variables_PublicProfileGetAchievements struct {
	UserID string `json:"userId"`
}


type Data_PublicProfileGetAchievements struct {
	Student_PublicProfileGetAchievements Student_PublicProfileGetAchievements `json:"student"`
}

type Student_PublicProfileGetAchievements struct {
	GetBadgesPublicProfile_PublicProfileGetAchievements []GetBadgesPublicProfile_PublicProfileGetAchievements `json:"getBadgesPublicProfile"`
	Typename               string                   `json:"__typename"`
}

type GetBadgesPublicProfile_PublicProfileGetAchievements struct {
	Points   int64  `json:"points"`
	ID       string `json:"id"`
	Badge_PublicProfileGetAchievements    Badge_PublicProfileGetAchievements  `json:"badge"`
	Typename string `json:"__typename"`
}

type Badge_PublicProfileGetAchievements struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetAchievements(variables Variables_PublicProfileGetAchievements) (Data_PublicProfileGetAchievements, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetAchievements](
		"query publicProfileGetAchievements($userId: UUID!) {   student {     getBadgesPublicProfile(userId: $userId) {       points       id       badge {         id         name         avatarUrl         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetAchievements](ctx, request)
}