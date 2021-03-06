// Copyright (c) 2021, Ricard Bejarano. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license, which can be found in the LICENSE file.

package cloudflareanalytics

import (
	"fmt"
	"net/http"
	"net/url"

	schema "github.com/Elbandi/go-cloudflareanalytics/schemas/v4"
)

// Reusable structure holding the required metadata for talking to Cloudflare's
// GraphQL Analytics API.
//
// Client should be reused (not reinstantiated) across multiple requests.
type Client struct {
	// Cloudflare GraphQL Analytics API URL.
	// Exported to future-proof the structure against Cloudflare-side changes.
	Endpoint *url.URL

	httpClient *http.Client
	httpHeader *http.Header
}

// NewWithKey instantiates a Client prepared to perform requests over to the
// API using Cloudflare's email+key authentication method.
//
// Cloudflare recommends using token-based authentication instead. See
// NewWithAPIToken.
func NewWithKey(key string, email string) (*Client, error) {
	httpHeader := http.Header{}
	httpHeader.Set("X-Auth-Email", email)
	httpHeader.Set("X-Auth-Key", key)
	return newWithCustomHeader(&httpHeader)
}

// NewWithAPIToken instantiates a Client prepared to perform requests over to
// the API using Cloudflare's token authentication method.
//
// More info on how to generate a token for accessing the GraphQL Analytics API:
//   https://developers.cloudflare.com/analytics/graphql-api/getting-started/authentication/api-token-auth
func NewWithAPIToken(token string) (*Client, error) {
	httpHeader := http.Header{}
	httpHeader.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	return newWithCustomHeader(&httpHeader)
}

func newWithCustomHeader(httpHeader *http.Header) (*Client, error) {
	endpoint, err := url.Parse(schema.Endpoint)
	if err != nil {
		return nil, err
	}

	httpHeader.Set("Content-Type", "application/json")

	return &Client{
		Endpoint:   endpoint,
		httpClient: &http.Client{},
		httpHeader: httpHeader,
	}, nil
}
