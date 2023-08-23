package requests

import "s21client/gql"

type Variables_PublicProfileGetSoftSkills struct {
	StudentID string `json:"studentId"`
}


type Data_PublicProfileGetSoftSkills struct {
	School21_PublicProfileGetSoftSkills School21_PublicProfileGetSoftSkills `json:"school21"`
}

type School21_PublicProfileGetSoftSkills struct {
	GetSoftSkillsByStudentID_PublicProfileGetSoftSkills     []GetSoftSkillsByStudentID_PublicProfileGetSoftSkills   `json:"getSoftSkillsByStudentId"`
	GetSoftSkillLimitByStudentID_PublicProfileGetSoftSkills GetSoftSkillLimitByStudentID_PublicProfileGetSoftSkills `json:"getSoftSkillLimitByStudentId"`
	Typename                     string                       `json:"__typename"`
}

type GetSoftSkillLimitByStudentID_PublicProfileGetSoftSkills struct {
	PowerLimit int64  `json:"powerLimit"`
	Typename   string `json:"__typename"`
}

type GetSoftSkillsByStudentID_PublicProfileGetSoftSkills struct {
	ID                     string `json:"id"`
	Type                   string `json:"type"`
	Code                   string `json:"code"`
	TotalPower             int64  `json:"totalPower"`
	HueSaturationLightness string `json:"hueSaturationLightness"`
	Typename               string `json:"__typename"`
}

func (ctx *RequestContext) PublicProfileGetSoftSkills(variables Variables_PublicProfileGetSoftSkills) (Data_PublicProfileGetSoftSkills, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetSoftSkills](
		"query publicProfileGetSoftSkills($studentId: UUID!) {   school21 {     getSoftSkillsByStudentId(studentId: $studentId) {       id       type       code       totalPower       hueSaturationLightness       __typename     }     getSoftSkillLimitByStudentId(studentId: $studentId) {       powerLimit       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetSoftSkills](ctx, request)
}