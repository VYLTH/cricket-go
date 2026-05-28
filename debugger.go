package kricket

import "context"

// DebuggerClient provides access to smart contract vulnerability scanning.
type DebuggerClient struct {
	http *httpClient
}

// Analyze analyzes source code for vulnerabilities and gas optimizations.
func (c *DebuggerClient) Analyze(ctx context.Context, source string, language Language) (*AnalysisResult, error) {
	var result AnalysisResult
	err := c.http.post(ctx, "/debugger/analyze", map[string]string{
		"source":   source,
		"language": string(language),
	}, &result)
	return &result, err
}
