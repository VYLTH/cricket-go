package cricket

import "time"

const defaultBaseURL = "https://api.cricket.vylth.com/api/v1"
const defaultTimeout = 30 * time.Second

// Config holds configuration for the Cricket client.
type Config struct {
	BaseURL string
	APIKey  string
	Timeout time.Duration
}

// Client is the unified entry point for all Cricket Protocol APIs.
type Client struct {
	Pulse    *PulseClient
	Mantis   *MantisClient
	Firefly  *FireflyClient
	Chirps   *ChirpsClient
	Debugger *DebuggerClient
}

// New creates a new Cricket client with the given configuration.
func New(cfg Config) *Client {
	if cfg.BaseURL == "" {
		cfg.BaseURL = defaultBaseURL
	}
	if cfg.Timeout == 0 {
		cfg.Timeout = defaultTimeout
	}

	h := newHTTPClient(cfg.BaseURL, cfg.APIKey, cfg.Timeout)

	return &Client{
		Pulse:    &PulseClient{http: h},
		Mantis:   &MantisClient{http: h},
		Firefly:  &FireflyClient{http: h},
		Chirps:   &ChirpsClient{http: h},
		Debugger: &DebuggerClient{http: h},
	}
}
