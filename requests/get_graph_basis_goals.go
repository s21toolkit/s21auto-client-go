package requests

import "s21client/gql"

type Variables_GetGraphBasisGoals struct {
	StudentID string `json:"studentId"`
}


type Data_GetGraphBasisGoals struct {
	Student_GetGraphBasisGoals Student_GetGraphBasisGoals `json:"student"`
}

type Student_GetGraphBasisGoals struct {
	GetBasisGraph_GetGraphBasisGoals GetBasisGraph_GetGraphBasisGoals `json:"getBasisGraph"`
	Typename      string        `json:"__typename"`
}

type GetBasisGraph_GetGraphBasisGoals struct {
	IsIntensiveGraphAvailable bool        `json:"isIntensiveGraphAvailable"`
	GraphNodes                []GraphNode_GetGraphBasisGoals `json:"graphNodes"`
	Typename                  string      `json:"__typename"`
}

type GraphNode_GetGraphBasisGoals struct {
	GraphNodeID     string           `json:"graphNodeId"`
	NodeCode        string           `json:"nodeCode"`
	StudyDirections []StudyDirection_GetGraphBasisGoals `json:"studyDirections"`
	EntityType      string           `json:"entityType"`
	EntityID        string           `json:"entityId"`
	Goal            *Course_GetGraphBasisGoals          `json:"goal"`
	Course_GetGraphBasisGoals          *Course_GetGraphBasisGoals          `json:"course"`
	ParentNodeCodes []string         `json:"parentNodeCodes"`
	Typename        string           `json:"__typename"`
}

type Course_GetGraphBasisGoals struct {
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

type StudyDirection_GetGraphBasisGoals struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Color     string `json:"color"`
	TextColor string `json:"textColor"`
	Typename  string `json:"__typename"`
}

func (ctx *RequestContext) GetGraphBasisGoals(variables Variables_GetGraphBasisGoals) (Data_GetGraphBasisGoals, error) {
	request := gql.NewQueryRequest[Variables_GetGraphBasisGoals](
		"query getGraphBasisGoals($studentId: UUID!) {   student {     getBasisGraph(studentId: $studentId) {       isIntensiveGraphAvailable       graphNodes {         ...BasisGraphNode         __typename       }       __typename     }     __typename   } }  fragment BasisGraphNode on GraphNode {   graphNodeId   nodeCode   studyDirections {     id     name     color     textColor     __typename   }   entityType   entityId   goal {     goalExecutionType     projectState     projectName     projectDescription     projectPoints     projectDate     duration     isMandatory     __typename   }   course {     courseType     projectState     projectName     projectDescription     projectPoints     projectDate     duration     localCourseId     __typename   }   parentNodeCodes   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetGraphBasisGoals](ctx, request)
}