package requests

import "s21client/gql"

type Variables_GetUserNotifications struct {
	Paging_GetUserNotifications Paging_GetUserNotifications `json:"paging"`
}

type Paging_GetUserNotifications struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetUserNotifications struct {
	Student_GetUserNotifications Student_GetUserNotifications `json:"student"`
}

type Student_GetUserNotifications struct {
	GetS21Notifications_GetUserNotifications GetS21Notifications_GetUserNotifications `json:"getS21Notifications"`
	Typename            string              `json:"__typename"`
}

type GetS21Notifications_GetUserNotifications struct {
	Notifications []Notification_GetUserNotifications `json:"notifications"`
	TotalCount    int64          `json:"totalCount"`
	GroupNames    []string       `json:"groupNames"`
	Typename      string         `json:"__typename"`
}

type Notification_GetUserNotifications struct {
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
		"query getUserNotifications($paging: PagingInput!) {   student {     getS21Notifications(paging: $paging) {       notifications {         id         relatedObjectType         relatedObjectId         message         time         wasRead         groupName         __typename       }       totalCount       groupNames       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetUserNotifications](ctx, request)
}