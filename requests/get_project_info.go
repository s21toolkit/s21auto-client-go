package requests

import "github.com/s21toolkit/s21client/gql"

type GetProjectInfo_Variables struct {
	GoalID string `json:"goalId"`
}


type GetProjectInfo_Data struct {
	Student GetProjectInfo_Data_Student `json:"student"`
}

type GetProjectInfo_Data_Student struct {
	GetModuleByID                GetProjectInfo_Data_GetModuleByID                `json:"getModuleById"`
	GetModuleCoverInformation    GetProjectInfo_Data_GetModuleCoverInformation    `json:"getModuleCoverInformation"`
	GetP2PChecksInfo             GetProjectInfo_Data_GetP2PChecksInfo             `json:"getP2PChecksInfo"`
	GetStudentCodeReviewByGoalID GetProjectInfo_Data_GetStudentCodeReviewByGoalID `json:"getStudentCodeReviewByGoalId"`
	Typename                     string                       `json:"__typename"`
}

type GetProjectInfo_Data_GetModuleByID struct {
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
	CourseBaseParameters              interface{}   `json:"courseBaseParameters"`
	TeamSettings                      *GetProjectInfo_Data_TeamSettings `json:"teamSettings"`
	StudyModule                       GetProjectInfo_Data_StudyModule   `json:"studyModule"`
	CurrentTask                       GetProjectInfo_Data_CurrentTask   `json:"currentTask"`
	Typename                          string        `json:"__typename"`
}

type GetProjectInfo_Data_CurrentTask struct {
	ID           string          `json:"id"`
	TaskID       string          `json:"taskId"`
	Task         GetProjectInfo_Data_CurrentTaskTask `json:"task"`
	LastAnswer   *GetProjectInfo_Data_LastAnswer     `json:"lastAnswer"`
	TeamSettings *GetProjectInfo_Data_TeamSettings   `json:"teamSettings"`
	Typename     string          `json:"__typename"`
}

