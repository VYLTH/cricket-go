package cricket

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// apiResponse is the standard envelope returned by Cricket API endpoints.
type apiResponse struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data,omitempty"`
	Error   *apiError       `json:"error,omitempty"`
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error represents an API error from Cricket.
type Error struct {
	Code    string
	Message string
	Status  int
}

func (e *Error) Error() string {
	return fmt.Sprintf("cricket: %s — %s (HTTP %d)", e.Code, e.Message, e.Status)
}

// httpClient is the internal HTTP transport used by all sub-clients.
type httpClient struct {
	baseURL string
	apiKey  string
	client  *http.Client
}

func newHTTPClient(baseURL, apiKey string, timeout time.Duration) *httpClient {
	return &httpClient{
		baseURL: strings.TrimRight(baseURL, "/"),
		apiKey:  apiKey,
		client:  &http.Client{Timeout: timeout},
	}
}

func (h *httpClient) get(ctx context.Context, path string, result any) error {
	return h.do(ctx, http.MethodGet, path, nil, result)
}

func (h *httpClient) post(ctx context.Context, path string, body, result any) error {
	return h.do(ctx, http.MethodPost, path, body, result)
}

func (h *httpClient) delete(ctx context.Context, path string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodDelete, h.baseURL+path, nil)
	if err != nil {
		return err
	}
	h.setHeaders(req)
	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return &Error{Code: "REQUEST_FAILED", Message: fmt.Sprintf("HTTP %d", resp.StatusCode), Status: resp.StatusCode}
	}
	return nil
}

func (h *httpClient) do(ctx context.Context, method, path string, body, result any) error {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, h.baseURL+path, bodyReader)
	if err != nil {
		return err
	}
	h.setHeaders(req)

	resp, err := h.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var envelope apiResponse
	if err := json.NewDecoder(resp.Body).Decode(&envelope); err != nil {
		return err
	}

	if !envelope.Success || envelope.Error != nil {
		code := "UNKNOWN_ERROR"
		msg := "unknown error"
		if envelope.Error != nil {
			code = envelope.Error.Code
			msg = envelope.Error.Message
		}
		return &Error{Code: code, Message: msg, Status: resp.StatusCode}
	}

	if result != nil && envelope.Data != nil {
		return json.Unmarshal(envelope.Data, result)
	}
	return nil
}

func (h *httpClient) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", h.apiKey)
}
