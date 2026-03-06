// Package cricket provides a Go client for the Cricket Protocol API.
package cricket

// --- Enums (typed string constants) ---

type Chain string

const (
	ChainSolana   Chain = "solana"
	ChainEthereum Chain = "ethereum"
	ChainBase     Chain = "base"
	ChainArbitrum Chain = "arbitrum"
)

type VenueType string

const (
	VenueTypeCex VenueType = "cex"
	VenueTypeDex VenueType = "dex"
)

type RiskRating string

const (
	RiskRatingLow      RiskRating = "low"
	RiskRatingModerate RiskRating = "moderate"
	RiskRatingHigh     RiskRating = "high"
	RiskRatingCritical RiskRating = "critical"
)

type Severity string

const (
	SeverityLow      Severity = "low"
	SeverityMedium   Severity = "medium"
	SeverityHigh     Severity = "high"
	SeverityCritical Severity = "critical"
)

type Tier string

const (
	TierScout  Tier = "scout"
	TierHunter Tier = "hunter"
	TierApex   Tier = "apex"
	TierColony Tier = "colony"
)

type WalletStyle string

const (
	WalletStyleEarlyAccumulator      WalletStyle = "early_accumulator"
	WalletStyleMomentumTrader        WalletStyle = "momentum_trader"
	WalletStyleSniper                WalletStyle = "sniper"
	WalletStyleWhaleMover            WalletStyle = "whale_mover"
	WalletStyleSmartContractDeployer WalletStyle = "smart_contract_deployer"
	WalletStyleUnknown               WalletStyle = "unknown"
)

type Language string

const (
	LanguageSolidity Language = "solidity"
	LanguageRust     Language = "rust"
	LanguageVyper    Language = "vyper"
)

type SignalType string

const (
	SignalTypeAccumulation SignalType = "accumulation"
	SignalTypeDivergence   SignalType = "divergence"
	SignalTypeExodus       SignalType = "exodus"
)

type SignalStrength string

const (
	SignalStrengthWeak     SignalStrength = "weak"
	SignalStrengthModerate SignalStrength = "moderate"
	SignalStrengthStrong   SignalStrength = "strong"
)

// --- Pulse types ---

type Venue struct {
	Name      string    `json:"name"`
	VenueType VenueType `json:"venue_type"`
	Chain     *Chain    `json:"chain,omitempty"`
}

type PriceQuote struct {
	Venue        string    `json:"venue"`
	VenueType    VenueType `json:"venue_type"`
	Bid          *float64  `json:"bid,omitempty"`
	Ask          *float64  `json:"ask,omitempty"`
	Mid          float64   `json:"mid"`
	Volume24h    float64   `json:"volume_24h"`
	LiquidityUSD *float64  `json:"liquidity_usd,omitempty"`
	Timestamp    string    `json:"timestamp"`
	LatencyUs    uint64    `json:"latency_us"`
}

type VenuePricePair struct {
	Venue string  `json:"venue"`
	Price float64 `json:"price"`
}

type SpreadAnalysis struct {
	BestBid    *VenuePricePair `json:"best_bid,omitempty"`
	BestAsk    *VenuePricePair `json:"best_ask,omitempty"`
	SpreadBps  float64         `json:"spread_bps"`
	VenueCount int             `json:"venue_count"`
}

// --- Mantis types ---

type RiskFlag struct {
	Check    string   `json:"check"`
	Severity Severity `json:"severity"`
	Status   string   `json:"status"`
	Detail   string   `json:"detail"`
}

type RiskScore struct {
	Score     int        `json:"score"`
	Rating    RiskRating `json:"rating"`
	Flags     []RiskFlag `json:"flags"`
	ScannedAt string     `json:"scanned_at"`
}

type ScanResult struct {
	TokenAddress            string     `json:"token_address"`
	MintAuthorityRevoked    bool       `json:"mint_authority_revoked"`
	FreezeAuthorityRevoked  bool       `json:"freeze_authority_revoked"`
	LpLocked                bool       `json:"lp_locked"`
	LpLockDurationDays      *int       `json:"lp_lock_duration_days,omitempty"`
	Top10HolderPct          float64    `json:"top_10_holder_pct"`
	DeployerWalletAgeDays   int        `json:"deployer_wallet_age_days"`
	UpgradeAuthorityRevoked bool       `json:"upgrade_authority_revoked"`
	MetadataMutable         bool       `json:"metadata_mutable"`
	Flags                   []RiskFlag `json:"flags"`
}

type ScanResponse struct {
	Scan      ScanResult `json:"scan"`
	RiskScore RiskScore  `json:"risk_score"`
}

// --- Firefly types ---

type WalletProfile struct {
	Address      string      `json:"address"`
	Chain        Chain       `json:"chain"`
	Score        int         `json:"score"`
	TotalTrades  int         `json:"total_trades"`
	WinRate      float64     `json:"win_rate"`
	AvgReturnPct float64     `json:"avg_return_pct"`
	TotalPnlUSD  float64     `json:"total_pnl_usd"`
	Style        WalletStyle `json:"style"`
	ActiveSince  string      `json:"active_since"`
}

type SignalEvidence struct {
	SmartWalletsCount int     `json:"smart_wallets_count"`
	AvgWalletScore    int     `json:"avg_wallet_score"`
	TotalVolumeUSD    float64 `json:"total_volume_usd"`
	TimeWindowHours   int     `json:"time_window_hours"`
}

type Signal struct {
	SignalType   SignalType     `json:"signal_type"`
	TokenAddress string         `json:"token_address"`
	TokenSymbol  string         `json:"token_symbol"`
	Strength     SignalStrength `json:"strength"`
	Evidence     SignalEvidence `json:"evidence"`
	DetectedAt   string         `json:"detected_at"`
}

// --- Debugger types ---

type SuggestedFix struct {
	Description string `json:"description"`
	CodeBefore  string `json:"code_before"`
	CodeAfter   string `json:"code_after"`
}

type Finding struct {
	Severity    Severity      `json:"severity"`
	FindingType string        `json:"finding_type"`
	Location    string        `json:"location"`
	Description string        `json:"description"`
	Fix         *SuggestedFix `json:"fix,omitempty"`
}

type GasOptimization struct {
	OptimizationType string       `json:"optimization_type"`
	Location         string       `json:"location"`
	EstimatedSavings string       `json:"estimated_savings"`
	Suggestion       SuggestedFix `json:"suggestion"`
}

type AnalysisResult struct {
	ScanID           string            `json:"scan_id"`
	Language         Language          `json:"language"`
	Findings         []Finding         `json:"findings"`
	GasOptimizations []GasOptimization `json:"gas_optimizations"`
	OverallRiskScore int               `json:"overall_risk_score"`
}

// --- Chirps types ---

type ChirpChannel struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Active    bool            `json:"active"`
	CreatedAt string          `json:"created_at"`
	UserID    *string         `json:"user_id,omitempty"`
	Config    map[string]any  `json:"config,omitempty"`
	Filters   map[string]any  `json:"filters,omitempty"`
}

type ChirpRecord struct {
	ID        string `json:"id"`
	ChannelID string `json:"channel_id"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	SentAt    string `json:"sent_at"`
}

// --- Watchlist ---

type WatchlistEntry struct {
	TokenAddress string     `json:"token_address"`
	AddedAt      string     `json:"added_at"`
	RiskScore    int        `json:"risk_score"`
	RiskRating   RiskRating `json:"risk_rating"`
}