type GetProjectInfo_Data_LastAnswer struct {
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type GetProjectInfo_Data_CurrentTaskTask struct {
	ID                              string                          `json:"id"`
	AssignmentType                  string                          `json:"assignmentType"`
	StudentTaskAdditionalAttributes GetProjectInfo_Data_StudentTaskAdditionalAttributes `json:"studentTaskAdditionalAttributes"`
	CheckTypes                      []string                        `json:"checkTypes"`
	Typename                        string                          `json:"__typename"`
}

type GetProjectInfo_Data_StudentTaskAdditionalAttributes struct {
	CookiesCount       int64  `json:"cookiesCount"`
	MaxCodeReviewCount int64  `json:"maxCodeReviewCount"`
	CodeReviewCost     int64  `json:"codeReviewCost"`
	CiCDMode           string `json:"ciCdMode"`
	Typename           string `json:"__typename"`
}

type GetProjectInfo_Data_TeamSettings struct {
	TeamCreateOption    string `json:"teamCreateOption"`
	MinAmountMember     int64  `json:"minAmountMember"`
	MaxAmountMember     int64  `json:"maxAmountMember"`
	EnableSurrenderTeam bool   `json:"enableSurrenderTeam"`
	Typename            string `json:"__typename"`
}

type GetProjectInfo_Data_StudyModule struct {
	ID            string        `json:"id"`
	Idea          string        `json:"idea"`
	Duration      int64         `json:"duration"`
	GoalPoint     int64         `json:"goalPoint"`
	RetrySettings GetProjectInfo_Data_RetrySettings `json:"retrySettings"`
	Levels        []GetProjectInfo_Data_Level       `json:"levels"`
	Typename      string        `json:"__typename"`
}

type GetProjectInfo_Data_Level struct {
	ID           string        `json:"id"`
	GoalElements []GetProjectInfo_Data_GoalElement `json:"goalElements"`
	Typename     string        `json:"__typename"`
}

type GetProjectInfo_Data_GoalElement struct {
	ID       string        `json:"id"`
	Tasks    []GetProjectInfo_Data_TaskElement `json:"tasks"`
	Typename string        `json:"__typename"`
}

type GetProjectInfo_Data_TaskElement struct {
	ID       string `json:"id"`
	TaskID   string `json:"taskId"`
	Typename string `json:"__typename"`
}

type GetProjectInfo_Data_RetrySettings struct {
	MaxModuleAttempts   int64  `json:"maxModuleAttempts"`
	IsUnlimitedAttempts bool   `json:"isUnlimitedAttempts"`
	Typename            string `json:"__typename"`
}

type GetProjectInfo_Data_GetModuleCoverInformation struct {
	IsOwnStudentTimeline bool              `json:"isOwnStudentTimeline"`
	SoftSkills           []GetProjectInfo_Data_SoftSkill       `json:"softSkills"`
	Timeline             []GetProjectInfo_Data_Timeline        `json:"timeline"`
	ProjectStatistics    GetProjectInfo_Data_ProjectStatistics `json:"projectStatistics"`
	Typename             string            `json:"__typename"`
}

type GetProjectInfo_Data_ProjectStatistics struct {
	RegisteredStudents        int64                   `json:"registeredStudents"`
	InProgressStudents        int64                   `json:"inProgressStudents"`
	EvaluationStudents        int64                   `json:"evaluationStudents"`
	FinishedStudents          int64                   `json:"finishedStudents"`
	AcceptedStudents          int64                   `json:"acceptedStudents"`
	FailedStudents            int64                   `json:"failedStudents"`
	RetriedStudentsPercentage int64                   `json:"retriedStudentsPercentage"`
	GroupProjectStatistics    *GetProjectInfo_Data_GroupProjectStatistics `json:"groupProjectStatistics"`
	Typename                  string                  `json:"__typename"`
}

type GetProjectInfo_Data_GroupProjectStatistics struct {
	InProgressTeams int64  `json:"inProgressTeams"`
	EvaluationTeams int64  `json:"evaluationTeams"`
	FinishedTeams   int64  `json:"finishedTeams"`
	AcceptedTeams   int64  `json:"acceptedTeams"`
	FailedTeams     int64  `json:"failedTeams"`
	Typename        string `json:"__typename"`
}

type GetProjectInfo_Data_SoftSkill struct {
	SoftSkillID       int64   `json:"softSkillId"`
	SoftSkillName     string  `json:"softSkillName"`
	TotalPower        int64   `json:"totalPower"`
	MaxPower          int64   `json:"maxPower"`
	CurrentUserPower  int64   `json:"currentUserPower"`
	AchievedUserPower *int64  `json:"achievedUserPower"`
	TeamRole          *string `json:"teamRole"`
	Typename          string  `json:"__typename"`
}

type GetProjectInfo_Data_Timeline struct {
	Type     string      `json:"type"`
	Status   string      `json:"status"`
	Start    interface{} `json:"start"`
	End      interface{} `json:"end"`
	Children []GetProjectInfo_Data_Child     `json:"children"`
	Typename string      `json:"__typename"`
}

type GetProjectInfo_Data_Child struct {
	Type        string  `json:"type"`
	ElementType string  `json:"elementType"`
	Status      string  `json:"status"`
	Start       *string `json:"start"`
	End         *string `json:"end"`
	Order       int64   `json:"order"`
	Typename    string  `json:"__typename"`
}

type GetProjectInfo_Data_GetP2PChecksInfo struct {
	CookiesCount         int64              `json:"cookiesCount"`
	PeriodOfVerification int64              `json:"periodOfVerification"`
	ProjectReviewsInfo   GetProjectInfo_Data_ProjectReviewsInfo `json:"projectReviewsInfo"`
	Typename             string             `json:"__typename"`
}

type GetProjectInfo_Data_ProjectReviewsInfo struct {
	ReviewByStudentCount                 int64  `json:"reviewByStudentCount"`
	RelevantReviewByStudentsCount        int64  `json:"relevantReviewByStudentsCount"`
	ReviewByInspectionStaffCount         int64  `json:"reviewByInspectionStaffCount"`
	RelevantReviewByInspectionStaffCount int64  `json:"relevantReviewByInspectionStaffCount"`
	Typename                             string `json:"__typename"`
}

type GetProjectInfo_Data_GetStudentCodeReviewByGoalID struct {
	CountRound1     int64       `json:"countRound1"`
	CountRound2     int64       `json:"countRound2"`
	CodeReviewsInfo interface{} `json:"codeReviewsInfo"`
	Typename        string      `json:"__typename"`
}


func (ctx *RequestContext) GetProjectInfo(variables GetProjectInfo_Variables) (GetProjectInfo_Data, error) {
	request := gql.NewQueryRequest[GetProjectInfo_Variables](
		"query getProjectInfo($goalId: ID!) {\n  student {\n    getModuleById(goalId: $goalId) {\n      ...ProjectInfo\n      __typename\n    }\n    getModuleCoverInformation(goalId: $goalId) {\n      ...ModuleCoverInfo\n      __typename\n    }\n    getP2PChecksInfo(goalId: $goalId) {\n      ...P2PInfo\n      __typename\n    }\n    getStudentCodeReviewByGoalId(goalId: $goalId) {\n      ...StudentsCodeReview\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment ProjectInfo on StudentModule {\n  id\n  moduleTitle\n  finalPercentage\n  finalPoint\n  goalExecutionType\n  displayedGoalStatus\n  accessBeforeStartProgress\n  resultModuleCompletion\n  finishedExecutionDateByScheduler\n  durationFromStageSubjectGroupPlan\n  currentAttemptNumber\n  isDeadlineFree\n  isRetryAvailable\n  localCourseId\n  courseBaseParameters {\n    isGradedCourse\n    __typename\n  }\n  teamSettings {\n    ...teamSettingsInfo\n    __typename\n  }\n  studyModule {\n    id\n    idea\n    duration\n    goalPoint\n    retrySettings {\n      ...RetrySettings\n      __typename\n    }\n    levels {\n      id\n      goalElements {\n        id\n        tasks {\n          id\n          taskId\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n  currentTask {\n    ...CurrentInternshipTaskInfo\n    __typename\n  }\n  __typename\n}\n\nfragment teamSettingsInfo on TeamSettings {\n  teamCreateOption\n  minAmountMember\n  maxAmountMember\n  enableSurrenderTeam\n  __typename\n}\n\nfragment RetrySettings on ModuleAttemptsSettings {\n  maxModuleAttempts\n  isUnlimitedAttempts\n  __typename\n}\n\nfragment CurrentInternshipTaskInfo on StudentTask {\n  id\n  taskId\n  task {\n    id\n    assignmentType\n    studentTaskAdditionalAttributes {\n      cookiesCount\n      maxCodeReviewCount\n      codeReviewCost\n      ciCdMode\n      __typename\n    }\n    checkTypes\n    __typename\n  }\n  lastAnswer {\n    id\n    __typename\n  }\n  teamSettings {\n    ...teamSettingsInfo\n    __typename\n  }\n  __typename\n}\n\nfragment ModuleCoverInfo on ModuleCoverInformation {\n  isOwnStudentTimeline\n  softSkills {\n    softSkillId\n    softSkillName\n    totalPower\n    maxPower\n    currentUserPower\n    achievedUserPower\n    teamRole\n    __typename\n  }\n  timeline {\n    ...TimelineItem\n    __typename\n  }\n  projectStatistics {\n    ...ProjectStatistics\n    __typename\n  }\n  __typename\n}\n\nfragment TimelineItem on ProjectTimelineItem {\n  type\n  status\n  start\n  end\n  children {\n    ...TimelineItemChildren\n    __typename\n  }\n  __typename\n}\n\nfragment TimelineItemChildren on ProjectTimelineItem {\n  type\n  elementType\n  status\n  start\n  end\n  order\n  __typename\n}\n\nfragment ProjectStatistics on ProjectStatistics {\n  registeredStudents\n  inProgressStudents\n  evaluationStudents\n  finishedStudents\n  acceptedStudents\n  failedStudents\n  retriedStudentsPercentage\n  groupProjectStatistics {\n    ...GroupProjectStatistics\n    __typename\n  }\n  __typename\n}\n\nfragment GroupProjectStatistics on GroupProjectStatistics {\n  inProgressTeams\n  evaluationTeams\n  finishedTeams\n  acceptedTeams\n  failedTeams\n  __typename\n}\n\nfragment P2PInfo on P2PChecksInfo {\n  cookiesCount\n  periodOfVerification\n  projectReviewsInfo {\n    ...ProjectReviewsInfo\n    __typename\n  }\n  __typename\n}\n\nfragment ProjectReviewsInfo on ProjectReviewsInfo {\n  reviewByStudentCount\n  relevantReviewByStudentsCount\n  reviewByInspectionStaffCount\n  relevantReviewByInspectionStaffCount\n  __typename\n}\n\nfragment StudentsCodeReview on StudentCodeReviewsWithCountRound {\n  countRound1\n  countRound2\n  codeReviewsInfo {\n    maxCodeReviewCount\n    codeReviewDuration\n    codeReviewCost\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetProjectInfo_Data](ctx, request)
}