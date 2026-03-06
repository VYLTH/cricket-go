package cricket

import (
	"context"
	"net/url"
)

// PulseClient provides access to unified price feeds across CEX and DEX venues.
type PulseClient struct {
	http *httpClient
}

// GetPrice returns price quotes for a token across all venues.
func (c *PulseClient) GetPrice(ctx context.Context, token string) ([]PriceQuote, error) {
	var result []PriceQuote
	err := c.http.get(ctx, "/pulse/price/"+url.PathEscape(token), &result)
	return result, err
}

// GetSpread returns cross-venue spread analysis for a token.
func (c *PulseClient) GetSpread(ctx context.Context, token string) (*SpreadAnalysis, error) {
	var result SpreadAnalysis
	err := c.http.get(ctx, "/pulse/spread/"+url.PathEscape(token), &result)
	return &result, err
}

// GetVenues lists all supported venues.
func (c *PulseClient) GetVenues(ctx context.Context) ([]Venue, error) {
	var result []Venue
	err := c.http.get(ctx, "/pulse/venues", &result)
	return result, err
}
