package requests

import "s21client/gql"

type Variables_GetAsapWidgets struct {
}


type Data_GetAsapWidgets struct {
	Student_GetAsapWidgets Student_GetAsapWidgets `json:"student"`
}

type Student_GetAsapWidgets struct {
	GetAsapWidgetList_GetAsapWidgets GetAsapWidgetList_GetAsapWidgets `json:"getAsapWidgetList"`
	Typename          string            `json:"__typename"`
}

type GetAsapWidgetList_GetAsapWidgets struct {
	WidgetList_GetAsapWidgets []WidgetList_GetAsapWidgets `json:"widgetList"`
	Typename   string       `json:"__typename"`
}

type WidgetList_GetAsapWidgets struct {
	ShortImg       string      `json:"shortImg"`
	ShortTitle     string      `json:"shortTitle"`
	ShortURL       interface{} `json:"shortUrl"`
	StartDate      string      `json:"startDate"`
	FinishDate     string      `json:"finishDate"`
	ShowFinishDate bool        `json:"showFinishDate"`
	FullTitle      string      `json:"fullTitle"`
	Text           string      `json:"text"`
	FullImgURL     string      `json:"fullImgUrl"`
	ADTTypeID      int64       `json:"adtTypeId"`
	ADTWidgetID    int64       `json:"adtWidgetId"`
	Typename       string      `json:"__typename"`
}

func (ctx *RequestContext) GetAsapWidgets(variables Variables_GetAsapWidgets) (Data_GetAsapWidgets, error) {
	request := gql.NewQueryRequest[Variables_GetAsapWidgets](
		"query getAsapWidgets {   student {     getAsapWidgetList {       widgetList {         ...AsapWidget         __typename       }       __typename     }     __typename   } }  fragment AsapWidget on AsapWidgetInfo {   shortImg   shortTitle   shortUrl   startDate   finishDate   showFinishDate   fullTitle   text   fullImgUrl   adtTypeId   adtWidgetId   __typename } ",
		variables,
	)

	return GqlRequest[Data_GetAsapWidgets](ctx, request)
}