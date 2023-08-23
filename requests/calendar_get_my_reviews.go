package requests

import "s21client/gql"

type Variables_CalendarGetMyReviews struct {
	Limit int64 `json:"limit"`
}


type Data_CalendarGetMyReviews struct {
	Student_CalendarGetMyReviews Student_CalendarGetMyReviews `json:"student"`
}

type Student_CalendarGetMyReviews struct {
	GetMyUpcomingBookings []interface{} `json:"getMyUpcomingBookings"`
	Typename              string        `json:"__typename"`
}

func (ctx *RequestContext) CalendarGetMyReviews(variables Variables_CalendarGetMyReviews) (Data_CalendarGetMyReviews, error) {
	request := gql.NewQueryRequest[Variables_CalendarGetMyReviews](
		"query calendarGetMyReviews($to: DateTime, $limit: Int) {   student {     getMyUpcomingBookings(to: $to, limit: $limit) {       ...Review       __typename     }     __typename   } }  fragment Review on CalendarBooking {   id   answerId   eventSlot {     id     start     end     __typename   }   task {     id     title     assignmentType     goalId     goalName     studentTaskAdditionalAttributes {       cookiesCount       __typename     }     __typename   }   verifierUser {     ...UserInBooking     __typename   }   verifiableStudent {     id     user {       ...UserInBooking       __typename     }     __typename   }   team {     ...ProjectTeamMembers     __typename   }   bookingStatus   isOnline   vcLinkUrl   __typename }  fragment UserInBooking on User {   id   login   avatarUrl   userExperience {     level {       id       range {         levelCode         __typename       }       __typename     }     __typename   }   __typename }  fragment ProjectTeamMembers on ProjectTeamMembers {   id   teamLead {     ...ProjectTeamMember     __typename   }   members {     ...ProjectTeamMember     __typename   }   invitedUsers {     ...ProjectTeamMember     __typename   }   teamName   teamStatus   minTeamMemberCount   maxTeamMemberCount   __typename }  fragment ProjectTeamMember on User {   id   avatarUrl   login   userExperience {     level {       id       range {         levelCode         __typename       }       __typename     }     cookiesCount     codeReviewPoints     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_CalendarGetMyReviews](ctx, request)
}