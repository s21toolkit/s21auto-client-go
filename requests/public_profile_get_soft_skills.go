package requests

import "github.com/s21toolkit/s21client/gql"

type PublicProfileGetSoftSkills_Variables struct {
	StudentID string `json:"studentId"`
}


type PublicProfileGetSoftSkills_Data struct {
	School21 PublicProfileGetSoftSkills_Data_School21 `json:"school21"`
}

type PublicProfileGetSoftSkills_Data_School21 struct {
	GetSoftSkillsByStudentID     []PublicProfileGetSoftSkills_Data_GetSoftSkillsByStudentID   `json:"getSoftSkillsByStudentId"`
	GetSoftSkillLimitByStudentID PublicProfileGetSoftSkills_Data_GetSoftSkillLimitByStudentID `json:"getSoftSkillLimitByStudentId"`
	Typename                     string                       `json:"__typename"`
}

type PublicProfileGetSoftSkills_Data_GetSoftSkillLimitByStudentID struct {
	PowerLimit int64  `json:"powerLimit"`
	Typename   string `json:"__typename"`
}

type PublicProfileGetSoftSkills_Data_GetSoftSkillsByStudentID struct {
	ID                     string `json:"id"`
	Type                   string `json:"type"`
	Code                   string `json:"code"`
	TotalPower             int64  `json:"totalPower"`
	HueSaturationLightness string `json:"hueSaturationLightness"`
	Typename               string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetSoftSkills(variables PublicProfileGetSoftSkills_Variables) (PublicProfileGetSoftSkills_Data, error) {
	request := gql.NewQueryRequest[PublicProfileGetSoftSkills_Variables](
		"query publicProfileGetSoftSkills($studentId: UUID!) {\n  school21 {\n    getSoftSkillsByStudentId(studentId: $studentId) {\n      id\n      type\n      code\n      totalPower\n      hueSaturationLightness\n      __typename\n    }\n    getSoftSkillLimitByStudentId(studentId: $studentId) {\n      powerLimit\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[PublicProfileGetSoftSkills_Data](ctx, request)
}