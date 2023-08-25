package requests

import "github.com/s21toolkit/s21client/gql"

type Request_Variables_GetTasks struct {
	IDS []string `json:"ids"`
}


type Response_Data_GetTasks struct {
	Response_Student_GetTasks Response_Student_GetTasks `json:"student"`
}

type Response_Student_GetTasks struct {
	GetTasksByIDS []Response_GetTasksByID_GetTasks `json:"getTasksByIds"`
	Typename      string         `json:"__typename"`
}

type Response_GetTasksByID_GetTasks struct {
	ID       string `json:"id"`
	Response_Task_GetTasks     Response_Task_GetTasks   `json:"task"`
	Typename string `json:"__typename"`
}

type Response_Task_GetTasks struct {
	ID       string  `json:"id"`
	Response_Content_GetTasks  Response_Content_GetTasks `json:"content"`
	Typename string  `json:"__typename"`
}

type Response_Content_GetTasks struct {
	ID       string `json:"id"`
	Body     string `json:"body"`
	Typename string `json:"__typename"`
}

func (ctx *RequestContext) GetTasks(variables Request_Variables_GetTasks) (Response_Data_GetTasks, error) {
	request := gql.NewQueryRequest[Request_Variables_GetTasks](
		"query getTasks($ids: [ID!]!) {   student {     getTasksByIds(ids: $ids) {       ...StudentTaskInProject       __typename     }     __typename   } }  fragment StudentTaskInProject on StudentTask {   id   task {     id     content {       id       body       __typename     }     __typename   }   __typename } ",
		variables,
	)

	return GqlRequest[Response_Data_GetTasks](ctx, request)
}