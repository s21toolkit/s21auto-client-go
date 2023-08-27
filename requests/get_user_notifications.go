package requests

import "github.com/s21toolkit/s21client/gql"

type Variables_GetUserNotifications struct {
	Variables_Paging_GetUserNotifications Variables_Paging_GetUserNotifications `json:"paging"`
}

type Variables_Paging_GetUserNotifications struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetUserNotifications struct {
	Data_Student_GetUserNotifications Data_Student_GetUserNotifications `json:"student"`
}

type Data_Student_GetUserNotifications struct {
	Data_GetS21Notifications_GetUserNotifications Data_GetS21Notifications_GetUserNotifications `json:"getS21Notifications"`
	Typename            string              `json:"__typename"`
}

type Data_GetS21Notifications_GetUserNotifications struct {
	Notifications []Data_Notification_GetUserNotifications `json:"notifications"`
	TotalCount    int64          `json:"totalCount"`
	GroupNames    []string       `json:"groupNames"`
	Typename      string         `json:"__typename"`
}

type Data_Notification_GetUserNotifications struct {
	ID                string `json:"id"`
	RelatedObjectType string `json:"relatedObjectType"`
	RelatedObjectID   string `json:"relatedObjectId"`
	Message           string `json:"message"`
	Time              string `json:"time"`
	WasRead           bool   `json:"wasRead"`
	GroupName         string `json:"groupName"`
	Typename          string `json:"__typename"`
}


func (ctx *RequestContext) GetUserNotifications(variables Variables_GetUserNotifications) (Data_GetUserNotifications, error) {
	request := gql.NewQueryRequest[Variables_GetUserNotifications](
		"query getUserNotifications($paging: PagingInput!) {\n  student {\n    getS21Notifications(paging: $paging) {\n      notifications {\n        id\n        relatedObjectType\n        relatedObjectId\n        message\n        time\n        wasRead\n        groupName\n        __typename\n      }\n      totalCount\n      groupNames\n      __typename\n    }\n    __typename\n  }\n}\n",
		variables,
	)

	return GqlRequest[Data_GetUserNotifications](ctx, request)
}