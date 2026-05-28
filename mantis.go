package kricket

import (
	"context"
	"net/url"
)

// MantisClient provides access to rug-pull detection and token risk scoring.
type MantisClient struct {
	http *httpClient
}

// Scan runs a full security scan on a token.
func (c *MantisClient) Scan(ctx context.Context, token string) (*ScanResponse, error) {
	var result ScanResponse
	err := c.http.get(ctx, "/mantis/scan/"+url.PathEscape(token), &result)
	return &result, err
}

// GetScore returns only the risk score for a token.
func (c *MantisClient) GetScore(ctx context.Context, token string) (*RiskScore, error) {
	var result RiskScore
	err := c.http.get(ctx, "/mantis/score/"+url.PathEscape(token), &result)
	return &result, err
}

// AddToWatchlist adds tokens to the watchlist.
func (c *MantisClient) AddToWatchlist(ctx context.Context, tokens []string) error {
	return c.http.post(ctx, "/mantis/watchlist", map[string][]string{"tokens": tokens}, nil)
}

// GetWatchlist returns all watched tokens.
func (c *MantisClient) GetWatchlist(ctx context.Context) ([]WatchlistEntry, error) {
	var result []WatchlistEntry
	err := c.http.get(ctx, "/mantis/watchlist", &result)
	return result, err
}
