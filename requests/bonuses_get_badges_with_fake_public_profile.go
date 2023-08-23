package requests

import "s21client/gql"

type Variables_BonusesGetBadgesWithFakePublicProfile struct {
	UserID *string `json:"userId"`
}


type Data_BonusesGetBadgesWithFakePublicProfile struct {
	Student_BonusesGetBadgesWithFakePublicProfile Student_BonusesGetBadgesWithFakePublicProfile `json:"student"`
}

type Student_BonusesGetBadgesWithFakePublicProfile struct {
	GetBadgesWithFakePublicProfile_BonusesGetBadgesWithFakePublicProfile []GetBadgesWithFakePublicProfile_BonusesGetBadgesWithFakePublicProfile `json:"getBadgesWithFakePublicProfile"`
	Typename                       string                           `json:"__typename"`
}

type GetBadgesWithFakePublicProfile_BonusesGetBadgesWithFakePublicProfile struct {
	ID        string    `json:"id"`
	Histories []History_BonusesGetBadgesWithFakePublicProfile `json:"histories"`
	Badge_BonusesGetBadgesWithFakePublicProfile     Badge_BonusesGetBadgesWithFakePublicProfile     `json:"badge"`
	Award_BonusesGetBadgesWithFakePublicProfile     Award_BonusesGetBadgesWithFakePublicProfile     `json:"award"`
	Points    int64     `json:"points"`
	IsFake    bool      `json:"isFake"`
	Typename  string    `json:"__typename"`
}

type Award_BonusesGetBadgesWithFakePublicProfile struct {
	ID             string         `json:"id"`
	AwardCondition_BonusesGetBadgesWithFakePublicProfile AwardCondition_BonusesGetBadgesWithFakePublicProfile `json:"awardCondition"`
	AwardBounties  []AwardBounty_BonusesGetBadgesWithFakePublicProfile  `json:"awardBounties"`
	Typename       string         `json:"__typename"`
}

type AwardBounty_BonusesGetBadgesWithFakePublicProfile struct {
	AwardBountyID   string           `json:"awardBountyId"`
	Description     string           `json:"description"`
	Cookies         int64            `json:"cookies"`
	Coins           int64            `json:"coins"`
	ExperienceValue int64            `json:"experienceValue"`
	CoalitionPoints int64            `json:"coalitionPoints"`
	SoftSkillPowers []SoftSkillPower_BonusesGetBadgesWithFakePublicProfile `json:"softSkillPowers"`
	Typename        string           `json:"__typename"`
}

type SoftSkillPower_BonusesGetBadgesWithFakePublicProfile struct {
	SoftSkillID int64     `json:"softSkillId"`
	Power       int64     `json:"power"`
	SoftSkill_BonusesGetBadgesWithFakePublicProfile   SoftSkill_BonusesGetBadgesWithFakePublicProfile `json:"softSkill"`
	Typename    string    `json:"__typename"`
}

type SoftSkill_BonusesGetBadgesWithFakePublicProfile struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}

type AwardCondition_BonusesGetBadgesWithFakePublicProfile struct {
	ID          int64  `json:"id"`
	Description string `json:"description"`
	Typename    string `json:"__typename"`
}

type Badge_BonusesGetBadgesWithFakePublicProfile struct {
	ID           string `json:"id"`
	Kind_BonusesGetBadgesWithFakePublicProfile         Kind_BonusesGetBadgesWithFakePublicProfile   `json:"kind"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	AvatarURL    string `json:"avatarUrl"`
	BigAvatarURL string `json:"bigAvatarUrl"`
	Typename     string `json:"__typename"`
}

type Kind_BonusesGetBadgesWithFakePublicProfile struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Order    int64  `json:"order"`
	Typename string `json:"__typename"`
}

type History_BonusesGetBadgesWithFakePublicProfile struct {
	ID          string `json:"id"`
	RewardDate  string `json:"rewardDate"`
	AwardPoints int64  `json:"awardPoints"`
	Typename    string `json:"__typename"`
}

func (ctx *RequestContext) BonusesGetBadgesWithFakePublicProfile(variables Variables_BonusesGetBadgesWithFakePublicProfile) (Data_BonusesGetBadgesWithFakePublicProfile, error) {
	request := gql.NewQueryRequest[Variables_BonusesGetBadgesWithFakePublicProfile](
		"query bonusesGetBadgesWithFakePublicProfile($userId: UUID) {   student {     getBadgesWithFakePublicProfile(userId: $userId) {       ...UserAchievements       __typename     }     __typename   } }  fragment UserAchievements on UserBadgeAward {   id   histories {     id     rewardDate     awardPoints     __typename   }   badge {     id     kind {       id       name       order       __typename     }     name     description     avatarUrl     bigAvatarUrl     __typename   }   award {     id     awardCondition {       id       description       __typename     }     awardBounties {       awardBountyId       description       cookies       coins       experienceValue       coalitionPoints       softSkillPowers {         softSkillId         power         softSkill {           id           name           __typename         }         __typename       }       __typename     }     __typename   }   points   isFake   __typename } ",
		variables,
	)

	return GqlRequest[Data_BonusesGetBadgesWithFakePublicProfile](ctx, request)
}