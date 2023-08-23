package requests

import "s21client/gql"

type Variables_GetCampusPlanOccupied struct {
	ClusterID string `json:"clusterId"`
}


type Data_GetCampusPlanOccupied struct {
	Student_GetCampusPlanOccupied Student_GetCampusPlanOccupied `json:"student"`
}

type Student_GetCampusPlanOccupied struct {
	GetClusterPlanStudentsByClusterID_GetCampusPlanOccupied GetClusterPlanStudentsByClusterID_GetCampusPlanOccupied `json:"getClusterPlanStudentsByClusterId"`
	Typename                          string                            `json:"__typename"`
}

type GetClusterPlanStudentsByClusterID_GetCampusPlanOccupied struct {
	OccupiedPlaces []OccupiedPlace_GetCampusPlanOccupied `json:"occupiedPlaces"`
	Typename       string          `json:"__typename"`
}

type OccupiedPlace_GetCampusPlanOccupied struct {
	Row            string     `json:"row"`
	Number         int64      `json:"number"`
	StageGroupName string     `json:"stageGroupName"`
	StageName      string     `json:"stageName"`
	User_GetCampusPlanOccupied           User_GetCampusPlanOccupied       `json:"user"`
	Experience_GetCampusPlanOccupied     Experience_GetCampusPlanOccupied `json:"experience"`
	StudentType    string     `json:"studentType"`
	Typename       string     `json:"__typename"`
}

type Experience_GetCampusPlanOccupied struct {
	ID       string `json:"id"`
	Value    int64  `json:"value"`
	Level_GetCampusPlanOccupied    Level_GetCampusPlanOccupied  `json:"level"`
	Typename string `json:"__typename"`
}

type Level_GetCampusPlanOccupied struct {
	ID       int64  `json:"id"`
	Range_GetCampusPlanOccupied    Range_GetCampusPlanOccupied  `json:"range"`
	Typename string `json:"__typename"`
}

type Range_GetCampusPlanOccupied struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}

type User_GetCampusPlanOccupied struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) GetCampusPlanOccupied(variables Variables_GetCampusPlanOccupied) (Data_GetCampusPlanOccupied, error) {
	request := gql.NewQueryRequest[Variables_GetCampusPlanOccupied](
		"query getCampusPlanOccupied($clusterId: ID!) {   student {     getClusterPlanStudentsByClusterId(clusterId: $clusterId) {       occupiedPlaces {         row         number         stageGroupName         stageName         user {           id           login           avatarUrl           __typename         }         experience {           id           value           level {             id             range {               id               levelCode               leftBorder               rightBorder               __typename             }             __typename           }           __typename         }         studentType         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetCampusPlanOccupied](ctx, request)
}