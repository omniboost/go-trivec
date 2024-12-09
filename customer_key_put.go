package trivec

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-trivec/omitempty"
	"github.com/omniboost/go-trivec/utils"
)

func (c *Client) NewCustomerKeyPutRequest() CustomerKeyPutRequest {
	r := CustomerKeyPutRequest{
		client:  c,
		method:  http.MethodPut,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CustomerKeyPutRequest struct {
	client      *Client
	queryParams *CustomerKeyPutRequestQueryParams
	pathParams  *CustomerKeyPutRequestPathParams
	method      string
	headers     http.Header
	requestBody CustomerKeyPutRequestBody
}

func (r CustomerKeyPutRequest) NewQueryParams() *CustomerKeyPutRequestQueryParams {
	return &CustomerKeyPutRequestQueryParams{}
}

type CustomerKeyPutRequestQueryParams struct {
}

func (p CustomerKeyPutRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CustomerKeyPutRequest) QueryParams() *CustomerKeyPutRequestQueryParams {
	return r.queryParams
}

func (r CustomerKeyPutRequest) NewPathParams() *CustomerKeyPutRequestPathParams {
	return &CustomerKeyPutRequestPathParams{}
}

type CustomerKeyPutRequestPathParams struct {
	AccountKey string `schema:"account_key"`
}

func (p *CustomerKeyPutRequestPathParams) Params() map[string]string {
	return map[string]string{
		"account_key": p.AccountKey,
	}
}

func (r *CustomerKeyPutRequest) PathParams() *CustomerKeyPutRequestPathParams {
	return r.pathParams
}

func (r *CustomerKeyPutRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CustomerKeyPutRequest) SetMethod(method string) {
	r.method = method
}

func (r *CustomerKeyPutRequest) Method() string {
	return r.method
}

func (r CustomerKeyPutRequest) NewRequestBody() CustomerKeyPutRequestBody {
	return CustomerKeyPutRequestBody{}
}

type CustomerKeyPutRequestBody Customer

func (r CustomerKeyPutRequestBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *CustomerKeyPutRequest) RequestBody() *CustomerKeyPutRequestBody {
	return &r.requestBody
}

func (r *CustomerKeyPutRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CustomerKeyPutRequest) SetRequestBody(body CustomerKeyPutRequestBody) {
	r.requestBody = body
}

func (r *CustomerKeyPutRequest) NewResponseBody() *CustomerKeyPutRequestResponseBody {
	return &CustomerKeyPutRequestResponseBody{}
}

type CustomerKeyPutRequestResponseBody struct {
	Customer Customer `json:"customer"`
	Sucess   bool     `json:"success"`
}

func (r CustomerKeyPutRequestResponseBody) MarshalJSON() ([]byte, error) {
	return omitempty.MarshalJSON(r)
}

func (r *CustomerKeyPutRequest) URL() *url.URL {
	u := r.client.GetEndpointURLLiteAPI("customer/key/{{.account_key}}", r.PathParams())
	return &u
}

func (r *CustomerKeyPutRequest) Do() (CustomerKeyPutRequestResponseBody, error) {
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
