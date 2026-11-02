package main

import (
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestPublishedSchemaRequiresExactReviewedBytesAndVersion(t *testing.T) {
	raw := `{"info":{"version":"test-version"}}`
	digest := sha256.Sum256([]byte(raw))
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) { _, _ = w.Write([]byte(raw)) }))
	defer server.Close()
	source := schemaSource{URL: server.URL, SHA256: hex.EncodeToString(digest[:]), Version: "test-version"}
	result, err := fetchSchema(t.Context(), source, *server.Client())
	if err != nil || string(result) != raw {
		t.Fatalf("valid published source rejected: %v", err)
	}
	source.SHA256 = strings.Repeat("0", 64)
	if _, err := fetchSchema(t.Context(), source, *server.Client()); err == nil {
		t.Fatal("accepted changed schema bytes")
	}
	source.SHA256, source.Version = hex.EncodeToString(digest[:]), "different"
	if _, err := fetchSchema(t.Context(), source, *server.Client()); err == nil {
		t.Fatal("accepted wrong contract version")
	}
}

func TestPublishedSchemaRejectsRedirectsFailuresAndOversize(t *testing.T) {
	for _, status := range []int{302, 503, 200} {
		t.Run(http.StatusText(status), func(t *testing.T) {
			server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Location", "/other")
				w.WriteHeader(status)
				if status == 200 {
					_, _ = w.Write([]byte(strings.Repeat(" ", (4<<20)+1)))
				}
			}))
			defer server.Close()
			source := schemaSource{URL: server.URL, SHA256: strings.Repeat("0", 64), Version: "test"}
			if _, err := fetchSchema(t.Context(), source, *server.Client()); err == nil {
				t.Fatal("accepted invalid published response")
			}
		})
	}
}
