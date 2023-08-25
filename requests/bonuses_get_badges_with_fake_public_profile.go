package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_BonusesGetBadgesWithFakePublicProfile struct {
	UserID *string `json:"userId"`
}


type Response_Data_BonusesGetBadgesWithFakePublicProfile struct {
	Response_Student_BonusesGetBadgesWithFakePublicProfile Response_Student_BonusesGetBadgesWithFakePublicProfile `json:"student"`
}

type Response_Student_BonusesGetBadgesWithFakePublicProfile struct {
	Response_GetBadgesWithFakePublicProfile_BonusesGetBadgesWithFakePublicProfile []Response_GetBadgesWithFakePublicProfile_BonusesGetBadgesWithFakePublicProfile `json:"getBadgesWithFakePublicProfile"`
	Typename                       string                           `json:"__typename"`
}

type Response_GetBadgesWithFakePublicProfile_BonusesGetBadgesWithFakePublicProfile struct {
	ID        string    `json:"id"`
	Histories []Response_History_BonusesGetBadgesWithFakePublicProfile `json:"histories"`
	Response_Badge_BonusesGetBadgesWithFakePublicProfile     Response_Badge_BonusesGetBadgesWithFakePublicProfile     `json:"badge"`
	Response_Award_BonusesGetBadgesWithFakePublicProfile     Response_Award_BonusesGetBadgesWithFakePublicProfile     `json:"award"`
	Points    int64     `json:"points"`
	IsFake    bool      `json:"isFake"`
	Typename  string    `json:"__typename"`
}

type Response_Award_BonusesGetBadgesWithFakePublicProfile struct {
	ID             string         `json:"id"`
	Response_AwardCondition_BonusesGetBadgesWithFakePublicProfile Response_AwardCondition_BonusesGetBadgesWithFakePublicProfile `json:"awardCondition"`
	AwardBounties  []Response_AwardBounty_BonusesGetBadgesWithFakePublicProfile  `json:"awardBounties"`
	Typename       string         `json:"__typename"`
}

type Response_AwardBounty_BonusesGetBadgesWithFakePublicProfile struct {
	AwardBountyID   string           `json:"awardBountyId"`
	Description     string           `json:"description"`
	Cookies         int64            `json:"cookies"`
	Coins           int64            `json:"coins"`
	ExperienceValue int64            `json:"experienceValue"`
	CoalitionPoints int64            `json:"coalitionPoints"`
	SoftSkillPowers []Response_SoftSkillPower_BonusesGetBadgesWithFakePublicProfile `json:"softSkillPowers"`
	Typename        string           `json:"__typename"`
}

type Response_SoftSkillPower_BonusesGetBadgesWithFakePublicProfile struct {
	SoftSkillID int64     `json:"softSkillId"`
	Power       int64     `json:"power"`
	Response_SoftSkill_BonusesGetBadgesWithFakePublicProfile   Response_SoftSkill_BonusesGetBadgesWithFakePublicProfile `json:"softSkill"`
	Typename    string    `json:"__typename"`
}

type Response_SoftSkill_BonusesGetBadgesWithFakePublicProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

type Response_AwardCondition_BonusesGetBadgesWithFakePublicProfile struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Typename    string `json:"__typename"`
}

type Response_Badge_BonusesGetBadgesWithFakePublicProfile struct {
	ID           string `json:"id"`
	Response_Kind_BonusesGetBadgesWithFakePublicProfile         Response_Kind_BonusesGetBadgesWithFakePublicProfile   `json:"kind"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	AvatarURL    string `json:"avatarUrl"`
	BigAvatarURL string `json:"bigAvatarUrl"`
	Typename     string `json:"__typename"`
}

type Response_Kind_BonusesGetBadgesWithFakePublicProfile struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Order    int64  `json:"order"`
	Typename string `json:"__typename"`
}

type Response_History_BonusesGetBadgesWithFakePublicProfile struct {
	ID          string `json:"id"`
	RewardDate  string `json:"rewardDate"`
	AwardPoints int64  `json:"awardPoints"`
	Typename    string `json:"__typename"`
}

func (ctx *RequestContext) BonusesGetBadgesWithFakePublicProfile(variables Request_Variables_BonusesGetBadgesWithFakePublicProfile) (Response_Data_BonusesGetBadgesWithFakePublicProfile, error) {
	request := gql.NewQueryRequest[Request_Variables_BonusesGetBadgesWithFakePublicProfile](
		"query bonusesGetBadgesWithFakePublicProfile($userId: UUID) {\n  student {\n    getBadgesWithFakePublicProfile(userId: $userId) {\n      ...UserAchievements\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment UserAchievements on UserBadgeAward {\n  id\n  histories {\n    id\n    rewardDate\n    awardPoints\n    __typename\n  }\n  badge {\n    id\n    kind {\n      id\n      name\n      order\n      __typename\n    }\n    name\n    description\n    avatarUrl\n    bigAvatarUrl\n    __typename\n  }\n  award {\n    id\n    awardCondition {\n      id\n      description\n      __typename\n    }\n    awardBounties {\n      awardBountyId\n      description\n      cookies\n      coins\n      experienceValue\n      coalitionPoints\n      softSkillPowers {\n        softSkillId\n        power\n        softSkill {\n          id\n          name\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  points\n  isFake\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[Response_Data_BonusesGetBadgesWithFakePublicProfile](ctx, request)
}