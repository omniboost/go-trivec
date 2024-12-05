package trivec

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-trivec/utils"
)

func (c *Client) NewAccountInfoRequest() AccountInfoRequest {
	r := AccountInfoRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type AccountInfoRequest struct {
	client      *Client
	queryParams *AccountInfoRequestQueryParams
	pathParams  *AccountInfoRequestPathParams
	method      string
	headers     http.Header
	requestBody AccountInfoRequestBody
}

func (r AccountInfoRequest) NewQueryParams() *AccountInfoRequestQueryParams {
	return &AccountInfoRequestQueryParams{}
}

type AccountInfoRequestQueryParams struct{}

func (p AccountInfoRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *AccountInfoRequest) QueryParams() *AccountInfoRequestQueryParams {
	return r.queryParams
}

func (r AccountInfoRequest) NewPathParams() *AccountInfoRequestPathParams {
	return &AccountInfoRequestPathParams{}
}

type AccountInfoRequestPathParams struct {
	AccountKey string
}

func (p *AccountInfoRequestPathParams) Params() map[string]string {
	return map[string]string{
		"account_key": p.AccountKey,
	}
}

func (r *AccountInfoRequest) PathParams() *AccountInfoRequestPathParams {
	return r.pathParams
}

func (r *AccountInfoRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *AccountInfoRequest) SetMethod(method string) {
	r.method = method
}

func (r *AccountInfoRequest) Method() string {
	return r.method
}

func (r AccountInfoRequest) NewRequestBody() AccountInfoRequestBody {
	return AccountInfoRequestBody{}
}

type AccountInfoRequestBody struct {
}

func (r *AccountInfoRequest) RequestBody() *AccountInfoRequestBody {
	return nil
}

func (r *AccountInfoRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *AccountInfoRequest) SetRequestBody(body AccountInfoRequestBody) {
	r.requestBody = body
}

func (r *AccountInfoRequest) NewResponseBody() *AccountInfoRequestResponseBody {
	return &AccountInfoRequestResponseBody{}
}

type AccountInfoRequestResponseBody AccountInfo

func (r *AccountInfoRequest) URL() *url.URL {
	u := r.client.GetEndpointURLExportService("account/info.json/{{.account_key}}", r.PathParams())
	return &u
}

func (r *AccountInfoRequest) Do() (AccountInfoRequestResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(context.TODO(), r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
