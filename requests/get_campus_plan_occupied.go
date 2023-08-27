package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetCampusPlanOccupied struct {
	ClusterID string `json:"clusterId"`
}


type Data_GetCampusPlanOccupied struct {
	Data_Student_GetCampusPlanOccupied Data_Student_GetCampusPlanOccupied `json:"student"`
}

type Data_Student_GetCampusPlanOccupied struct {
	Data_GetClusterPlanStudentsByClusterID_GetCampusPlanOccupied Data_GetClusterPlanStudentsByClusterID_GetCampusPlanOccupied `json:"getClusterPlanStudentsByClusterId"`
	Typename                          string                            `json:"__typename"`
}

type Data_GetClusterPlanStudentsByClusterID_GetCampusPlanOccupied struct {
	OccupiedPlaces []Data_OccupiedPlace_GetCampusPlanOccupied `json:"occupiedPlaces"`
	Typename       string          `json:"__typename"`
}

type Data_OccupiedPlace_GetCampusPlanOccupied struct {
	Row            string     `json:"row"`
	Number         int64      `json:"number"`
	StageGroupName string     `json:"stageGroupName"`
	StageName      string     `json:"stageName"`
	Data_User_GetCampusPlanOccupied           Data_User_GetCampusPlanOccupied       `json:"user"`
	Data_Experience_GetCampusPlanOccupied     Data_Experience_GetCampusPlanOccupied `json:"experience"`
	StudentType    string     `json:"studentType"`
	Typename       string     `json:"__typename"`
}

type Data_Experience_GetCampusPlanOccupied struct {
	ID       string `json:"id"`
	Value    int64  `json:"value"`
	Data_Level_GetCampusPlanOccupied    Data_Level_GetCampusPlanOccupied  `json:"level"`
	Typename string `json:"__typename"`
}

type Data_Level_GetCampusPlanOccupied struct {
	ID       int64  `json:"id"`
	Data_Range_GetCampusPlanOccupied    Data_Range_GetCampusPlanOccupied  `json:"range"`
	Typename string `json:"__typename"`
}

type Data_Range_GetCampusPlanOccupied struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}

type Data_User_GetCampusPlanOccupied struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) GetCampusPlanOccupied(variables Variables_GetCampusPlanOccupied) (Data_GetCampusPlanOccupied, error) {
	request := gql.NewQueryRequest[Variables_GetCampusPlanOccupied](
		"query getCampusPlanOccupied($clusterId: ID!) {\n  student {\n    getClusterPlanStudentsByClusterId(clusterId: $clusterId) {\n      occupiedPlaces {\n        row\n        number\n        stageGroupName\n        stageName\n        user {\n          id\n          login\n          avatarUrl\n          __typename\n        }\n        experience {\n          id\n          value\n          level {\n            id\n            range {\n              id\n              levelCode\n              leftBorder\n              rightBorder\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        studentType\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetCampusPlanOccupied](ctx, request)
}