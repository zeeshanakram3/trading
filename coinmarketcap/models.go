package coinmarketcap

import "time"

// Common
type status struct {
	Timestamp    time.Time `json:"timestamp" bson:"timestamp"`
	ErrorCode    int       `json:"error_code" bson:"error_code"`
	ErrorMessage string    `json:"error_message" bson:"error_message"`
	Elapsed      int       `json:"elapsed" bson:"elapsed"`
	CreditCount  int       `json:"credit_count" bson:"credit_count"`
}

// Map data
type _map struct {
	Data   []data `json:"data" bson:"data"`
	Status status `json:"status" bson:"status"`
}
type mapData struct {
	ID                  int         `json:"id" bson:"id"`
	Name                string      `json:"name" bson:"name"`
	Symbol              string      `json:"symbol" bson:"symbol"`
	Slug                string      `json:"slug" bson:"slug"`
	IsActive            int         `json:"is_active" bson:"is_active"`
	FirstHistoricalData time.Time   `json:"first_historical_data" bson:"first_historical_data"`
	LastHistoricalData  time.Time   `json:"last_historical_data" bson:"last_historical_data"`
	Platform            interface{} `json:"platform" bson:"platform"`
}

// Coin latest market data
type market struct {
	Data   []marketData `json:"data" bson:"data"`
	Status status       `json:"status" bson:"status"`
}
type usd struct {
	Price            float64   `json:"price" bson:"price"`
	Volume24H        int64     `json:"volume_24h" bson:"volume_24_h"`
	PercentChange1H  float64   `json:"percent_change_1h" bson:"percent_change_1_h"`
	PercentChange24H float64   `json:"percent_change_24h" bson:"percent_change_24_h"`
	PercentChange7D  float64   `json:"percent_change_7d" bson:"percent_change_7_d"`
	MarketCap        int64     `json:"market_cap" bson:"market_cap"`
	LastUpdated      time.Time `json:"last_updated" bson:"last_updated"`
}
type btc struct {
	Price            int       `json:"price" bson:"price"`
	Volume24H        int       `json:"volume_24h" bson:"volume_24_h"`
	PercentChange1H  int       `json:"percent_change_1h" bson:"percent_change_1_h"`
	PercentChange24H int       `json:"percent_change_24h" bson:"percent_change_24_h"`
	PercentChange7D  int       `json:"percent_change_7d" bson:"percent_change_7_d"`
	MarketCap        int       `json:"market_cap" bson:"market_cap"`
	LastUpdated      time.Time `json:"last_updated" bson:"last_updated"`
}
type quote struct {
	USD usd `json:"USD" bson:"usd"`
	BTC btc `json:"BTC" bson:"btc"`
}
type marketData struct {
	ID                int         `json:"id" bson:"id"`
	Name              string      `json:"name" bson:"name"`
	Symbol            string      `json:"symbol" bson:"symbol"`
	Slug              string      `json:"slug" bson:"slug"`
	CmcRank           int         `json:"cmc_rank" bson:"cmc_rank"`
	NumMarketPairs    int         `json:"num_market_pairs" bson:"num_market_pairs"`
	CirculatingSupply int         `json:"circulating_supply" bson:"circulating_supply"`
	TotalSupply       int         `json:"total_supply" bson:"total_supply"`
	MaxSupply         int         `json:"max_supply" bson:"max_supply"`
	LastUpdated       time.Time   `json:"last_updated" bson:"last_updated"`
	DateAdded         time.Time   `json:"date_added" bson:"date_added"`
	Tags              []string    `json:"tags" bson:"tags"`
	Platform          interface{} `json:"platform" bson:"platform"`
	Quote             quote       `json:"quote" bson:"quote"`
}

// Coins meta data
type metadata struct {
	Data   data   `json:"data" bson:"data"`
	Status status `json:"status" bson:"status"`
}
type data struct {
	Coins map[string]coinMetaInfo `bson:"coins"`
}
type urls struct {
	Website      []string `json:"website" bson:"website"`
	TechnicalDoc []string `json:"technical_doc" bson:"technical_doc"`
	Twitter      []string `json:"twitter" bson:"twitter"`
	Reddit       []string `json:"reddit" bson:"reddit"`
	MessageBoard []string `json:"message_board" bson:"message_board"`
	Announcement []string `json:"announcement" bson:"announcement"`
	Chat         []string `json:"chat" bson:"chat"`
	Explorer     []string `json:"explorer" bson:"explorer"`
	SourceCode   []string `json:"source_code" bson:"source_code"`
}
type coinMetaInfo struct {
	Urls        urls        `json:"urls" bson:"urls"`
	Logo        string      `json:"logo" bson:"logo"`
	ID          int         `json:"id" bson:"id"`
	Name        string      `json:"name" bson:"name"`
	Symbol      string      `json:"symbol" bson:"symbol"`
	Slug        string      `json:"slug" bson:"slug"`
	Description string      `json:"description" bson:"description"`
	DateAdded   time.Time   `json:"date_added" bson:"date_added"`
	Tags        []string    `json:"tags" bson:"tags"`
	Platform    interface{} `json:"platform" bson:"platform"`
	Category    string      `json:"category" bson:"category"`
}
