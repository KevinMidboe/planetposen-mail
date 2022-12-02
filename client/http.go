// Package client contains a HTTP client.
package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	opentracing "github.com/opentracing/opentracing-go"
)

// Parameters provides the parameters used when creating a new HTTP client.
type Parameters struct {
	Timeout *time.Duration
}

// NewHTTPClient instantiates a new HTTPClient based on provided parameters.
func NewHTTPClient(parameters Parameters) HTTPClient {
	if parameters.Timeout == nil {
		timeout := 1 * time.Second
		parameters.Timeout = &timeout
	}

	client := &http.Client{
		Timeout: *parameters.Timeout,
	}

	return HTTPClient{client}
}

// HTTPRequestData contains the request data.
type HTTPRequestData struct {
	Method      string
	URL         string
	Headers     map[string]string
	PostPayload []byte
	GetPayload  *url.Values
}

// HTTPClient contains the HTTP client.
type HTTPClient struct {
	*http.Client
}

// HTTPStatusCodeError is an error that occurs when receiving an unexpected status code (>= 400).
type HTTPStatusCodeError struct {
	URL        string
	StatusCode int
	Message    string
}

// Error return an error string.
func (e HTTPStatusCodeError) Error() string {
	return fmt.Sprintf("Error response from %s, got status: %d", e.URL, e.StatusCode)
}

// RequestBytes does the actual HTTP request.
// Returns a slice of bytes or an error.
func (client *HTTPClient) RequestBytes(ctx context.Context, reqData HTTPRequestData) ([]byte, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "RequestBytes")
	defer span.Finish()

	r, err := client.request(ctx, reqData)

	if err != nil {
		return nil, err
	}

	defer r.Body.Close()

	if r.StatusCode >= 400 {
		resp, _ := ioutil.ReadAll(r.Body)
		message := string(resp)
		span.SetTag("error", true)
		span.LogKV("message", fmt.Errorf("error making request to %s, got error: %s", reqData.URL, message))
		return nil, HTTPStatusCodeError{
			URL:        reqData.URL,
			StatusCode: r.StatusCode,
			Message:    message,
		}
	}

	return ioutil.ReadAll(r.Body)
}

func (client *HTTPClient) request(ctx context.Context, reqData HTTPRequestData) (*http.Response, error) {
	var req *http.Request
	var err error

	if reqData.Method == http.MethodPost {
		req, err = http.NewRequest(reqData.Method, reqData.URL, bytes.NewBuffer(reqData.PostPayload))
	} else {
		req, err = http.NewRequest(reqData.Method, reqData.URL, nil)
	}

	if err != nil {
		return nil, err
	}

	if reqData.GetPayload != nil {
		req.URL.RawQuery = reqData.GetPayload.Encode()
	}

	span := opentracing.SpanFromContext(ctx)

	if span != nil {
		opentracing.GlobalTracer().Inject(
			span.Context(),
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(req.Header),
		)
	}

	for k, v := range reqData.Headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("User-Agent", "planetposen-mail")

	resp, err := client.Do(req)

	if err != nil {
		if reqData.Method == http.MethodPost {
			return resp, fmt.Errorf("Error making request: %v. Body: %s", err, reqData.PostPayload)
		}

		return resp, fmt.Errorf("Error making request: %v. Query: %v", err, req.URL.RawQuery)
	}

	return resp, nil
}