package requests

import "github.com/s21toolkit/s21client/gql"

type CreateFilledChecklist_Variables struct {
	StudentAnswerID string `json:"studentAnswerId"`
}


type CreateFilledChecklist_Data struct {
	Student CreateFilledChecklist_Data_Student `json:"student"`
}

type CreateFilledChecklist_Data_Student struct {
	CreateFilledChecklist CreateFilledChecklist_Data_CreateFilledChecklist `json:"createFilledChecklist"`
	Typename              string                `json:"__typename"`
}

type CreateFilledChecklist_Data_CreateFilledChecklist struct {
	ID                      string                  `json:"id"`
	GitlabStudentProjectURL CreateFilledChecklist_Data_GitlabStudentProjectURL `json:"gitlabStudentProjectUrl"`
	Checklist               CreateFilledChecklist_Data_Checklist               `json:"checklist"`
	ModuleInfoP2P           CreateFilledChecklist_Data_ModuleInfoP2P           `json:"moduleInfoP2P"`
	ProgressCheckInfo       CreateFilledChecklist_Data_ProgressCheckInfo       `json:"progressCheckInfo"`
	VerifiableUsers         CreateFilledChecklist_Data_VerifiableUsers         `json:"verifiableUsers"`
	Video                   interface{}             `json:"video"`
	Typename                string                  `json:"__typename"`
}

type CreateFilledChecklist_Data_Checklist struct {
	Language     string        `json:"language"`
	Introduction string        `json:"introduction"`
	Guidelines   string        `json:"guidelines"`
	SectionList  []CreateFilledChecklist_Data_SectionList `json:"sectionList"`
	QuickActions []string      `json:"quickActions"`
	Typename     string        `json:"__typename"`
}

type CreateFilledChecklist_Data_SectionList struct {
	ChecklistSectionID string         `json:"checklistSectionId"`
	Name               string         `json:"name"`
	Description        string         `json:"description"`
	KindQuestionID     string         `json:"kindQuestionId"`
	QuestionList       []CreateFilledChecklist_Data_QuestionList `json:"questionList"`
	Typename           string         `json:"__typename"`
}

type CreateFilledChecklist_Data_QuestionList struct {
	SectionQuestionID   string              `json:"sectionQuestionId"`
	Name                string              `json:"name"`
	Description         string              `json:"description"`
	TaskAssessmentScale CreateFilledChecklist_Data_TaskAssessmentScale `json:"taskAssessmentScale"`
	Typename            string              `json:"__typename"`
}

type CreateFilledChecklist_Data_TaskAssessmentScale struct {
	CriterionScaleID string        `json:"criterionScaleId"`
	Type             string        `json:"type"`
	Description      string        `json:"description"`
	ScaleWeights     []CreateFilledChecklist_Data_ScaleWeight `json:"scaleWeights"`
	Typename         string        `json:"__typename"`
}

type CreateFilledChecklist_Data_ScaleWeight struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Typename string `json:"__typename"`
}

type CreateFilledChecklist_Data_GitlabStudentProjectURL struct {
	SSHLink   string `json:"sshLink"`
	HTTPSLink string `json:"httpsLink"`
	Typename  string `json:"__typename"`
}

type CreateFilledChecklist_Data_ModuleInfoP2P struct {
	ModuleName           string `json:"moduleName"`
	ExecutionType        string `json:"executionType"`
	PeriodOfVerification int64  `json:"periodOfVerification"`
	Typename             string `json:"__typename"`
}

type CreateFilledChecklist_Data_ProgressCheckInfo struct {
	ReviewUserCount         int64  `json:"reviewUserCount"`
	ReviewUserCountExecuted int64  `json:"reviewUserCountExecuted"`
	Typename                string `json:"__typename"`
}

type CreateFilledChecklist_Data_VerifiableUsers struct {
	TeamWithMembers CreateFilledChecklist_Data_TeamWithMembers `json:"teamWithMembers"`
	User            interface{}     `json:"user"`
	Typename        string          `json:"__typename"`
}

