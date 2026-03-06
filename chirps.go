package cricket

import (
	"context"
	"net/url"
)

// ChirpsClient provides access to notification channel management.
type ChirpsClient struct {
	http *httpClient
}

// CreateChannel creates a notification channel.
func (c *ChirpsClient) CreateChannel(ctx context.Context, channel map[string]any) (*ChirpChannel, error) {
	var result ChirpChannel
	err := c.http.post(ctx, "/chirps/channels", channel, &result)
	return &result, err
}

// ListChannels returns all notification channels.
func (c *ChirpsClient) ListChannels(ctx context.Context) ([]ChirpChannel, error) {
	var result []ChirpChannel
	err := c.http.get(ctx, "/chirps/channels", &result)
	return result, err
}

// DeleteChannel deletes a notification channel.
func (c *ChirpsClient) DeleteChannel(ctx context.Context, channelID string) error {
	return c.http.delete(ctx, "/chirps/channels/"+url.PathEscape(channelID))
}

// TestChannel sends a test chirp to verify channel configuration.
func (c *ChirpsClient) TestChannel(ctx context.Context, channelID string) error {
	return c.http.post(ctx, "/chirps/test/"+url.PathEscape(channelID), nil, nil)
}

// GetHistory returns notification delivery history.
func (c *ChirpsClient) GetHistory(ctx context.Context) ([]ChirpRecord, error) {
	var result []ChirpRecord
	err := c.http.get(ctx, "/chirps/history", &result)
	return result, err
}
