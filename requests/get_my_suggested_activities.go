package requests

import "s21client/gql"

type Variables_GetMySuggestedActivities struct {
	Page_GetMySuggestedActivities Page_GetMySuggestedActivities `json:"page"`
}

type Page_GetMySuggestedActivities struct {
	Offset int64 `json:"offset"`
	Limit  int64 `json:"limit"`
}


type Data_GetMySuggestedActivities struct {
	School21_GetMySuggestedActivities School21_GetMySuggestedActivities `json:"school21"`
}

type School21_GetMySuggestedActivities struct {
	GetMySuggestedActivities []interface{} `json:"getMySuggestedActivities"`
	Typename                 string        `json:"__typename"`
}

func (ctx *RequestContext) GetMySuggestedActivities(variables Variables_GetMySuggestedActivities) (Data_GetMySuggestedActivities, error) {
	request := gql.NewQueryRequest[Variables_GetMySuggestedActivities](
		"query getMySuggestedActivities($page: PagingInput!, $statuses: [ParticipantEventStatus]) {   school21 {     getMySuggestedActivities(page: $page, statuses: $statuses) {       id       start       end       eventType       description       eventCode       activity {         organizers {           id           login           __typename         }         eventCreator {           id           login           __typename         }         comments {           type           createTs           comment           __typename         }         averageFeedbackRating         isVisible         activityType         status         activityEventId         eventId         name         description         location         currentStudentsCount         maxStudentCount         isRegistered         isInWaitList         isWaitListActive         stopRegisterDate         beginDate         endDate         __typename       }       __typename     }     __typename   } } ",
		variables,
	)

	return GqlRequest[Data_GetMySuggestedActivities](ctx, request)
}