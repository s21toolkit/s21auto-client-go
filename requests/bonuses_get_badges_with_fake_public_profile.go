package requests

import "github.com/s21toolkit/s21client/gql"

type BonusesGetBadgesWithFakePublicProfile_Variables struct {
	UserID *string `json:"userId"`
}


type BonusesGetBadgesWithFakePublicProfile_Data struct {
	Student BonusesGetBadgesWithFakePublicProfile_Data_Student `json:"student"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_Student struct {
	GetBadgesWithFakePublicProfile []BonusesGetBadgesWithFakePublicProfile_Data_GetBadgesWithFakePublicProfile `json:"getBadgesWithFakePublicProfile"`
	Typename                       string                           `json:"__typename"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_GetBadgesWithFakePublicProfile struct {
	ID        string    `json:"id"`
	Histories []BonusesGetBadgesWithFakePublicProfile_Data_History `json:"histories"`
	Badge     BonusesGetBadgesWithFakePublicProfile_Data_Badge     `json:"badge"`
	Award     BonusesGetBadgesWithFakePublicProfile_Data_Award     `json:"award"`
	Points    int64     `json:"points"`
	IsFake    bool      `json:"isFake"`
	Typename  string    `json:"__typename"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_Award struct {
	ID             string         `json:"id"`
	AwardCondition BonusesGetBadgesWithFakePublicProfile_Data_AwardCondition `json:"awardCondition"`
	AwardBounties  []BonusesGetBadgesWithFakePublicProfile_Data_AwardBounty  `json:"awardBounties"`
	Typename       string         `json:"__typename"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_AwardBounty struct {
	AwardBountyID   string           `json:"awardBountyId"`
	Description     string           `json:"description"`
	Cookies         int64            `json:"cookies"`
	Coins           int64            `json:"coins"`
	ExperienceValue int64            `json:"experienceValue"`
	CoalitionPoints int64            `json:"coalitionPoints"`
	SoftSkillPowers []BonusesGetBadgesWithFakePublicProfile_Data_SoftSkillPower `json:"softSkillPowers"`
	Typename        string           `json:"__typename"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_SoftSkillPower struct {
	SoftSkillID int64     `json:"softSkillId"`
	Power       int64     `json:"power"`
	SoftSkill   BonusesGetBadgesWithFakePublicProfile_Data_SoftSkill `json:"softSkill"`
	Typename    string    `json:"__typename"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_SoftSkill struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_AwardCondition struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Typename    string `json:"__typename"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_Badge struct {
	ID           string `json:"id"`
	Kind         BonusesGetBadgesWithFakePublicProfile_Data_Kind   `json:"kind"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	AvatarURL    string `json:"avatarUrl"`
	BigAvatarURL string `json:"bigAvatarUrl"`
	Typename     string `json:"__typename"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_Kind struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Order    int64  `json:"order"`
	Typename string `json:"__typename"`
}

type BonusesGetBadgesWithFakePublicProfile_Data_History struct {
	ID          string `json:"id"`
	RewardDate  string `json:"rewardDate"`
	AwardPoints int64  `json:"awardPoints"`
	Typename    string `json:"__typename"`
}


func (ctx *RequestContext) BonusesGetBadgesWithFakePublicProfile(variables BonusesGetBadgesWithFakePublicProfile_Variables) (BonusesGetBadgesWithFakePublicProfile_Data, error) {
	request := gql.NewQueryRequest[BonusesGetBadgesWithFakePublicProfile_Variables](
		"query bonusesGetBadgesWithFakePublicProfile($userId: UUID) {\n  student {\n    getBadgesWithFakePublicProfile(userId: $userId) {\n      ...UserAchievements\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment UserAchievements on UserBadgeAward {\n  id\n  histories {\n    id\n    rewardDate\n    awardPoints\n    __typename\n  }\n  badge {\n    id\n    kind {\n      id\n      name\n      order\n      __typename\n    }\n    name\n    description\n    avatarUrl\n    bigAvatarUrl\n    __typename\n  }\n  award {\n    id\n    awardCondition {\n      id\n      description\n      __typename\n    }\n    awardBounties {\n      awardBountyId\n      description\n      cookies\n      coins\n      experienceValue\n      coalitionPoints\n      softSkillPowers {\n        softSkillId\n        power\n        softSkill {\n          id\n          name\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  points\n  isFake\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[BonusesGetBadgesWithFakePublicProfile_Data](ctx, request)
}