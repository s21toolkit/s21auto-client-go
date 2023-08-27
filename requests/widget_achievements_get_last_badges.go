package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_WidgetAchievementsGetLastBadges struct {
	Limit int64 `json:"limit"`
}


type Data_WidgetAchievementsGetLastBadges struct {
	Data_Student_WidgetAchievementsGetLastBadges Data_Student_WidgetAchievementsGetLastBadges `json:"student"`
}

type Data_Student_WidgetAchievementsGetLastBadges struct {
	GetLastBadges []Data_GetLastBadge_WidgetAchievementsGetLastBadges `json:"getLastBadges"`
	Typename      string         `json:"__typename"`
}

type Data_GetLastBadge_WidgetAchievementsGetLastBadges struct {
	ID       string `json:"id"`
	Points   int64  `json:"points"`
	Data_Badge_WidgetAchievementsGetLastBadges    Data_Badge_WidgetAchievementsGetLastBadges  `json:"badge"`
	Data_Award_WidgetAchievementsGetLastBadges    Data_Award_WidgetAchievementsGetLastBadges  `json:"award"`
	Typename string `json:"__typename"`
}

type Data_Award_WidgetAchievementsGetLastBadges struct {
	AwardBounties []Data_AwardBounty_WidgetAchievementsGetLastBadges `json:"awardBounties"`
	Typename      string        `json:"__typename"`
}

type Data_AwardBounty_WidgetAchievementsGetLastBadges struct {
	AwardLevelID *int64 `json:"awardLevelId"`
	Typename     string `json:"__typename"`
}

type Data_Badge_WidgetAchievementsGetLastBadges struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) WidgetAchievementsGetLastBadges(variables Variables_WidgetAchievementsGetLastBadges) (Data_WidgetAchievementsGetLastBadges, error) {
	request := gql.NewQueryRequest[Variables_WidgetAchievementsGetLastBadges](
		"query widgetAchievementsGetLastBadges($limit: Int) {\n  student {\n    getLastBadges(limit: $limit) {\n      id\n      points\n      badge {\n        name\n        avatarUrl\n        __typename\n      }\n      award {\n        awardBounties {\n          awardLevelId\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_WidgetAchievementsGetLastBadges](ctx, request)
}