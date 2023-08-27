package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_PublicProfileGetAchievements struct {
	UserID string `json:"userId"`
}


type Data_PublicProfileGetAchievements struct {
	Data_Student_PublicProfileGetAchievements Data_Student_PublicProfileGetAchievements `json:"student"`
}

type Data_Student_PublicProfileGetAchievements struct {
	Data_GetBadgesPublicProfile_PublicProfileGetAchievements []Data_GetBadgesPublicProfile_PublicProfileGetAchievements `json:"getBadgesPublicProfile"`
	Typename               string                   `json:"__typename"`
}

type Data_GetBadgesPublicProfile_PublicProfileGetAchievements struct {
	Points   int64  `json:"points"`
	ID       string `json:"id"`
	Data_Badge_PublicProfileGetAchievements    Data_Badge_PublicProfileGetAchievements  `json:"badge"`
	Typename string `json:"__typename"`
}

type Data_Badge_PublicProfileGetAchievements struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetAchievements(variables Variables_PublicProfileGetAchievements) (Data_PublicProfileGetAchievements, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetAchievements](
		"query publicProfileGetAchievements($userId: UUID!) {\n  student {\n    getBadgesPublicProfile(userId: $userId) {\n      points\n      id\n      badge {\n        id\n        name\n        avatarUrl\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetAchievements](ctx, request)
}