package trivec

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strings"
	"text/template"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
)

const (
	libraryVersion     = "0.0.1"
	userAgent          = "go-trivec/" + libraryVersion
	mediaType          = "application/json"
	charset            = "utf-8"
	defaultEnvironment = "staging"
)

var (
	BaseURL = url.URL{
		Scheme: "https",
		Host:   "{{.service_id}}.trivecgateway.com",
		Path:   "",
	}
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client, subscriptionKey, serviceKey, serviceID, appID string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetSubscriptionKey(subscriptionKey)
	client.SetServiceKey(serviceKey)
	client.SetServiceID(serviceID)
	client.SetAppID(appID)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)
	client.SetEnvironment(defaultEnvironment)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL url.URL

	// credentials
	subscriptionKey string
	serviceKey      string
	serviceID       string
	appID           string
	environment     string

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	beforeRequestDo    BeforeRequestDoCallback
	onRequestCompleted RequestCompletionCallback
}

type BeforeRequestDoCallback func(*http.Client, *http.Request, interface{})

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) SubscriptionKey() string {
	return c.subscriptionKey
}

func (c *Client) SetSubscriptionKey(subscriptionKey string) {
	c.subscriptionKey = subscriptionKey
}

func (c Client) ServiceKey() string {
	return c.serviceKey
}

func (c *Client) SetServiceKey(serviceKey string) {
	c.serviceKey = serviceKey
}

func (c Client) ServiceID() string {
	return c.serviceID
}

func (c *Client) SetServiceID(serviceID string) {
	c.serviceID = serviceID
}

func (c Client) AppID() string {
	return c.appID
}

func (c *Client) SetAppID(appID string) {
	c.appID = appID
}

func (c Client) Environment() string {
	return c.environment
}

func (c *Client) SetEnvironment(environment string) {
	c.environment = environment
}

func (c Client) BaseURL() url.URL {
	return c.baseURL
}

func (c *Client) SetBaseURL(baseURL url.URL) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetBeforeRequestDo(fun BeforeRequestDoCallback) {
	c.beforeRequestDo = fun
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) url.URL {
	clientURL := c.BaseURL()
	clientURL.Host = strings.Replace(clientURL.Host, "{{.service_id}}", c.ServiceID(), -1)

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	q.Add("appid", c.AppID())
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	// params["app_id"] = c.AppID()
	err = tmpl.Execute(buf, params)
	if err != nil {
		log.Fatal(err)
	}

	clientURL.Path = buf.String()
	return clientURL
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if req.RequestBodyInterface() != nil {
		err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	r, err := http.NewRequest(req.Method(), req.URL().String(), buf)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())
	r.Header.Add("Subscription-Key", c.SubscriptionKey())
	r.Header.Add("ServiceKey", c.ServiceKey())
	if c.Environment() != "" {
		r.Header.Add("Environment", c.Environment())
	}

	return r, nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.beforeRequestDo != nil {
		c.beforeRequestDo(c.http, req, body)
	}

	if c.debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	if httpResp.ContentLength == 0 {
		return httpResp, nil
	}

	errResp := &ErrorResponse{Response: httpResp}
	httpErr := &HTTPError{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, body, errResp, httpErr)
	if err != nil {
		return httpResp, err
	}

	if errResp.Error() != "" {
		return httpResp, errResp
	}

	if httpErr.Error() != "" {
		return httpResp, httpErr
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)
		if c.disallowUnknownFields {
			dec.DisallowUnknownFields()
		}

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			errs = append(errs, err)
		}

	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert json to struct
	if len(data) != 0 {
		err = json.Unmarshal(data, &errorResponse)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if errorResponse.Error() != "" {
		return errorResponse
	}

	return nil
}

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	Errors struct {
		Date []string `json:"date"`
	} `json:"errors"`
	Type    string `json:"type"`
	Title   string `json:"title"`
	Status  int    `json:"status"`
	TraceID string `json:"traceId"`
}

func (r *ErrorResponse) Error() string {
	var errs *multierror.Error
	for _, m := range r.Errors.Date {
		errs = multierror.Append(errs, errors.New(m))
	}

	if errs == nil {
		return ""
	}
	return errs.Error()
}

// {
//   "$type": "System.Web.Http.HttpError, System.Web.Http",
//   "Message": "No HTTP resource was found that matches the request URI 'http://6676-92-70-210-65.ngrok-free.app/Catalog/PaymodeList.json?appid=d8f44c37d2d4465cbc03c55b45d6ab43'.",
//   "MessageDetail": "No action was found on the controller 'Catalog' that matches the name 'PaymodeList.json'."
// }

type HTTPError struct {
	// HTTP response that caused this error
	Response *http.Response

	Type          string `json:"$type"`
	Message       string `json:"Message"`
	MessageDetail string `json:"MessageDetail"`
}

func (e HTTPError) Error() string {
	if e.Type != "" || e.Message != "" || e.MessageDetail != "" {
		return fmt.Sprintf("%s: %s: %s", e.Type, e.Message, e.MessageDetail)
	}
	return ""
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != mediaType {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}
