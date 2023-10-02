package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileGetAchievements_Variables struct {
	UserID string `json:"userId"`
}


type PublicProfileGetAchievements_Data struct {
	Student PublicProfileGetAchievements_Data_Student `json:"student"`
}

type PublicProfileGetAchievements_Data_Student struct {
	GetBadgesPublicProfile []PublicProfileGetAchievements_Data_GetBadgesPublicProfile `json:"getBadgesPublicProfile"`
	Typename               string                   `json:"__typename"`
}

type PublicProfileGetAchievements_Data_GetBadgesPublicProfile struct {
	Points   int64  `json:"points"`
	ID       string `json:"id"`
	Badge    PublicProfileGetAchievements_Data_Badge  `json:"badge"`
	Typename string `json:"__typename"`
}

type PublicProfileGetAchievements_Data_Badge struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetAchievements(variables PublicProfileGetAchievements_Variables) (PublicProfileGetAchievements_Data, error) {
	request := gql.NewQueryRequest[PublicProfileGetAchievements_Variables](
		"query publicProfileGetAchievements($userId: UUID!) {\n  student {\n    getBadgesPublicProfile(userId: $userId) {\n      points\n      id\n      badge {\n        id\n        name\n        avatarUrl\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileGetAchievements_Data](ctx, request)
}