package requests

import "github.com/s21toolkit/s21client/gql"

type GetCampusPlanOccupied_Variables struct {
	ClusterID string `json:"clusterId"`
}


type GetCampusPlanOccupied_Data struct {
	Student GetCampusPlanOccupied_Data_Student `json:"student"`
}

type GetCampusPlanOccupied_Data_Student struct {
	GetClusterPlanStudentsByClusterID GetCampusPlanOccupied_Data_GetClusterPlanStudentsByClusterID `json:"getClusterPlanStudentsByClusterId"`
	Typename                          string                            `json:"__typename"`
}

type GetCampusPlanOccupied_Data_GetClusterPlanStudentsByClusterID struct {
	OccupiedPlaces []GetCampusPlanOccupied_Data_OccupiedPlace `json:"occupiedPlaces"`
	Typename       string          `json:"__typename"`
}

type GetCampusPlanOccupied_Data_OccupiedPlace struct {
	Row            string     `json:"row"`
	Number         int64      `json:"number"`
	StageGroupName string     `json:"stageGroupName"`
	StageName      string     `json:"stageName"`
	User           GetCampusPlanOccupied_Data_User       `json:"user"`
	Experience     GetCampusPlanOccupied_Data_Experience `json:"experience"`
	StudentType    string     `json:"studentType"`
	Typename       string     `json:"__typename"`
}

type GetCampusPlanOccupied_Data_Experience struct {
	ID       string `json:"id"`
	Value    int64  `json:"value"`
	Level    GetCampusPlanOccupied_Data_Level  `json:"level"`
	Typename string `json:"__typename"`
}

type GetCampusPlanOccupied_Data_Level struct {
	ID       int64  `json:"id"`
	Range    GetCampusPlanOccupied_Data_Range  `json:"range"`
	Typename string `json:"__typename"`
}

type GetCampusPlanOccupied_Data_Range struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}

type GetCampusPlanOccupied_Data_User struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) GetCampusPlanOccupied(variables GetCampusPlanOccupied_Variables) (GetCampusPlanOccupied_Data, error) {
	request := gql.NewQueryRequest[GetCampusPlanOccupied_Variables](
		"query getCampusPlanOccupied($clusterId: ID!) {\n  student {\n    getClusterPlanStudentsByClusterId(clusterId: $clusterId) {\n      occupiedPlaces {\n        row\n        number\n        stageGroupName\n        stageName\n        user {\n          id\n          login\n          avatarUrl\n          __typename\n        }\n        experience {\n          id\n          value\n          level {\n            id\n            range {\n              id\n              levelCode\n              leftBorder\n              rightBorder\n              __typename\n            }\n            __typename\n          }\n          __typename\n        }\n        studentType\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetCampusPlanOccupied_Data](ctx, request)
}