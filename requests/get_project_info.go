package requests

import "s21client/gql"

type Variables_GetProjectInfo struct {
	GoalID string `json:"goalId"`
}


type Data_GetProjectInfo struct {
	Student_GetProjectInfo Student_GetProjectInfo `json:"student"`
}

type Student_GetProjectInfo struct {
	GetModuleByID_GetProjectInfo                GetModuleByID_GetProjectInfo                `json:"getModuleById"`
	GetModuleCoverInformation_GetProjectInfo    GetModuleCoverInformation_GetProjectInfo    `json:"getModuleCoverInformation"`
	GetP2PChecksInfo_GetProjectInfo             GetP2PChecksInfo_GetProjectInfo             `json:"getP2PChecksInfo"`
	GetStudentCodeReviewByGoalID_GetProjectInfo GetStudentCodeReviewByGoalID_GetProjectInfo `json:"getStudentCodeReviewByGoalId"`
	Typename                     string                       `json:"__typename"`
}

type GetModuleByID_GetProjectInfo struct {
	ID                                string        `json:"id"`
	ModuleTitle                       string        `json:"moduleTitle"`
	FinalPercentage                   *int64        `json:"finalPercentage"`
	FinalPoint                        *int64        `json:"finalPoint"`
	GoalExecutionType                 string        `json:"goalExecutionType"`
	DisplayedGoalStatus               string        `json:"displayedGoalStatus"`
	AccessBeforeStartProgress         bool          `json:"accessBeforeStartProgress"`
	ResultModuleCompletion            *string       `json:"resultModuleCompletion"`
	FinishedExecutionDateByScheduler  interface{}   `json:"finishedExecutionDateByScheduler"`
	DurationFromStageSubjectGroupPlan int64         `json:"durationFromStageSubjectGroupPlan"`
	CurrentAttemptNumber              int64         `json:"currentAttemptNumber"`
	IsDeadlineFree                    bool          `json:"isDeadlineFree"`
	IsRetryAvailable                  bool          `json:"isRetryAvailable"`
	LocalCourseID                     interface{}   `json:"localCourseId"`
	TeamSettings_GetProjectInfo                      *TeamSettings_GetProjectInfo `json:"teamSettings"`
	StudyModule_GetProjectInfo                       StudyModule_GetProjectInfo   `json:"studyModule"`
	CurrentTask_GetProjectInfo                       CurrentTask_GetProjectInfo   `json:"currentTask"`
	Typename                          string        `json:"__typename"`
}

type CurrentTask_GetProjectInfo struct {
	ID           string          `json:"id"`
	TaskID       string          `json:"taskId"`
	Task         CurrentTaskTask_GetProjectInfo `json:"task"`
	LastAnswer_GetProjectInfo   *LastAnswer_GetProjectInfo     `json:"lastAnswer"`
	TeamSettings_GetProjectInfo *TeamSettings_GetProjectInfo   `json:"teamSettings"`
	Typename     string          `json:"__typename"`
}

type LastAnswer_GetProjectInfo struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type CurrentTaskTask_GetProjectInfo struct {
	ID                              string                          `json:"id"`
	AssignmentType                  string                          `json:"assignmentType"`
	StudentTaskAdditionalAttributes_GetProjectInfo StudentTaskAdditionalAttributes_GetProjectInfo `json:"studentTaskAdditionalAttributes"`
	CheckTypes                      []string                        `json:"checkTypes"`
	Typename                        string                          `json:"__typename"`
}

type StudentTaskAdditionalAttributes_GetProjectInfo struct {
	CookiesCount       int64  `json:"cookiesCount"`
	MaxCodeReviewCount int64  `json:"maxCodeReviewCount"`
	CodeReviewCost     int64  `json:"codeReviewCost"`
	CiCDMode           string `json:"ciCdMode"`
	Typename           string `json:"__typename"`
}

type TeamSettings_GetProjectInfo struct {
	TeamCreateOption    string `json:"teamCreateOption"`
	MinAmountMember     int64  `json:"minAmountMember"`
	MaxAmountMember     int64  `json:"maxAmountMember"`
	EnableSurrenderTeam bool   `json:"enableSurrenderTeam"`
	Typename            string `json:"__typename"`
}

