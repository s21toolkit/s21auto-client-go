package requests

import "s21client/gql"

type Request_Variables_WidgetAchievementsGetLastBadges struct {
	Limit int64 `json:"limit"`
}


type Response_Data_WidgetAchievementsGetLastBadges struct {
	Response_Student_WidgetAchievementsGetLastBadges Response_Student_WidgetAchievementsGetLastBadges `json:"student"`
}

type Response_Student_WidgetAchievementsGetLastBadges struct {
	GetLastBadges []Response_GetLastBadge_WidgetAchievementsGetLastBadges `json:"getLastBadges"`
	Typename      string         `json:"__typename"`
}

type Response_GetLastBadge_WidgetAchievementsGetLastBadges struct {
	ID       string `json:"id"`
	Points   int64  `json:"points"`
	Response_Badge_WidgetAchievementsGetLastBadges    Response_Badge_WidgetAchievementsGetLastBadges  `json:"badge"`
	Response_Award_WidgetAchievementsGetLastBadges    Response_Award_WidgetAchievementsGetLastBadges  `json:"award"`
	Typename string `json:"__typename"`
}

type Response_Award_WidgetAchievementsGetLastBadges struct {
	AwardBounties []Response_AwardBounty_WidgetAchievementsGetLastBadges `json:"awardBounties"`
	Typename      string        `json:"__typename"`
}

type Response_AwardBounty_WidgetAchievementsGetLastBadges struct {
	AwardLevelID *int64 `json:"awardLevelId"`
	Typename     string `json:"__typename"`
}

type Response_Badge_WidgetAchievementsGetLastBadges struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) WidgetAchievementsGetLastBadges(variables Request_Variables_WidgetAchievementsGetLastBadges) (Response_Data_WidgetAchievementsGetLastBadges, error) {
	request := gql.NewQueryRequest[Request_Variables_WidgetAchievementsGetLastBadges](
		"query widgetAchievementsGetLastBadges($limit: Int) {   student {     getLastBadges(limit: $limit) {       id       points       badge {         name         avatarUrl         __typename       }       award {         awardBounties {           awardLevelId           __typename         }         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_WidgetAchievementsGetLastBadges](ctx, request)
}