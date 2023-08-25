package requests

import (
	"context"

	"github.com/s21toolkit/s21client/gql"

	"github.com/go-resty/resty/v2"
)

type RequestContext struct {
	request *resty.Request
	gqlUrl  string
}

func (ctx *RequestContext) Request() *resty.Request {
	return ctx.request
}

func NewRequestContext(request *resty.Request, gqlUrl string) *RequestContext {
	return &RequestContext{
		request: request,
		gqlUrl:  gqlUrl,
	}
}

func (context *RequestContext) SetContext(ctx context.Context) *RequestContext {
	context.Request().SetContext(ctx)

	return context
}

func GqlRequest[TData any, TVariables any](ctx *RequestContext, request gql.Request[TVariables]) (data TData, err error) {
	response, err := ctx.Request().SetBody(request).Post(ctx.gqlUrl)

	if err != nil {
		return
	}

	data, err = gql.ExtractResponseData[TData](response)

	return
}
