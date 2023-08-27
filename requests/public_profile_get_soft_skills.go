package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_PublicProfileGetSoftSkills struct {
	StudentID string `json:"studentId"`
}


type Data_PublicProfileGetSoftSkills struct {
	Data_School21_PublicProfileGetSoftSkills Data_School21_PublicProfileGetSoftSkills `json:"school21"`
}

type Data_School21_PublicProfileGetSoftSkills struct {
	Data_GetSoftSkillsByStudentID_PublicProfileGetSoftSkills     []Data_GetSoftSkillsByStudentID_PublicProfileGetSoftSkills   `json:"getSoftSkillsByStudentId"`
	Data_GetSoftSkillLimitByStudentID_PublicProfileGetSoftSkills Data_GetSoftSkillLimitByStudentID_PublicProfileGetSoftSkills `json:"getSoftSkillLimitByStudentId"`
	Typename                     string                       `json:"__typename"`
}

type Data_GetSoftSkillLimitByStudentID_PublicProfileGetSoftSkills struct {
	PowerLimit int64  `json:"powerLimit"`
	Typename   string `json:"__typename"`
}

type Data_GetSoftSkillsByStudentID_PublicProfileGetSoftSkills struct {
	ID                     string `json:"id"`
	Type                   string `json:"type"`
	Code                   string `json:"code"`
	TotalPower             int64  `json:"totalPower"`
	HueSaturationLightness string `json:"hueSaturationLightness"`
	Typename               string `json:"__typename"`
}


func (ctx *RequestContext) PublicProfileGetSoftSkills(variables Variables_PublicProfileGetSoftSkills) (Data_PublicProfileGetSoftSkills, error) {
	request := gql.NewQueryRequest[Variables_PublicProfileGetSoftSkills](
		"query publicProfileGetSoftSkills($studentId: UUID!) {\n  school21 {\n    getSoftSkillsByStudentId(studentId: $studentId) {\n      id\n      type\n      code\n      totalPower\n      hueSaturationLightness\n      __typename\n    }\n    getSoftSkillLimitByStudentId(studentId: $studentId) {\n      powerLimit\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_PublicProfileGetSoftSkills](ctx, request)
}