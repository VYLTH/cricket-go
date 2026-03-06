package cricket

import (
	"context"
	"net/url"
)

// FireflyClient provides access to wallet intelligence and smart money signals.
type FireflyClient struct {
	http *httpClient
}

// GetWallet returns a wallet's intelligence profile.
func (c *FireflyClient) GetWallet(ctx context.Context, address string) (*WalletProfile, error) {
	var result WalletProfile
	err := c.http.get(ctx, "/firefly/wallet/"+url.PathEscape(address), &result)
	return &result, err
}

// GetSignals returns active smart money signals.
func (c *FireflyClient) GetSignals(ctx context.Context) ([]Signal, error) {
	var result []Signal
	err := c.http.get(ctx, "/firefly/signals", &result)
	return result, err
}

// GetLeaderboard returns the top wallets leaderboard.
func (c *FireflyClient) GetLeaderboard(ctx context.Context) ([]WalletProfile, error) {
	var result []WalletProfile
	err := c.http.get(ctx, "/firefly/leaderboard", &result)
	return result, err
}

// Track starts tracking a wallet.
func (c *FireflyClient) Track(ctx context.Context, address string) error {
	return c.http.post(ctx, "/firefly/track", map[string]string{"address": address}, nil)
}
