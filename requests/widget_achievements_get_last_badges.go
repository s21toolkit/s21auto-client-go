package requests

import "github.com/s21toolkit/s21client/gql"

type WidgetAchievementsGetLastBadges_Variables struct {
	Limit int64 `json:"limit"`
}


type WidgetAchievementsGetLastBadges_Data struct {
	Student WidgetAchievementsGetLastBadges_Data_Student `json:"student"`
}

type WidgetAchievementsGetLastBadges_Data_Student struct {
	GetLastBadges []WidgetAchievementsGetLastBadges_Data_GetLastBadge `json:"getLastBadges"`
	Typename      string         `json:"__typename"`
}

type WidgetAchievementsGetLastBadges_Data_GetLastBadge struct {
	ID       string `json:"id"`
	Points   int64  `json:"points"`
	Badge    WidgetAchievementsGetLastBadges_Data_Badge  `json:"badge"`
	Award    WidgetAchievementsGetLastBadges_Data_Award  `json:"award"`
	Typename string `json:"__typename"`
}

type WidgetAchievementsGetLastBadges_Data_Award struct {
	AwardBounties []WidgetAchievementsGetLastBadges_Data_AwardBounty `json:"awardBounties"`
	Typename      string        `json:"__typename"`
}

type WidgetAchievementsGetLastBadges_Data_AwardBounty struct {
	AwardLevelID *int64 `json:"awardLevelId"`
	Typename     string `json:"__typename"`
}

type WidgetAchievementsGetLastBadges_Data_Badge struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) WidgetAchievementsGetLastBadges(variables WidgetAchievementsGetLastBadges_Variables) (WidgetAchievementsGetLastBadges_Data, error) {
	request := gql.NewQueryRequest[WidgetAchievementsGetLastBadges_Variables](
		"query widgetAchievementsGetLastBadges($limit: Int) {\n  student {\n    getLastBadges(limit: $limit) {\n      id\n      points\n      badge {\n        name\n        avatarUrl\n        __typename\n      }\n      award {\n        awardBounties {\n          awardLevelId\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[WidgetAchievementsGetLastBadges_Data](ctx, request)
}