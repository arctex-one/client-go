package transport

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestOriginAndRequestValidationBeforeNetwork(t *testing.T) {
	for _, origin := range []string{"", "ftp://example.com", "https://u:p@example.com", "https://example.com/v1", "https://example.com?", "https://example.com#fragment"} {
		if _, err := New(origin, http.Client{}); !errors.Is(err, ErrInvalidArgument) {
			t.Errorf("accepted invalid origin %q", origin)
		}
	}
	c, err := New("https://example.com", http.Client{})
	if err != nil {
		t.Fatal(err)
	}
	if c.http.Timeout != 30*time.Second {
		t.Fatal("unbounded default timeout")
	}
	for _, path := range []string{"https://other.example/v1/transfers", "//other.example/v1/transfers", "/v2/accounts", "/v1/accounts#x", "/v1/accounts\r\nx", "/v1/../admin", "/v1/%2e%2e/admin", "/v1/transfers/%5c.."} {
		if _, err := c.Call(t.Context(), "test", "GET", path, "", nil); !errors.Is(err, ErrInvalidArgument) {
			t.Errorf("accepted path %q: %v", path, err)
		}
	}
	if _, err := c.Call(t.Context(), "test", "POST", "/v1/transfers", "", nil); !errors.Is(err, ErrInvalidArgument) {
		t.Fatal("POST accepted without identity")
	}
	if _, err := c.Call(t.Context(), "", "GET", "/v1/accounts", "", nil); !errors.Is(err, ErrInvalidArgument) {
		t.Fatal("empty bearer accepted")
	}
}

func TestRawBoundaryRetainsBytesAndUsesCurrentCredential(t *testing.T) {
	body := []byte(` { "amount" : "1.0100" } `)
	response := []byte(` { "version" : "18446744073709551615" } `)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer "+r.URL.Query().Get("token") || r.Header.Get("Idempotency-Key") != "original" {
			t.Error("credential/identity changed")
		}
		raw, _ := io.ReadAll(r.Body)
		if !bytes.Equal(raw, body) {
			t.Error("request bytes changed")
		}
		_, _ = w.Write(response)
	}))
	defer server.Close()
	c, _ := New(server.URL, *server.Client())
	for _, token := range []string{"first", "rotated"} {
		raw, err := c.Call(t.Context(), token, "POST", "/v1/quotes?token="+token, "original", body)
		if err != nil || !bytes.Equal(raw, response) {
			t.Fatalf("raw evidence changed: %v", err)
		}
	}
}

func TestRequestErrorRetainsCancellationWithoutRenderingURL(t *testing.T) {
	c, _ := New("http://127.0.0.1:1", http.Client{})
	ctx, cancel := context.WithCancel(t.Context())
	cancel()
	_, err := c.Call(ctx, "private-bearer", "GET", "/v1/accounts?private=value", "", nil)
	if !errors.Is(err, ErrUnavailable) || !errors.Is(err, context.Canceled) {
		t.Fatalf("cause lost: %v", err)
	}
	if strings.Contains(err.Error(), "private") {
		t.Fatal("public error exposes private request")
	}
}

func TestResponseErrorPreservesCodeWithoutRenderingEvidence(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(503)
		_, _ = io.WriteString(w, `{"code":"FUTURE_CODE","detail":"private-evidence"}`)
	}))
	defer server.Close()
	c, _ := New(server.URL, *server.Client())
	_, err := c.Call(t.Context(), "private-bearer", "POST", "/v1/transfers", "original", []byte(`{}`))
	var response *ResponseError
	if !errors.As(err, &response) || response.Code != "FUTURE_CODE" || !bytes.Contains(response.Body, []byte("private-evidence")) {
		t.Fatalf("evidence lost: %v", err)
	}
	if strings.Contains(err.Error(), "private") {
		t.Fatal("error renders private evidence")
	}
}

type trackedTransport struct{ closed bool }

func (t *trackedTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, context.Canceled
}
func (t *trackedTransport) CloseIdleConnections() { t.closed = true }

func TestCloseReleasesSuppliedIdleConnections(t *testing.T) {
	pool := &trackedTransport{}
	c, err := New("https://example.com", http.Client{Transport: pool})
	if err != nil {
		t.Fatal(err)
	}
	c.CloseIdleConnections()
	if !pool.closed {
		t.Fatal("shutdown did not release the supplied transport")
	}
}
