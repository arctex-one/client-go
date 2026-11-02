// Package transport implements the bounded public HTTP boundary shared by typed
// clients and gateways which must preserve exact response evidence.
package transport

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ErrInvalidArgument  = errors.New("invalid Arctex public request")
	ErrUnavailable      = errors.New("Arctex response unavailable; retain original submission")
	ErrResponseTooLarge = errors.New("Arctex response exceeds supported size")
)

const maxResponseBytes = 4 << 20

// Client is immutable after construction and safe for concurrent calls. Bearers
// are supplied per call so applications retain credential rotation ownership.
type Client struct {
	origin string
	http   http.Client
}

// New accepts an API origin without a path, query, fragment or user information.
// HTTP supports private development stands; production callers must supply an
// approved HTTPS origin and trusted TLS transport. Redirects are always refused.
// The caller's client is copied; a zero timeout becomes 30 seconds.
func New(origin string, client http.Client) (*Client, error) {
	u, err := url.Parse(origin)
	if err != nil || u == nil || u.Host == "" || u.User != nil || u.RawQuery != "" || u.ForceQuery || u.Fragment != "" || u.Opaque != "" || (u.Path != "" && u.Path != "/") || (u.Scheme != "https" && u.Scheme != "http") || client.Timeout < 0 {
		return nil, ErrInvalidArgument
	}
	client.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	if client.Timeout == 0 {
		client.Timeout = 30 * time.Second
	}
	return &Client{origin: strings.TrimRight(origin, "/"), http: client}, nil
}

func (c *Client) Origin() string { return c.origin }

// CloseIdleConnections releases idle connections owned by the supplied HTTP
// transport. It does not interrupt in-flight calls or prevent later calls.
func (c *Client) CloseIdleConnections() { c.http.CloseIdleConnections() }

// ResponseError retains bounded private evidence. Error never renders the body
// or credential. Truncated evidence must not be treated as a complete receipt.
type ResponseError struct {
	Status    int
	Body      []byte
	Truncated bool
	Code      string
}

func (e *ResponseError) Error() string { return "Arctex public HTTP request failed" }
func (e *ResponseError) Unwrap() error {
	if e.Truncated {
		return ErrResponseTooLarge
	}
	return nil
}

// RequestError exposes the cause to errors.Is/As without rendering request URLs,
// remote payloads or credential-provider errors in its public message.
type RequestError struct{ Cause error }

func (e *RequestError) Error() string        { return ErrUnavailable.Error() }
func (e *RequestError) Unwrap() error        { return e.Cause }
func (e *RequestError) Is(target error) bool { return target == ErrUnavailable }

// Call makes one application-level attempt. Persist a POST's key and original
// body before calling; uncertainty never authorizes a replacement submission.
// Body bytes and successful response bytes are returned without re-encoding.
func (c *Client) Call(ctx context.Context, bearer, method, path, key string, body json.RawMessage) (json.RawMessage, error) {
	u, err := url.ParseRequestURI(path)
	if err != nil || u.IsAbs() || u.Host != "" || !strings.HasPrefix(u.Path, "/v1/") || strings.ContainsAny(path, "\r\n#") || (method != http.MethodGet && method != http.MethodPost) || strings.TrimSpace(bearer) == "" || strings.ContainsAny(bearer, "\r\n") || (method == http.MethodPost && strings.TrimSpace(key) == "") {
		return nil, ErrInvalidArgument
	}
	for _, segment := range strings.Split(u.Path, "/") {
		if segment == "." || segment == ".." || strings.Contains(segment, "\\") {
			return nil, ErrInvalidArgument
		}
	}
	req, err := http.NewRequestWithContext(ctx, method, c.origin+path, bytes.NewReader(body))
	if err != nil {
		return nil, ErrInvalidArgument
	}
	req.Header.Set("Authorization", "Bearer "+bearer)
	if method == http.MethodPost {
		req.Header.Set("Idempotency-Key", key)
		req.Header.Set("Content-Type", "application/json")
	}
	response, err := c.http.Do(req)
	if err != nil {
		return nil, &RequestError{Cause: err}
	}
	defer func() { _ = response.Body.Close() }()
	raw, err := io.ReadAll(io.LimitReader(response.Body, maxResponseBytes+1))
	if err != nil {
		return nil, &RequestError{Cause: err}
	}
	truncated := len(raw) > maxResponseBytes
	if truncated {
		raw = raw[:maxResponseBytes]
	}
	if truncated || response.StatusCode < 200 || response.StatusCode >= 300 {
		var detail struct {
			Code string `json:"code"`
		}
		if !truncated {
			_ = json.Unmarshal(raw, &detail)
		}
		return nil, &ResponseError{Status: response.StatusCode, Body: raw, Truncated: truncated, Code: detail.Code}
	}
	return raw, nil
}
