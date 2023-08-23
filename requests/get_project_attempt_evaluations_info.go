package requests

import "s21client/gql"

type Variables_GetProjectAttemptEvaluationsInfo struct {
	GoalID    string `json:"goalId"`
	StudentID string `json:"studentId"`
}


type Data_GetProjectAttemptEvaluationsInfo struct {
	School21_GetProjectAttemptEvaluationsInfo School21_GetProjectAttemptEvaluationsInfo `json:"school21"`
}

type School21_GetProjectAttemptEvaluationsInfo struct {
	GetProjectAttemptEvaluationsInfoV1_GetProjectAttemptEvaluationsInfo []GetProjectAttemptEvaluationsInfoV1_GetProjectAttemptEvaluationsInfo `json:"getProjectAttemptEvaluationsInfo_V1"`
	Typename                           string                               `json:"__typename"`
}

type GetProjectAttemptEvaluationsInfoV1_GetProjectAttemptEvaluationsInfo struct {
	StudentAnswerID      string        `json:"studentAnswerId"`
	StudentGoalAttemptID string        `json:"studentGoalAttemptId"`
	AttemptResult_GetProjectAttemptEvaluationsInfo        AttemptResult_GetProjectAttemptEvaluationsInfo `json:"attemptResult"`
	Team                 interface{}   `json:"team"`
	P2P_GetProjectAttemptEvaluationsInfo                  []P2P_GetProjectAttemptEvaluationsInfo         `json:"p2p"`
	Auto_GetProjectAttemptEvaluationsInfo                 Auto_GetProjectAttemptEvaluationsInfo          `json:"auto"`
	CodeReview_GetProjectAttemptEvaluationsInfo           CodeReview_GetProjectAttemptEvaluationsInfo    `json:"codeReview"`
	Typename             string        `json:"__typename"`
}

type AttemptResult_GetProjectAttemptEvaluationsInfo struct {
	FinalPointProject      int64  `json:"finalPointProject"`
	FinalPercentageProject int64  `json:"finalPercentageProject"`
	ResultModuleCompletion string `json:"resultModuleCompletion"`
	ResultDate             string `json:"resultDate"`
	Typename               string `json:"__typename"`
}

type Auto_GetProjectAttemptEvaluationsInfo struct {
	Status             string      `json:"status"`
	ReceivedPercentage int64       `json:"receivedPercentage"`
	EndTimeCheck       interface{} `json:"endTimeCheck"`
	ResultInfo         interface{} `json:"resultInfo"`
	Typename           string      `json:"__typename"`
}

type CodeReview_GetProjectAttemptEvaluationsInfo struct {
	AverageMark        string        `json:"averageMark"`
	StudentCodeReviews []interface{} `json:"studentCodeReviews"`
	Typename           string        `json:"__typename"`
}

type P2P_GetProjectAttemptEvaluationsInfo struct {
	Status    string    `json:"status"`
	Checklist_GetProjectAttemptEvaluationsInfo Checklist_GetProjectAttemptEvaluationsInfo `json:"checklist"`
	Typename  string    `json:"__typename"`
}

type Checklist_GetProjectAttemptEvaluationsInfo struct {
	ID                 string          `json:"id"`
	ChecklistID        string          `json:"checklistId"`
	EndTimeCheck       string          `json:"endTimeCheck"`
	StartTimeCheck     string          `json:"startTimeCheck"`
	Reviewer_GetProjectAttemptEvaluationsInfo           Reviewer_GetProjectAttemptEvaluationsInfo        `json:"reviewer"`
	ReviewFeedback_GetProjectAttemptEvaluationsInfo     *ReviewFeedback_GetProjectAttemptEvaluationsInfo `json:"reviewFeedback"`
	Comment            string          `json:"comment"`
	ReceivedPoint      int64           `json:"receivedPoint"`
	ReceivedPercentage int64           `json:"receivedPercentage"`
	QuickAction        interface{}     `json:"quickAction"`
	CheckType          string          `json:"checkType"`
	Video              interface{}     `json:"video"`
	Typename           string          `json:"__typename"`
}

