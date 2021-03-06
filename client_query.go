// Copyright (c) 2021, Ricard Bejarano. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license, which can be found in the LICENSE file.

package cloudflareanalytics

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	schema "github.com/ricardbejarano/go-cloudflareanalytics/schemas/v4"
)

type requestBody struct {
	Query     string            `json:"query"`
	Variables map[string]string `json:"variables"`
}

type responseBody struct {
	Data   schema.Data
	Errors schema.Errors
}

// Query sends `query` and `variables` over to the API, and (hopefully) returns
// the result as a *schema.Data structure.
//
// Errors are returned as []error since GraphQL may send more than one.
func (c *Client) Query(query string, variables map[string]string) (*schema.Data, []error) {
	requestBody, err := json.Marshal(requestBody{Query: query, Variables: variables})
	if err != nil {
		return nil, []error{err}
	}

	response, err := c.httpClient.Do(&http.Request{
		Method: http.MethodPost,
		URL:    c.Endpoint,
		Header: *c.httpHeader,
		Body:   ioutil.NopCloser(io.Reader(bytes.NewReader(requestBody))),
	})
	if err != nil {
		return nil, []error{err}
	}

	responseBodyBuffer := new(bytes.Buffer)
	if _, err := responseBodyBuffer.ReadFrom(response.Body); err != nil {
		return nil, []error{err}
	}

	var responseBody responseBody
	if err := json.Unmarshal(responseBodyBuffer.Bytes(), &responseBody); err != nil {
		return nil, []error{err}
	}

	return &responseBody.Data, responseBody.Errors.Slice()
}
