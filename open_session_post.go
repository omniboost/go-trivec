package trivec

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-trivec/utils"
)

func (c *Client) NewOpenSessionPostRequest() OpenSessionPostRequest {
	r := OpenSessionPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type OpenSessionPostRequest struct {
	client      *Client
	queryParams *OpenSessionPostRequestQueryParams
	pathParams  *OpenSessionPostRequestPathParams
	method      string
	headers     http.Header
	requestBody OpenSessionPostRequestBody
}

func (r OpenSessionPostRequest) NewQueryParams() *OpenSessionPostRequestQueryParams {
	return &OpenSessionPostRequestQueryParams{}
}

type OpenSessionPostRequestQueryParams struct {
}

func (p OpenSessionPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *OpenSessionPostRequest) QueryParams() *OpenSessionPostRequestQueryParams {
	return r.queryParams
}

func (r OpenSessionPostRequest) NewPathParams() *OpenSessionPostRequestPathParams {
	return &OpenSessionPostRequestPathParams{}
}

type OpenSessionPostRequestPathParams struct {
	ClientNr string `schema:"client_nr"`
}

func (p *OpenSessionPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"client_nr": p.ClientNr,
	}
}

func (r *OpenSessionPostRequest) PathParams() *OpenSessionPostRequestPathParams {
	return r.pathParams
}

func (r *OpenSessionPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *OpenSessionPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *OpenSessionPostRequest) Method() string {
	return r.method
}

func (r OpenSessionPostRequest) NewRequestBody() OpenSessionPostRequestBody {
	return OpenSessionPostRequestBody{}
}

type OpenSessionPostRequestBody Customer

func (r *OpenSessionPostRequest) RequestBody() *OpenSessionPostRequestBody {
	return &r.requestBody
}

func (r *OpenSessionPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *OpenSessionPostRequest) SetRequestBody(body OpenSessionPostRequestBody) {
	r.requestBody = body
}

func (r *OpenSessionPostRequest) NewResponseBody() *OpenSessionPostRequestResponseBody {
	return &OpenSessionPostRequestResponseBody{}
}

type OpenSessionPostRequestResponseBody struct {
}

func (r *OpenSessionPostRequest) URL() *url.URL {
	u := r.client.GetEndpointURLLiteAPI("OpenSession.json", r.PathParams())
	return &u
}

func (r *OpenSessionPostRequest) Do() (OpenSessionPostRequestResponseBody, error) {
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

