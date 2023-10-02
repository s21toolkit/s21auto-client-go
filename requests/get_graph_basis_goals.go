package requests

import "github.com/s21toolkit/s21client/gql"

type GetGraphBasisGoals_Variables struct {
	StudentID string `json:"studentId"`
}


type GetGraphBasisGoals_Data struct {
	Student GetGraphBasisGoals_Data_Student `json:"student"`
}

type GetGraphBasisGoals_Data_Student struct {
	GetBasisGraph GetGraphBasisGoals_Data_GetBasisGraph `json:"getBasisGraph"`
	Typename      string        `json:"__typename"`
}

type GetGraphBasisGoals_Data_GetBasisGraph struct {
	IsIntensiveGraphAvailable bool        `json:"isIntensiveGraphAvailable"`
	GraphNodes                []GetGraphBasisGoals_Data_GraphNode `json:"graphNodes"`
	Typename                  string      `json:"__typename"`
}

type GetGraphBasisGoals_Data_GraphNode struct {
	GraphNodeID     string           `json:"graphNodeId"`
	NodeCode        string           `json:"nodeCode"`
	StudyDirections []GetGraphBasisGoals_Data_StudyDirection `json:"studyDirections"`
	EntityType      string           `json:"entityType"`
	EntityID        string           `json:"entityId"`
	Goal            *GetGraphBasisGoals_Data_Course          `json:"goal"`
	Course          *GetGraphBasisGoals_Data_Course          `json:"course"`
	ParentNodeCodes []string         `json:"parentNodeCodes"`
	Typename        string           `json:"__typename"`
}

type GetGraphBasisGoals_Data_Course struct {
	CourseType         *string     `json:"courseType,omitempty"`
	ProjectState       string      `json:"projectState"`
	ProjectName        string      `json:"projectName"`
	ProjectDescription string      `json:"projectDescription"`
	ProjectPoints      int64       `json:"projectPoints"`
	ProjectDate        interface{} `json:"projectDate"`
	Duration           int64       `json:"duration"`
	LocalCourseID      *int64      `json:"localCourseId,omitempty"`
	Typename           string      `json:"__typename"`
	GoalExecutionType  *string     `json:"goalExecutionType,omitempty"`
	IsMandatory        *bool       `json:"isMandatory,omitempty"`
}

type GetGraphBasisGoals_Data_StudyDirection struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	TextColor string `json:"textColor"`
	Typename  string `json:"__typename"`
}


func (ctx *RequestContext) GetGraphBasisGoals(variables GetGraphBasisGoals_Variables) (GetGraphBasisGoals_Data, error) {
	request := gql.NewQueryRequest[GetGraphBasisGoals_Variables](
		"query getGraphBasisGoals($studentId: UUID!) {\n  student {\n    getBasisGraph(studentId: $studentId) {\n      isIntensiveGraphAvailable\n      graphNodes {\n        ...BasisGraphNode\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}\n\nfragment BasisGraphNode on GraphNode {\n  graphNodeId\n  nodeCode\n  studyDirections {\n    id\n    name\n    color\n    textColor\n    __typename\n  }\n  entityType\n  entityId\n  goal {\n    goalExecutionType\n    projectState\n    projectName\n    projectDescription\n    projectPoints\n    projectDate\n    duration\n    isMandatory\n    __typename\n  }\n  course {\n    courseType\n    projectState\n    projectName\n    projectDescription\n    projectPoints\n    projectDate\n    duration\n    localCourseId\n    __typename\n  }\n  parentNodeCodes\n  __typename\n}\n",
		variables,
	)

	return GqlRequest[GetGraphBasisGoals_Data](ctx, request)
}