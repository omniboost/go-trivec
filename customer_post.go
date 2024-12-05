package trivec

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-trivec/utils"
)

func (c *Client) NewCustomerPostRequest() CustomerPostRequest {
	r := CustomerPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerPostRequest struct {
	client      *Client
	queryParams *CustomerPostRequestQueryParams
	pathParams  *CustomerPostRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomerPostRequestBody
}

func (r CustomerPostRequest) NewQueryParams() *CustomerPostRequestQueryParams {
	return &CustomerPostRequestQueryParams{}
}

type CustomerPostRequestQueryParams struct {
}

func (p CustomerPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerPostRequest) QueryParams() *CustomerPostRequestQueryParams {
	return r.queryParams
}

func (r CustomerPostRequest) NewPathParams() *CustomerPostRequestPathParams {
	return &CustomerPostRequestPathParams{}
}

type CustomerPostRequestPathParams struct {
	ClientNr string `schema:"client_nr"`
}

func (p *CustomerPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"client_nr": p.ClientNr,
	}
}

func (r *CustomerPostRequest) PathParams() *CustomerPostRequestPathParams {
	return r.pathParams
}

func (r *CustomerPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerPostRequest) Method() string {
	return r.method
}

func (r CustomerPostRequest) NewRequestBody() CustomerPostRequestBody {
	return CustomerPostRequestBody{}
}

type CustomerPostRequestBody Customer

func (r *CustomerPostRequest) RequestBody() *CustomerPostRequestBody {
	return &r.requestBody
}

func (r *CustomerPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerPostRequest) SetRequestBody(body CustomerPostRequestBody) {
	r.requestBody = body
}

func (r *CustomerPostRequest) NewResponseBody() *CustomerPostRequestResponseBody {
	return &CustomerPostRequestResponseBody{}
}

type CustomerPostRequestResponseBody struct {
}

func (r *CustomerPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURLLiteAPI("Customer/key/{{.client_nr}}", r.PathParams())
	return &u
}

func (r *CustomerPostRequest) Do() (CustomerPostRequestResponseBody, error) {
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