type ReviewFeedback_GetProjectAttemptEvaluationsInfo struct {
	ID                           string                        `json:"id"`
	Comment                      string                        `json:"comment"`
	FilledChecklist_GetProjectAttemptEvaluationsInfo              FilledChecklist_GetProjectAttemptEvaluationsInfo               `json:"filledChecklist"`
	ReviewFeedbackCategoryValues []ReviewFeedbackCategoryValue_GetProjectAttemptEvaluationsInfo `json:"reviewFeedbackCategoryValues"`
	Typename                     string                        `json:"__typename"`
}

type FilledChecklist_GetProjectAttemptEvaluationsInfo struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type ReviewFeedbackCategoryValue_GetProjectAttemptEvaluationsInfo struct {
	FeedbackCategory string `json:"feedbackCategory"`
	FeedbackValue    string `json:"feedbackValue"`
	ID               string `json:"id"`
	Typename         string `json:"__typename"`
}

type Reviewer_GetProjectAttemptEvaluationsInfo struct {
	AvatarURL          string        `json:"avatarUrl"`
	Login              string        `json:"login"`
	BusinessAdminRoles []interface{} `json:"businessAdminRoles"`
	Typename           string        `json:"__typename"`
}

func (ctx *RequestContext) GetProjectAttemptEvaluationsInfo(variables Variables_GetProjectAttemptEvaluationsInfo) (Data_GetProjectAttemptEvaluationsInfo, error) {
	request := gql.NewQueryRequest[Variables_GetProjectAttemptEvaluationsInfo](
		"query getProjectAttemptEvaluationsInfo($goalId: ID!, $studentId: UUID!) {   school21 {     getProjectAttemptEvaluationsInfo_V1(goalId: $goalId, studentId: $studentId) {       ...ProjectAttemptEvaluations_V1       __typename     }     __typename   } }  fragment ProjectAttemptEvaluations_V1 on ProjectAttemptEvaluationsInfo_V1 {   studentAnswerId   studentGoalAttemptId   attemptResult {     ...AtemptResult     __typename   }   team {     ...AttemptTeamWithMembers     __typename   }   p2p {     ...P2PEvaluation     __typename   }   auto {     status     receivedPercentage     endTimeCheck     resultInfo     __typename   }   codeReview {     averageMark     studentCodeReviews {       user {         avatarUrl         login         __typename       }       finalMark       markTime       reviewerCommentsCount       __typename     }     __typename   }   __typename }  fragment AtemptResult on StudentGoalAttempt {   finalPointProject   finalPercentageProject   resultModuleCompletion   resultDate   __typename }  fragment AttemptTeamWithMembers on TeamWithMembers {   team {     id     name     __typename   }   members {     ...TeamMemberWithRole     __typename   }   __typename }  fragment TeamMemberWithRole on TeamMember {   role   user {     ...ProjectTeamMember     __typename   }   __typename }  fragment ProjectTeamMember on User {   id   avatarUrl   login   userExperience {     level {       id       range {         levelCode         __typename       }       __typename     }     cookiesCount     codeReviewPoints     __typename   }   __typename }  fragment P2PEvaluation on P2PEvaluationInfo {   status   checklist {     ...Checklist     __typename   }   __typename }  fragment Checklist on FilledChecklist {   id   checklistId   endTimeCheck   startTimeCheck   reviewer {     avatarUrl     login     businessAdminRoles {       id       school {         id         organizationType         __typename       }       __typename     }     __typename   }   reviewFeedback {     ...EvaluationFeedback     __typename   }   comment   receivedPoint   receivedPercentage   quickAction   checkType   video {     ...P2PReviewVideo     __typename   }   __typename }  fragment EvaluationFeedback on ReviewFeedback {   id   comment   filledChecklist {     id     __typename   }   reviewFeedbackCategoryValues {     feedbackCategory     feedbackValue     id     __typename   }   __typename }  fragment P2PReviewVideo on OnlineReviewVideo {   link   status   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetProjectAttemptEvaluationsInfo](ctx, request)
}