type StudyModule_GetProjectInfo struct {
	ID            string        `json:"id"`
	Idea          string        `json:"idea"`
	Duration      int64         `json:"duration"`
	GoalPoint     int64         `json:"goalPoint"`
	RetrySettings_GetProjectInfo RetrySettings_GetProjectInfo `json:"retrySettings"`
	Levels        []Level_GetProjectInfo       `json:"levels"`
	Typename      string        `json:"__typename"`
}

type Level_GetProjectInfo struct {
	ID           string        `json:"id"`
	GoalElements []GoalElement_GetProjectInfo `json:"goalElements"`
	Typename     string        `json:"__typename"`
}

type GoalElement_GetProjectInfo struct {
	ID       string        `json:"id"`
	Tasks    []TaskElement_GetProjectInfo `json:"tasks"`
	Typename string        `json:"__typename"`
}

type TaskElement_GetProjectInfo struct {
	ID       string `json:"id"`
	TaskID   string `json:"taskId"`
	Typename string `json:"__typename"`
}

type RetrySettings_GetProjectInfo struct {
	MaxModuleAttempts   int64  `json:"maxModuleAttempts"`
	IsUnlimitedAttempts bool   `json:"isUnlimitedAttempts"`
	Typename            string `json:"__typename"`
}

type GetModuleCoverInformation_GetProjectInfo struct {
	IsOwnStudentTimeline bool              `json:"isOwnStudentTimeline"`
	SoftSkills           []SoftSkill_GetProjectInfo       `json:"softSkills"`
	Timeline_GetProjectInfo             []Timeline_GetProjectInfo        `json:"timeline"`
	ProjectStatistics_GetProjectInfo    ProjectStatistics_GetProjectInfo `json:"projectStatistics"`
	Typename             string            `json:"__typename"`
}

type ProjectStatistics_GetProjectInfo struct {
	RegisteredStudents        int64                   `json:"registeredStudents"`
	InProgressStudents        int64                   `json:"inProgressStudents"`
	EvaluationStudents        int64                   `json:"evaluationStudents"`
	FinishedStudents          int64                   `json:"finishedStudents"`
	AcceptedStudents          int64                   `json:"acceptedStudents"`
	FailedStudents            int64                   `json:"failedStudents"`
	RetriedStudentsPercentage int64                   `json:"retriedStudentsPercentage"`
	GroupProjectStatistics_GetProjectInfo    *GroupProjectStatistics_GetProjectInfo `json:"groupProjectStatistics"`
	Typename                  string                  `json:"__typename"`
}

type GroupProjectStatistics_GetProjectInfo struct {
	InProgressTeams int64  `json:"inProgressTeams"`
	EvaluationTeams int64  `json:"evaluationTeams"`
	FinishedTeams   int64  `json:"finishedTeams"`
	AcceptedTeams   int64  `json:"acceptedTeams"`
	FailedTeams     int64  `json:"failedTeams"`
	Typename        string `json:"__typename"`
}

type SoftSkill_GetProjectInfo struct {
	SoftSkillID       int64   `json:"softSkillId"`
	SoftSkillName     string  `json:"softSkillName"`
	TotalPower        int64   `json:"totalPower"`
	MaxPower          int64   `json:"maxPower"`
	CurrentUserPower  int64   `json:"currentUserPower"`
	AchievedUserPower *int64  `json:"achievedUserPower"`
	TeamRole          *string `json:"teamRole"`
	Typename          string  `json:"__typename"`
}

type Timeline_GetProjectInfo struct {
	Type     string      `json:"type"`
	Status   string      `json:"status"`
	Start    interface{} `json:"start"`
	End      interface{} `json:"end"`
	Children []Child_GetProjectInfo     `json:"children"`
	Typename string      `json:"__typename"`
}

type Child_GetProjectInfo struct {
	Type        string  `json:"type"`
	ElementType string  `json:"elementType"`
	Status      string  `json:"status"`
	Start       *string `json:"start"`
	End         *string `json:"end"`
	Order       int64   `json:"order"`
	Typename    string  `json:"__typename"`
}

type GetP2PChecksInfo_GetProjectInfo struct {
	CookiesCount         int64              `json:"cookiesCount"`
	PeriodOfVerification int64              `json:"periodOfVerification"`
	ProjectReviewsInfo_GetProjectInfo   ProjectReviewsInfo_GetProjectInfo `json:"projectReviewsInfo"`
	Typename             string             `json:"__typename"`
}