type CreateFilledChecklist_Data_TeamWithMembers struct {
	Team     CreateFilledChecklist_Data_Team     `json:"team"`
	Members  []CreateFilledChecklist_Data_Member `json:"members"`
	Typename string   `json:"__typename"`
}

type CreateFilledChecklist_Data_Member struct {
	User     CreateFilledChecklist_Data_User   `json:"user"`
	Role     string `json:"role"`
	Typename string `json:"__typename"`
}

type CreateFilledChecklist_Data_User struct {
	ID             string         `json:"id"`
	AvatarURL      string         `json:"avatarUrl"`
	Login          string         `json:"login"`
	UserExperience CreateFilledChecklist_Data_UserExperience `json:"userExperience"`
	Typename       string         `json:"__typename"`
}

type CreateFilledChecklist_Data_UserExperience struct {
	Level            CreateFilledChecklist_Data_Level  `json:"level"`
	CookiesCount     int64  `json:"cookiesCount"`
	CodeReviewPoints int64  `json:"codeReviewPoints"`
	Typename         string `json:"__typename"`
}

type CreateFilledChecklist_Data_Level struct {
	ID        int64  `json:"id"`
	LevelCode int64  `json:"levelCode"`
	Range     CreateFilledChecklist_Data_Range  `json:"range"`
	Typename  string `json:"__typename"`
}

type CreateFilledChecklist_Data_Range struct {
	LevelCode int64  `json:"levelCode"`
	Typename  string `json:"__typename"`
}

type CreateFilledChecklist_Data_Team struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Typename string `json:"__typename"`
}


func (ctx *RequestContext) CreateFilledChecklist(variables CreateFilledChecklist_Variables) (CreateFilledChecklist_Data, error) {
	request := gql.NewQueryRequest[CreateFilledChecklist_Variables](
		"mutation createFilledChecklist($studentAnswerId: ID!) {\n  student {\n    createFilledChecklist(studentAnswerId: $studentAnswerId) {\n      id\n      gitlabStudentProjectUrl {\n        sshLink\n        httpsLink\n        __typename\n      }\n      checklist {\n        ...FormChecklist\n        __typename\n      }\n      moduleInfoP2P {\n        ...FilledChecklistModuleInfo\n        __typename\n      }\n      progressCheckInfo {\n        reviewUserCount\n        reviewUserCountExecuted\n        __typename\n      }\n      verifiableUsers {\n        teamWithMembers {\n          team {\n            id\n            name\n            __typename\n          }\n          members {\n            ...TeamMember\n            __typename\n          }\n          __typename\n        }\n        user {\n          ...TeamMemberUser\n          __typename\n        }\n        __typename\n      }\n      video {\n        filledChecklistId\n        link\n        status\n        statusDetails\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment FormChecklist on Checklist {\n  language\n  introduction\n  guidelines\n  sectionList {\n    ...FormChecklistSection\n    __typename\n  }\n  quickActions\n  __typename\n}\n\nfragment FormChecklistSection on ChecklistSection {\n  checklistSectionId\n  name\n  description\n  kindQuestionId\n  questionList {\n    ...FormChecklistQuestion\n    __typename\n  }\n  __typename\n}\n\nfragment FormChecklistQuestion on SectionQuestion {\n  sectionQuestionId\n  name\n  description\n  taskAssessmentScale {\n    criterionScaleId\n    type\n    description\n    scaleWeights {\n      key\n      value\n      __typename\n    }\n    __typename\n  }\n  __typename\n}\n\nfragment FilledChecklistModuleInfo on ModuleInfoP2P {\n  moduleName\n  executionType\n  periodOfVerification\n  __typename\n}\n\nfragment TeamMember on TeamMember {\n  user {\n    ...TeamMemberUser\n    __typename\n  }\n  role\n  __typename\n}\n\nfragment TeamMemberUser on User {\n  id\n  avatarUrl\n  login\n  userExperience {\n    level {\n      id\n      levelCode\n      range {\n        levelCode\n        __typename\n      }\n      __typename\n    }\n    cookiesCount\n    codeReviewPoints\n    __typename\n  }\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[CreateFilledChecklist_Data](ctx, request)
}