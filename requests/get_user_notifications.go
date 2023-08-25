package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetUserNotifications struct {
	Request_Paging_GetUserNotifications Request_Paging_GetUserNotifications `json:"paging"`
}

type Request_Paging_GetUserNotifications struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Response_Data_GetUserNotifications struct {
	Response_Student_GetUserNotifications Response_Student_GetUserNotifications `json:"student"`
}

type Response_Student_GetUserNotifications struct {
	Response_GetS21Notifications_GetUserNotifications Response_GetS21Notifications_GetUserNotifications `json:"getS21Notifications"`
	Typename            string              `json:"__typename"`
}

type Response_GetS21Notifications_GetUserNotifications struct {
	Notifications []Response_Notification_GetUserNotifications `json:"notifications"`
	TotalCount    int64          `json:"totalCount"`
	GroupNames    []string       `json:"groupNames"`
	Typename      string         `json:"__typename"`
}

type Response_Notification_GetUserNotifications struct {
	ID                string `json:"id"`
	RelatedObjectType string `json:"relatedObjectType"`
	RelatedObjectID   string `json:"relatedObjectId"`
	Message           string `json:"message"`
	Time              string `json:"time"`
	WasRead           bool   `json:"wasRead"`
	GroupName         string `json:"groupName"`
	Typename          string `json:"__typename"`
}

func (ctx *RequestContext) GetUserNotifications(variables Request_Variables_GetUserNotifications) (Response_Data_GetUserNotifications, error) {
	request := gql.NewQueryRequest[Request_Variables_GetUserNotifications](
		"query getUserNotifications($paging: PagingInput!) {   student {     getS21Notifications(paging: $paging) {       notifications {         id         relatedObjectType         relatedObjectId         message         time         wasRead         groupName         __typename       }       totalCount       groupNames       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Response_Data_GetUserNotifications](ctx, request)
}