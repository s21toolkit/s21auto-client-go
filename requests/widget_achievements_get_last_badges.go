package requests

import "s21client/gql"

type Variables_WidgetAchievementsGetLastBadges struct {
	Limit int64 `json:"limit"`
}


type Data_WidgetAchievementsGetLastBadges struct {
	Student_WidgetAchievementsGetLastBadges Student_WidgetAchievementsGetLastBadges `json:"student"`
}

type Student_WidgetAchievementsGetLastBadges struct {
	GetLastBadges []GetLastBadge_WidgetAchievementsGetLastBadges `json:"getLastBadges"`
	Typename      string         `json:"__typename"`
}

type GetLastBadge_WidgetAchievementsGetLastBadges struct {
	ID       string `json:"id"`
	Points   int64  `json:"points"`
	Badge_WidgetAchievementsGetLastBadges    Badge_WidgetAchievementsGetLastBadges  `json:"badge"`
	Award_WidgetAchievementsGetLastBadges    Award_WidgetAchievementsGetLastBadges  `json:"award"`
	Typename string `json:"__typename"`
}

type Award_WidgetAchievementsGetLastBadges struct {
	AwardBounties []AwardBounty_WidgetAchievementsGetLastBadges `json:"awardBounties"`
	Typename      string        `json:"__typename"`
}

type AwardBounty_WidgetAchievementsGetLastBadges struct {
	AwardLevelID *int64 `json:"awardLevelId"`
	Typename     string `json:"__typename"`
}

type Badge_WidgetAchievementsGetLastBadges struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) WidgetAchievementsGetLastBadges(variables Variables_WidgetAchievementsGetLastBadges) (Data_WidgetAchievementsGetLastBadges, error) {
	request := gql.NewQueryRequest[Variables_WidgetAchievementsGetLastBadges](
		"query widgetAchievementsGetLastBadges($limit: Int) {   student {     getLastBadges(limit: $limit) {       id       points       badge {         name         avatarUrl         __typename       }       award {         awardBounties {           awardLevelId           __typename         }         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_WidgetAchievementsGetLastBadges](ctx, request)
}