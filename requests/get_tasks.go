package requests

import "s21client/gql"

type Variables_GetTasks struct {
	IDS []string `json:"ids"`
}


type Data_GetTasks struct {
	Student_GetTasks Student_GetTasks `json:"student"`
}

type Student_GetTasks struct {
	GetTasksByIDS []GetTasksByID_GetTasks `json:"getTasksByIds"`
	Typename      string         `json:"__typename"`
}

type GetTasksByID_GetTasks struct {
	ID       string `json:"id"`
	Task_GetTasks     Task_GetTasks   `json:"task"`
	Typename string `json:"__typename"`
}

type Task_GetTasks struct {
	ID       string  `json:"id"`
	Content_GetTasks  Content_GetTasks `json:"content"`
	Typename string  `json:"__typename"`
}

type Content_GetTasks struct {
	ID       string `json:"id"`
	Body     string `json:"body"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetTasks(variables Variables_GetTasks) (Data_GetTasks, error) {
	request := gql.NewQueryRequest[Variables_GetTasks](
		"query getTasks($ids: [ID!]!) {   student {     getTasksByIds(ids: $ids) {       ...StudentTaskInProject       __typename     }     __typename   } }  fragment StudentTaskInProject on StudentTask {   id   task {     id     content {       id       body       __typename     }     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetTasks](ctx, request)
}