type ProjectReviewsInfo_GetProjectInfo struct {
	ReviewByStudentCount                 int64  `json:"reviewByStudentCount"`
	RelevantReviewByStudentsCount        int64  `json:"relevantReviewByStudentsCount"`
	ReviewByInspectionStaffCount         int64  `json:"reviewByInspectionStaffCount"`
	RelevantReviewByInspectionStaffCount int64  `json:"relevantReviewByInspectionStaffCount"`
	Typename                             string `json:"__typename"`
}

type GetStudentCodeReviewByGoalID_GetProjectInfo struct {
	CountRound1     int64       `json:"countRound1"`
	CountRound2     int64       `json:"countRound2"`
	CodeReviewsInfo interface{} `json:"codeReviewsInfo"`
	Typename        string      `json:"__typename"`
}

func (ctx *RequestContext) GetProjectInfo(variables Variables_GetProjectInfo) (Data_GetProjectInfo, error) {
	request := gql.NewQueryRequest[Variables_GetProjectInfo](
		"query getProjectInfo($goalId: ID!) {   student {     getModuleById(goalId: $goalId) {       ...ProjectInfo       __typename     }     getModuleCoverInformation(goalId: $goalId) {       ...ModuleCoverInfo       __typename     }     getP2PChecksInfo(goalId: $goalId) {       ...P2PInfo       __typename     }     getStudentCodeReviewByGoalId(goalId: $goalId) {       ...StudentsCodeReview       __typename     }     __typename   } }  fragment ProjectInfo on StudentModule {   id   moduleTitle   finalPercentage   finalPoint   goalExecutionType   displayedGoalStatus   accessBeforeStartProgress   resultModuleCompletion   finishedExecutionDateByScheduler   durationFromStageSubjectGroupPlan   currentAttemptNumber   isDeadlineFree   isRetryAvailable   localCourseId   teamSettings {     ...teamSettingsInfo     __typename   }   studyModule {     id     idea     duration     goalPoint     retrySettings {       ...RetrySettings       __typename     }     levels {       id       goalElements {         id         tasks {           id           taskId           __typename         }         __typename       }       __typename     }     __typename   }   currentTask {     ...CurrentInternshipTaskInfo     __typename   }   __typename }  fragment teamSettingsInfo on TeamSettings {   teamCreateOption   minAmountMember   maxAmountMember   enableSurrenderTeam   __typename }  fragment RetrySettings on ModuleAttemptsSettings {   maxModuleAttempts   isUnlimitedAttempts   __typename }  fragment CurrentInternshipTaskInfo on StudentTask {   id   taskId   task {     id     assignmentType     studentTaskAdditionalAttributes {       cookiesCount       maxCodeReviewCount       codeReviewCost       ciCdMode       __typename     }     checkTypes     __typename   }   lastAnswer {     id     __typename   }   teamSettings {     ...teamSettingsInfo     __typename   }   __typename }  fragment ModuleCoverInfo on ModuleCoverInformation {   isOwnStudentTimeline   softSkills {     softSkillId     softSkillName     totalPower     maxPower     currentUserPower     achievedUserPower     teamRole     __typename   }   timeline {     ...TimelineItem     __typename   }   projectStatistics {     ...ProjectStatistics     __typename   }   __typename }  fragment TimelineItem on ProjectTimelineItem {   type   status   start   end   children {     ...TimelineItemChildren     __typename   }   __typename }  fragment TimelineItemChildren on ProjectTimelineItem {   type   elementType   status   start   end   order   __typename }  fragment ProjectStatistics on ProjectStatistics {   registeredStudents   inProgressStudents   evaluationStudents   finishedStudents   acceptedStudents   failedStudents   retriedStudentsPercentage   groupProjectStatistics {     ...GroupProjectStatistics     __typename   }   __typename }  fragment GroupProjectStatistics on GroupProjectStatistics {   inProgressTeams   evaluationTeams   finishedTeams   acceptedTeams   failedTeams   __typename }  fragment P2PInfo on P2PChecksInfo {   cookiesCount   periodOfVerification   projectReviewsInfo {     ...ProjectReviewsInfo     __typename   }   __typename }  fragment ProjectReviewsInfo on ProjectReviewsInfo {   reviewByStudentCount   relevantReviewByStudentsCount   reviewByInspectionStaffCount   relevantReviewByInspectionStaffCount   __typename }  fragment StudentsCodeReview on StudentCodeReviewsWithCountRound {   countRound1   countRound2   codeReviewsInfo {     maxCodeReviewCount     codeReviewDuration     codeReviewCost     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetProjectInfo](ctx, request)
}