package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetCampusPlanOccupied struct {
	ClusterID string `json:"clusterId"`
}


type Response_Data_GetCampusPlanOccupied struct {
	Response_Student_GetCampusPlanOccupied Response_Student_GetCampusPlanOccupied `json:"student"`
}

type Response_Student_GetCampusPlanOccupied struct {
	Response_GetClusterPlanStudentsByClusterID_GetCampusPlanOccupied Response_GetClusterPlanStudentsByClusterID_GetCampusPlanOccupied `json:"getClusterPlanStudentsByClusterId"`
	Typename                          string                            `json:"__typename"`
}

type Response_GetClusterPlanStudentsByClusterID_GetCampusPlanOccupied struct {
	OccupiedPlaces []Response_OccupiedPlace_GetCampusPlanOccupied `json:"occupiedPlaces"`
	Typename       string          `json:"__typename"`
}

type Response_OccupiedPlace_GetCampusPlanOccupied struct {
	Row            string     `json:"row"`
	Number         int64      `json:"number"`
	StageGroupName string     `json:"stageGroupName"`
	StageName      string     `json:"stageName"`
	Response_User_GetCampusPlanOccupied           Response_User_GetCampusPlanOccupied       `json:"user"`
	Response_Experience_GetCampusPlanOccupied     Response_Experience_GetCampusPlanOccupied `json:"experience"`
	StudentType    string     `json:"studentType"`
	Typename       string     `json:"__typename"`
}

type Response_Experience_GetCampusPlanOccupied struct {
	ID       string `json:"id"`
	Value    int64  `json:"value"`
	Response_Level_GetCampusPlanOccupied    Response_Level_GetCampusPlanOccupied  `json:"level"`
	Typename string `json:"__typename"`
}

type Response_Level_GetCampusPlanOccupied struct {
	ID       int64  `json:"id"`
	Response_Range_GetCampusPlanOccupied    Response_Range_GetCampusPlanOccupied  `json:"range"`
	Typename string `json:"__typename"`
}

type Response_Range_GetCampusPlanOccupied struct {
	ID          string `json:"id"`
	LevelCode   int64  `json:"levelCode"`
	LeftBorder  int64  `json:"leftBorder"`
	RightBorder int64  `json:"rightBorder"`
	Typename    string `json:"__typename"`
}

type Response_User_GetCampusPlanOccupied struct {
	ID        string `json:"id"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatarUrl"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) GetCampusPlanOccupied(variables Request_Variables_GetCampusPlanOccupied) (Response_Data_GetCampusPlanOccupied, error) {
	request := gql.NewQueryRequest[Request_Variables_GetCampusPlanOccupied](
		"query getCampusPlanOccupied($clusterId: ID!) {   student {     getClusterPlanStudentsByClusterId(clusterId: $clusterId) {       occupiedPlaces {         row         number         stageGroupName         stageName         user {           id           login           avatarUrl           __typename         }         experience {           id           value           level {             id             range {               id               levelCode               leftBorder               rightBorder               __typename             }             __typename           }           __typename         }         studentType         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetCampusPlanOccupied](ctx, request)
}