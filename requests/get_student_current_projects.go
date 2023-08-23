package requests

import "s21client/gql"

type Variables_GetStudentCurrentProjects struct {
	UserID string `json:"userId"`
}


type Data_GetStudentCurrentProjects struct {
	Student_GetStudentCurrentProjects Student_GetStudentCurrentProjects `json:"student"`
}

type Student_GetStudentCurrentProjects struct {
	GetStudentCurrentProjects []GetStudentCurrentProject_GetStudentCurrentProjects `json:"getStudentCurrentProjects"`
	Typename                  string                     `json:"__typename"`
}

type GetStudentCurrentProject_GetStudentCurrentProjects struct {
	GoalID                         *int64      `json:"goalId"`
	Name                           string      `json:"name"`
	Description                    string      `json:"description"`
	Experience                     int64       `json:"experience"`
	DateTime                       interface{} `json:"dateTime"`
	FinalPercentage                interface{} `json:"finalPercentage"`
	Laboriousness                  int64       `json:"laboriousness"`
	ExecutionType                  *string     `json:"executionType"`
	GoalStatus                     *string     `json:"goalStatus"`
	CourseType                     *string     `json:"courseType"`
	DisplayedCourseStatus          *string     `json:"displayedCourseStatus"`
	AmountAnswers                  interface{} `json:"amountAnswers"`
	AmountMembers                  *int64      `json:"amountMembers"`
	AmountJoinedMembers            *int64      `json:"amountJoinedMembers"`
	AmountReviewedAnswers          interface{} `json:"amountReviewedAnswers"`
	AmountCodeReviewMembers        *int64      `json:"amountCodeReviewMembers"`
	AmountCurrentCodeReviewMembers *int64      `json:"amountCurrentCodeReviewMembers"`
	GroupName                      string      `json:"groupName"`
	LocalCourseID                  *int64      `json:"localCourseId"`
	Typename                       string      `json:"__typename"`
}

func (ctx *RequestContext) GetStudentCurrentProjects(variables Variables_GetStudentCurrentProjects) (Data_GetStudentCurrentProjects, error) {
	request := gql.NewQueryRequest[Variables_GetStudentCurrentProjects](
		"query getStudentCurrentProjects($userId: ID!) {   student {     getStudentCurrentProjects(userId: $userId) {       ...StudentProjectItem       __typename     }     __typename   } }  fragment StudentProjectItem on StudentItem {   goalId   name   description   experience   dateTime   finalPercentage   laboriousness   executionType   goalStatus   courseType   displayedCourseStatus   amountAnswers   amountMembers   amountJoinedMembers   amountReviewedAnswers   amountCodeReviewMembers   amountCurrentCodeReviewMembers   groupName   localCourseId   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetStudentCurrentProjects](ctx, request)
}