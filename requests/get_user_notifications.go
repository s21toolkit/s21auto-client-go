package requests

import "github.com/s21toolkit/s21client/gql"

type GetUserNotifications_Variables struct {
	Paging GetUserNotifications_Variables_Paging `json:"paging"`
}

type GetUserNotifications_Variables_Paging struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type GetUserNotifications_Data struct {
	Student GetUserNotifications_Data_Student `json:"student"`
}

type GetUserNotifications_Data_Student struct {
	GetS21Notifications GetUserNotifications_Data_GetS21Notifications `json:"getS21Notifications"`
	Typename            string              `json:"__typename"`
}

type GetUserNotifications_Data_GetS21Notifications struct {
	Notifications []GetUserNotifications_Data_Notification `json:"notifications"`
	TotalCount    int64          `json:"totalCount"`
	GroupNames    []string       `json:"groupNames"`
	Typename      string         `json:"__typename"`
}

type GetUserNotifications_Data_Notification struct {
	ID                string `json:"id"`
	RelatedObjectType string `json:"relatedObjectType"`
	RelatedObjectID   string `json:"relatedObjectId"`
	Message           string `json:"message"`
	Time              string `json:"time"`
	WasRead           bool   `json:"wasRead"`
	GroupName         string `json:"groupName"`
	Typename          string `json:"__typename"`
}


func (ctx *RequestContext) GetUserNotifications(variables GetUserNotifications_Variables) (GetUserNotifications_Data, error) {
	request := gql.NewQueryRequest[GetUserNotifications_Variables](
		"query getUserNotifications($paging: PagingInput!) {\n  student {\n    getS21Notifications(paging: $paging) {\n      notifications {\n        id\n        relatedObjectType\n        relatedObjectId\n        message\n        time\n        wasRead\n        groupName\n        __typename\n      }\n      totalCount\n      groupNames\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[GetUserNotifications_Data](ctx, request)
}