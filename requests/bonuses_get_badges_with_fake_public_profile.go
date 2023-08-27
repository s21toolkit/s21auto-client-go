package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_BonusesGetBadgesWithFakePublicProfile struct {
	UserID *string `json:"userId"`
}


type Data_BonusesGetBadgesWithFakePublicProfile struct {
	Data_Student_BonusesGetBadgesWithFakePublicProfile Data_Student_BonusesGetBadgesWithFakePublicProfile `json:"student"`
}

type Data_Student_BonusesGetBadgesWithFakePublicProfile struct {
	Data_GetBadgesWithFakePublicProfile_BonusesGetBadgesWithFakePublicProfile []Data_GetBadgesWithFakePublicProfile_BonusesGetBadgesWithFakePublicProfile `json:"getBadgesWithFakePublicProfile"`
	Typename                       string                           `json:"__typename"`
}

type Data_GetBadgesWithFakePublicProfile_BonusesGetBadgesWithFakePublicProfile struct {
	ID        string    `json:"id"`
	Histories []Data_History_BonusesGetBadgesWithFakePublicProfile `json:"histories"`
	Data_Badge_BonusesGetBadgesWithFakePublicProfile     Data_Badge_BonusesGetBadgesWithFakePublicProfile     `json:"badge"`
	Data_Award_BonusesGetBadgesWithFakePublicProfile     Data_Award_BonusesGetBadgesWithFakePublicProfile     `json:"award"`
	Points    int64     `json:"points"`
	IsFake    bool      `json:"isFake"`
	Typename  string    `json:"__typename"`
}

type Data_Award_BonusesGetBadgesWithFakePublicProfile struct {
	ID             string         `json:"id"`
	Data_AwardCondition_BonusesGetBadgesWithFakePublicProfile Data_AwardCondition_BonusesGetBadgesWithFakePublicProfile `json:"awardCondition"`
	AwardBounties  []Data_AwardBounty_BonusesGetBadgesWithFakePublicProfile  `json:"awardBounties"`
	Typename       string         `json:"__typename"`
}

type Data_AwardBounty_BonusesGetBadgesWithFakePublicProfile struct {
	AwardBountyID   string           `json:"awardBountyId"`
	Description     string           `json:"description"`
	Cookies         int64            `json:"cookies"`
	Coins           int64            `json:"coins"`
	ExperienceValue int64            `json:"experienceValue"`
	CoalitionPoints int64            `json:"coalitionPoints"`
	SoftSkillPowers []Data_SoftSkillPower_BonusesGetBadgesWithFakePublicProfile `json:"softSkillPowers"`
	Typename        string           `json:"__typename"`
}

type Data_SoftSkillPower_BonusesGetBadgesWithFakePublicProfile struct {
	SoftSkillID int64     `json:"softSkillId"`
	Power       int64     `json:"power"`
	Data_SoftSkill_BonusesGetBadgesWithFakePublicProfile   Data_SoftSkill_BonusesGetBadgesWithFakePublicProfile `json:"softSkill"`
	Typename    string    `json:"__typename"`
}

type Data_SoftSkill_BonusesGetBadgesWithFakePublicProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

type Data_AwardCondition_BonusesGetBadgesWithFakePublicProfile struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Typename    string `json:"__typename"`
}

type Data_Badge_BonusesGetBadgesWithFakePublicProfile struct {
	ID           string `json:"id"`
	Data_Kind_BonusesGetBadgesWithFakePublicProfile         Data_Kind_BonusesGetBadgesWithFakePublicProfile   `json:"kind"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	AvatarURL    string `json:"avatarUrl"`
	BigAvatarURL string `json:"bigAvatarUrl"`
	Typename     string `json:"__typename"`
}

type Data_Kind_BonusesGetBadgesWithFakePublicProfile struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Order    int64  `json:"order"`
	Typename string `json:"__typename"`
}

type Data_History_BonusesGetBadgesWithFakePublicProfile struct {
	ID          string `json:"id"`
	RewardDate  string `json:"rewardDate"`
	AwardPoints int64  `json:"awardPoints"`
	Typename    string `json:"__typename"`
}


func (ctx *RequestContext) BonusesGetBadgesWithFakePublicProfile(variables Variables_BonusesGetBadgesWithFakePublicProfile) (Data_BonusesGetBadgesWithFakePublicProfile, error) {
	request := gql.NewQueryRequest[Variables_BonusesGetBadgesWithFakePublicProfile](
		"query bonusesGetBadgesWithFakePublicProfile($userId: UUID) {\n  student {\n    getBadgesWithFakePublicProfile(userId: $userId) {\n      ...UserAchievements\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment UserAchievements on UserBadgeAward {\n  id\n  histories {\n    id\n    rewardDate\n    awardPoints\n    __typename\n  }\n  badge {\n    id\n    kind {\n      id\n      name\n      order\n      __typename\n    }\n    name\n    description\n    avatarUrl\n    bigAvatarUrl\n    __typename\n  }\n  award {\n    id\n    awardCondition {\n      id\n      description\n      __typename\n    }\n    awardBounties {\n      awardBountyId\n      description\n      cookies\n      coins\n      experienceValue\n      coalitionPoints\n      softSkillPowers {\n        softSkillId\n        power\n        softSkill {\n          id\n          name\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  points\n  isFake\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Data_BonusesGetBadgesWithFakePublicProfile](ctx, request)
}