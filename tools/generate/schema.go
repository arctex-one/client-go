package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

type schemaSource struct {
	URL     string `json:"url"`
	SHA256  string `json:"sha256"`
	Version string `json:"version"`
}

func loadSchema(lockPath string) ([]byte, error) {
	lock, err := os.ReadFile(lockPath)
	if err != nil {
		return nil, err
	}
	var source schemaSource
	if err := json.Unmarshal(lock, &source); err != nil {
		return nil, err
	}
	return fetchSchema(context.Background(), source, http.Client{Timeout: 30 * time.Second})
}

func fetchSchema(ctx context.Context, source schemaSource, client http.Client) ([]byte, error) {
	u, err := url.Parse(source.URL)
	digest, digestErr := hex.DecodeString(source.SHA256)
	if err != nil || u.Scheme != "https" || u.Host == "" || u.User != nil || u.Fragment != "" || digestErr != nil || len(digest) != sha256.Size || source.Version == "" {
		return nil, errors.New("schema lock requires an HTTPS URL, SHA256 and contract version")
	}
	client.CheckRedirect = func(*http.Request, []*http.Request) error { return http.ErrUseLastResponse }
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, source.URL, nil)
	if err != nil {
		return nil, err
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("fetch published schema: %w", err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("published schema HTTP status %d", response.StatusCode)
	}
	const limit = 4 << 20
	raw, err := io.ReadAll(io.LimitReader(response.Body, limit+1))
	if err != nil {
		return nil, err
	}
	if len(raw) > limit {
		return nil, errors.New("published schema exceeds supported size")
	}
	if actual := sha256.Sum256(raw); hex.EncodeToString(actual[:]) != source.SHA256 {
		return nil, errors.New("published schema checksum differs from reviewed lock")
	}
	var document struct {
		Info struct {
			Version string `json:"version"`
		} `json:"info"`
	}
	if err := json.Unmarshal(raw, &document); err != nil {
		return nil, err
	}
	if document.Info.Version != source.Version {
		return nil, errors.New("published schema version differs from reviewed lock")
	}
	return raw, nil
}
