package coingeko

import "time"

//gomodifytags -file demo.go -struct Server -add-tags json
type CoinsListBasic struct {
	ID        string            `json:"id" bson:"id"`
	Symbol    string            `json:"symbol" bson:"symbol"`
	Name      string            `json:"name" bson:"name"`
	Platforms map[string]string `json:"platforms" bson:"platforms"`
}

type CoinsList struct {
	ID                           string                 `json:"id" bson:"id"`
	Symbol                       string                 `json:"symbol" bson:"symbol"`
	Name                         string                 `json:"name" bson:"name"`
	Image                        string                 `json:"image" bson:"image"`
	CurrentPrice                 float64                `json:"current_price" bson:"current_price"`
	MarketCap                    float64                `json:"market_cap" bson:"market_cap"`
	MarketCapRank                int                    `json:"market_cap_rank" bson:"market_cap_rank"`
	FullyDilutedValuation        float64                `json:"fully_diluted_valuation" bson:"fully_diluted_valuation"`
	TotalVolume                  float64                `json:"total_volume" bson:"total_volume"`
	High24H                      float64                `json:"high_24h" bson:"high_24_h"`
	Low24H                       float64                `json:"low_24h" bson:"low_24_h"`
	PriceChange24H               float64                `json:"price_change_24h" bson:"price_change_24_h"`
	PriceChangePercentage24H     float64                `json:"price_change_percentage_24h" bson:"price_change_percentage_24_h"`
	MarketCapChange24H           float64                `json:"market_cap_change_24h" bson:"market_cap_change_24_h"`
	MarketCapChangePercentage24H float64                `json:"market_cap_change_percentage_24h" bson:"market_cap_change_percentage_24_h"`
	CirculatingSupply            float64                `json:"circulating_supply" bson:"circulating_supply"`
	TotalSupply                  float64                `json:"total_supply" bson:"total_supply"`
	MaxSupply                    float64                `json:"max_supply" bson:"max_supply"`
	Ath                          float64                `json:"ath" bson:"ath"`
	AthChangePercentage          float64                `json:"ath_change_percentage" bson:"ath_change_percentage"`
	AthDate                      time.Time              `json:"ath_date" bson:"ath_date"`
	Atl                          float64                `json:"atl" bson:"atl"`
	AtlChangePercentage          float64                `json:"atl_change_percentage" bson:"atl_change_percentage"`
	AtlDate                      time.Time              `json:"atl_date" bson:"atl_date"`
	Roi                          map[string]interface{} `json:"roi" bson:"roi"`
	LastUpdated                  time.Time              `json:"last_updated" bson:"last_updated"`
}
type CoinInfo struct {
	ID                           string              `json:"id" bson:"id"`
	Symbol                       string              `json:"symbol" bson:"symbol"`
	Name                         string              `json:"name" bson:"name"`
	AssetPlatformID              string              `json:"asset_platform_id" bson:"asset_platform_id"`
	Platforms                    map[string]string   `json:"platforms" bson:"platforms"`
	BlockTimeInMinutes           int                 `json:"block_time_in_minutes" bson:"block_time_in_minutes"`
	HashingAlgorithm             string              `json:"hashing_algorithm" bson:"hashing_algorithm"`
	Categories                   []string            `json:"categories" bson:"categories"`
	PublicNotice                 string              `json:"public_notice" bson:"public_notice"`
	AdditionalNotices            []string            `json:"additional_notices" bson:"additional_notices"`
	Description                  description         `json:"description" bson:"description"`
	Links                        links               `json:"links" bson:"links"`
	Image                        image               `json:"image" bson:"image"`
	CountryOrigin                string              `json:"country_origin" bson:"country_origin"`
	GenesisDate                  time.Time           `json:"genesis_date" bson:"genesis_date"`
	SentimentVotesUpPercentage   float64             `json:"sentiment_votes_up_percentage" bson:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64             `json:"sentiment_votes_down_percentage" bson:"sentiment_votes_down_percentage"`
	IcoData                      icoData             `json:"ico_data" bson:"ico_data"`
	MarketCapRank                int                 `json:"market_cap_rank" bson:"market_cap_rank"`
	CoingeckoRank                int                 `json:"coingecko_rank" bson:"coingecko_rank"`
	CoingeckoScore               float64             `json:"coingecko_score" bson:"coingecko_score"`
	DeveloperScore               float64             `json:"developer_score" bson:"developer_score"`
	CommunityScore               float64             `json:"community_score" bson:"community_score"`
	LiquidityScore               float64             `json:"liquidity_score" bson:"liquidity_score"`
	PublicInterestScore          float64             `json:"public_interest_score" bson:"public_interest_score"`
	MarketData                   marketData          `json:"market_data" bson:"market_data"`
	CommunityData                communityData       `json:"community_data" bson:"community_data"`
	DeveloperData                developerData       `json:"developer_data" bson:"developer_data"`
	PublicInterestStats          publicInterestStats `json:"public_interest_stats" bson:"public_interest_stats"`
	StatusUpdates                interface{}         `json:"status_updates" bson:"status_updates"`
	LastUpdated                  time.Time           `json:"last_updated" bson:"last_updated"`
	Tickers                      []Tickers           `json:"tickers" bson:"tickers"`
	DocumentLastUpdated          time.Time           `json:"document_last_updated" bson:"document_last_updated"`
}
type description struct {
	En string `json:"en" bson:"en"`
}
type reposURL struct {
	Github    []string `json:"github" bson:"github"`
	Bitbucket []string `json:"bitbucket" bson:"bitbucket"`
}
type links struct {
	Homepage                    []string `json:"homepage" bson:"homepage"`
	BlockchainSite              []string `json:"blockchain_site" bson:"blockchain_site"`
	OfficialForumURL            []string `json:"official_forum_url" bson:"official_forum_url"`
	ChatURL                     []string `json:"chat_url" bson:"chat_url"`
	AnnouncementURL             []string `json:"announcement_url" bson:"announcement_url"`
	TwitterScreenName           string   `json:"twitter_screen_name" bson:"twitter_screen_name"`
	FacebookUsername            string   `json:"facebook_username" bson:"facebook_username"`
	BitcointalkThreadIdentifier int      `json:"bitcointalk_thread_identifier" bson:"bitcointalk_thread_identifier"`
	TelegramChannelIdentifier   string   `json:"telegram_channel_identifier" bson:"telegram_channel_identifier"`
	SubredditURL                string   `json:"subreddit_url" bson:"subreddit_url"`
	ReposURL                    reposURL `json:"repos_url" bson:"repos_url"`
}
type image struct {
	Thumb string `json:"thumb" bson:"thumb"`
	Small string `json:"small" bson:"small"`
	Large string `json:"large" bson:"large"`
}

type icoData struct {
	IcoStartDate            time.Time `json:"ico_start_date" bson:"ico_start_date"`
	IcoEndDate              time.Time `json:"ico_end_date" bson:"ico_end_date"`
	ShortDesc               string    `json:"short_desc" bson:"short_desc"`
	Description             string    `json:"description" bson:"description"`
	SoftcapCurrency         string    `json:"softcap_currency" bson:"softcap_currency"`
	HardcapCurrency         string    `json:"hardcap_currency" bson:"hardcap_currency"`
	TotalRaisedCurrency     string    `json:"total_raised_currency" bson:"total_raised_currency"`
	SoftcapAmount           string    `json:"softcap_amount" bson:"softcap_amount"`
	HardcapAmount           string    `json:"hardcap_amount" bson:"hardcap_amount"`
	TotalRaised             string    `json:"total_raised" bson:"total_raised"`
	QuotePreSaleCurrency    string    `json:"quote_pre_sale_currency" bson:"quote_pre_sale_currency"`
	BasePreSaleAmount       string    `json:"base_pre_sale_amount" bson:"base_pre_sale_amount"`
	QuotePreSaleAmount      string    `json:"quote_pre_sale_amount" bson:"quote_pre_sale_amount"`
	QuotePublicSaleCurrency string    `json:"quote_public_sale_currency" bson:"quote_public_sale_currency"`
	BasePublicSaleAmount    float64   `json:"base_public_sale_amount" bson:"base_public_sale_amount"`
	QuotePublicSaleAmount   float64   `json:"quote_public_sale_amount" bson:"quote_public_sale_amount"`
	AcceptingCurrencies     string    `json:"accepting_currencies" bson:"accepting_currencies"`
	CountryOrigin           string    `json:"country_origin" bson:"country_origin"`
	PreSaleStartDate        time.Time `json:"pre_sale_start_date" bson:"pre_sale_start_date"`
	PreSaleEndDate          time.Time `json:"pre_sale_end_date" bson:"pre_sale_end_date"`
	WhitelistURL            string    `json:"whitelist_url" bson:"whitelist_url"`
	WhitelistStartDate      string    `json:"whitelist_start_date" bson:"whitelist_start_date"`
	WhitelistEndDate        string    `json:"whitelist_end_date" bson:"whitelist_end_date"`
	BountyDetailURL         string    `json:"bounty_detail_url" bson:"bounty_detail_url"`
	AmountForSale           string    `json:"amount_for_sale" bson:"amount_for_sale"`
	KycRequired             bool      `json:"kyc_required" bson:"kyc_required"`
	WhitelistAvailable      bool      `json:"whitelist_available" bson:"whitelist_available"`
	PreSaleAvailable        bool      `json:"pre_sale_available" bson:"pre_sale_available"`
	PreSaleEnded            bool      `json:"pre_sale_ended" bson:"pre_sale_ended"`
}
type currentPrice struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type roi struct {
	Times      float64 `json:"times" bson:"times"`
	Currency   string  `json:"currency" bson:"currency"`
	Percentage float64 `json:"percentage" bson:"percentage"`
}
type ath struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type athChangePercentage struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type athDate struct {
	Usd time.Time `json:"usd" bson:"usd"`
}
type atl struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type atlChangePercentage struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type atlDate struct {
	Usd time.Time `json:"usd" bson:"usd"`
}
type marketCap struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type fullyDilutedValuation struct {
}
type totalVolume struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type high24H struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type low24H struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type priceChange24HInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type priceChangePercentage1HInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type priceChangePercentage24HInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type priceChangePercentage7DInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type priceChangePercentage14DInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type priceChangePercentage30DInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type priceChangePercentage60DInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type priceChangePercentage200DInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type priceChangePercentage1YInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type marketCapChange24HInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type marketCapChangePercentage24HInCurrency struct {
	Usd float64 `json:"usd" bson:"usd"`
}
type marketData struct {
	CurrentPrice                           currentPrice                           `json:"current_price" bson:"current_price"`
	Roi                                    roi                                    `json:"roi" bson:"roi"`
	Ath                                    ath                                    `json:"ath" bson:"ath"`
	AthChangePercentage                    athChangePercentage                    `json:"ath_change_percentage" bson:"ath_change_percentage"`
	AthDate                                athDate                                `json:"ath_date" bson:"ath_date"`
	Atl                                    atl                                    `json:"atl" bson:"atl"`
	AtlChangePercentage                    atlChangePercentage                    `json:"atl_change_percentage" bson:"atl_change_percentage"`
	AtlDate                                atlDate                                `json:"atl_date" bson:"atl_date"`
	MarketCap                              marketCap                              `json:"market_cap" bson:"market_cap"`
	MarketCapRank                          int                                    `json:"market_cap_rank" bson:"market_cap_rank"`
	FullyDilutedValuation                  fullyDilutedValuation                  `json:"fully_diluted_valuation" bson:"fully_diluted_valuation"`
	TotalVolume                            totalVolume                            `json:"total_volume" bson:"total_volume"`
	High24H                                high24H                                `json:"high_24h" bson:"high_24_h"`
	Low24H                                 low24H                                 `json:"low_24h" bson:"low_24_h"`
	PriceChange24H                         float64                                `json:"price_change_24h" bson:"price_change_24_h"`
	PriceChangePercentage24H               float64                                `json:"price_change_percentage_24h" bson:"price_change_percentage_24_h"`
	PriceChangePercentage7D                float64                                `json:"price_change_percentage_7d" bson:"price_change_percentage_7_d"`
	PriceChangePercentage14D               float64                                `json:"price_change_percentage_14d" bson:"price_change_percentage_14_d"`
	PriceChangePercentage30D               float64                                `json:"price_change_percentage_30d" bson:"price_change_percentage_30_d"`
	PriceChangePercentage60D               float64                                `json:"price_change_percentage_60d" bson:"price_change_percentage_60_d"`
	PriceChangePercentage200D              float64                                `json:"price_change_percentage_200d" bson:"price_change_percentage_200_d"`
	PriceChangePercentage1Y                float64                                `json:"price_change_percentage_1y" bson:"price_change_percentage_1_y"`
	MarketCapChange24H                     float64                                `json:"market_cap_change_24h" bson:"market_cap_change_24_h"`
	MarketCapChangePercentage24H           float64                                `json:"market_cap_change_percentage_24h" bson:"market_cap_change_percentage_24_h"`
	PriceChange24HInCurrency               priceChange24HInCurrency               `json:"price_change_24h_in_currency" bson:"price_change_24_h_in_currency"`
	PriceChangePercentage1HInCurrency      priceChangePercentage1HInCurrency      `json:"price_change_percentage_1h_in_currency" bson:"price_change_percentage_1_h_in_currency"`
	PriceChangePercentage24HInCurrency     priceChangePercentage24HInCurrency     `json:"price_change_percentage_24h_in_currency" bson:"price_change_percentage_24_h_in_currency"`
	PriceChangePercentage7DInCurrency      priceChangePercentage7DInCurrency      `json:"price_change_percentage_7d_in_currency" bson:"price_change_percentage_7_d_in_currency"`
	PriceChangePercentage14DInCurrency     priceChangePercentage14DInCurrency     `json:"price_change_percentage_14d_in_currency" bson:"price_change_percentage_14_d_in_currency"`
	PriceChangePercentage30DInCurrency     priceChangePercentage30DInCurrency     `json:"price_change_percentage_30d_in_currency" bson:"price_change_percentage_30_d_in_currency"`
	PriceChangePercentage60DInCurrency     priceChangePercentage60DInCurrency     `json:"price_change_percentage_60d_in_currency" bson:"price_change_percentage_60_d_in_currency"`
	PriceChangePercentage200DInCurrency    priceChangePercentage200DInCurrency    `json:"price_change_percentage_200d_in_currency" bson:"price_change_percentage_200_d_in_currency"`
	PriceChangePercentage1YInCurrency      priceChangePercentage1YInCurrency      `json:"price_change_percentage_1y_in_currency" bson:"price_change_percentage_1_y_in_currency"`
	MarketCapChange24HInCurrency           marketCapChange24HInCurrency           `json:"market_cap_change_24h_in_currency" bson:"market_cap_change_24_h_in_currency"`
	MarketCapChangePercentage24HInCurrency marketCapChangePercentage24HInCurrency `json:"market_cap_change_percentage_24h_in_currency" bson:"market_cap_change_percentage_24_h_in_currency"`
	TotalSupply                            float64                                `json:"total_supply" bson:"total_supply"`
	MaxSupply                              float64                                `json:"max_supply" bson:"max_supply"`
	CirculatingSupply                      float64                                `json:"circulating_supply" bson:"circulating_supply"`
	LastUpdated                            time.Time                              `json:"last_updated" bson:"last_updated"`
}
type communityData struct {
	FacebookLikes            int     `json:"facebook_likes" bson:"facebook_likes"`
	TwitterFollowers         int     `json:"twitter_followers" bson:"twitter_followers"`
	RedditAveragePosts48H    float64 `json:"reddit_average_posts_48h" bson:"reddit_average_posts_48_h"`
	RedditAverageComments48H float64 `json:"reddit_average_comments_48h" bson:"reddit_average_comments_48_h"`
	RedditSubscribers        int     `json:"reddit_subscribers" bson:"reddit_subscribers"`
	RedditAccountsActive48H  int     `json:"reddit_accounts_active_48h" bson:"reddit_accounts_active_48_h"`
	TelegramChannelUserCount int     `json:"telegram_channel_user_count" bson:"telegram_channel_user_count"`
}
type codeAdditionsDeletions4Weeks struct {
	Additions int `json:"additions" bson:"additions"`
	Deletions int `json:"deletions" bson:"deletions"`
}
type developerData struct {
	Forks                          int                          `json:"forks" bson:"forks"`
	Stars                          int                          `json:"stars" bson:"stars"`
	Subscribers                    int                          `json:"subscribers" bson:"subscribers"`
	TotalIssues                    int                          `json:"total_issues" bson:"total_issues"`
	ClosedIssues                   int                          `json:"closed_issues" bson:"closed_issues"`
	PullRequestsMerged             int                          `json:"pull_requests_merged" bson:"pull_requests_merged"`
	PullRequestContributors        int                          `json:"pull_request_contributors" bson:"pull_request_contributors"`
	CodeAdditionsDeletions4Weeks   codeAdditionsDeletions4Weeks `json:"code_additions_deletions_4_weeks" bson:"code_additions_deletions_4_weeks"`
	CommitCount4Weeks              int                          `json:"commit_count_4_weeks" bson:"commit_count_4_weeks"`
	Last4WeeksCommitActivitySeries []int                        `json:"last_4_weeks_commit_activity_series" bson:"last_4_weeks_commit_activity_series"`
}
type publicInterestStats struct {
	AlexaRank   int    `json:"alexa_rank" bson:"alexa_rank"`
	BingMatches string `json:"bing_matches" bson:"bing_matches"`
}
type Market struct {
	Name                string `json:"name" bson:"name"`
	Identifier          string `json:"identifier" bson:"identifier"`
	HasTradingIncentive bool   `json:"has_trading_incentive" bson:"has_trading_incentive"`
}
type convertedLast struct {
	Btc float64 `json:"btc" bson:"btc"`
	Eth float64 `json:"eth" bson:"eth"`
	Usd float64 `json:"usd" bson:"usd"`
}
type convertedVolume struct {
	Btc float64 `json:"btc" bson:"btc"`
	Eth float64 `json:"eth" bson:"eth"`
	Usd float64 `json:"usd" bson:"usd"`
}
type Tickers struct {
	Base                   string          `json:"base" bson:"base"`
	Target                 string          `json:"target" bson:"target"`
	Market                 Market          `json:"market" bson:"market"`
	Last                   float64         `json:"last" bson:"last"`
	Volume                 float64         `json:"volume" bson:"volume"`
	ConvertedLast          convertedLast   `json:"converted_last" bson:"converted_last"`
	ConvertedVolume        convertedVolume `json:"converted_volume" bson:"converted_volume"`
	TrustScore             string          `json:"trust_score" bson:"trust_score"`
	BidAskSpreadPercentage float64         `json:"bid_ask_spread_percentage" bson:"bid_ask_spread_percentage"`
	Timestamp              time.Time       `json:"timestamp" bson:"timestamp"`
	LastTradedAt           time.Time       `json:"last_traded_at" bson:"last_traded_at"`
	LastFetchAt            time.Time       `json:"last_fetch_at" bson:"last_fetch_at"`
	IsAnomaly              bool            `json:"is_anomaly" bson:"is_anomaly"`
	IsStale                bool            `json:"is_stale" bson:"is_stale"`
	TradeURL               string          `json:"trade_url" bson:"trade_url"`
	TokenInfoURL           string          `json:"token_info_url" bson:"token_info_url"`
	CoinID                 string          `json:"coin_id" bson:"coin_id"`
	TargetCoinID           string          `json:"target_coin_id,omitempty" bson:"target_coin_id"`
}
