package messari

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Messari struct {
	baseUrl   string
	apiKey    string
	apiHeader string
}

func New(baseUrl string, apiKey string) Messari {
	return Messari{
		baseUrl:   baseUrl,
		apiKey:    apiKey,
		apiHeader: messariHeader,
	}
}

func (ms Messari) GetCoin() {

}

func (ms Messari) GetCoins() {

}

func (ms Messari) AssetProfile(queryVals url.Values) (*assetProfile, error) {
	res, err := ms.makeRequest(assetProfileEndpoint, queryVals)
	if err != nil {
		return nil, err
	}
	var data assetProfile
	err = json.Unmarshal(res, &data)
	return &data, err
}

func (ms Messari) AssetMetrics(queryVals url.Values) (*assetMetrics, error) {
	res, err := ms.makeRequest(assetMetricEndpoint, queryVals)
	if err != nil {
		return nil, err
	}
	var data assetMetrics
	err = json.Unmarshal(res, &data)
	return &data, err
}

func (ms Messari) AssetMarketData(queryVals url.Values) (*assetMarketData, error) {
	res, err := ms.makeRequest(assetMarketDataEndpoint, queryVals)
	if err != nil {
		return nil, err
	}
	var data assetMarketData
	err = json.Unmarshal(res, &data)
	return &data, err
}

func (ms Messari) makeRequest(endpoint string, queryVals url.Values) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(getMethod, ms.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add(ms.apiHeader, ms.apiKey)
	if queryVals != nil {
		req.URL.RawQuery = queryVals.Encode()
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

const (
	assetsEndpoint          = "/v2/assets"
	assetEndpoint           = "/v1/assets/%s"
	assetProfileEndpoint    = "/v2/assets/%s/profile"
	assetMetricEndpoint     = "/v1/assets/%s/metrics"
	assetMarketDataEndpoint = "/v1/assets/%s/metrics/market-data"
	getMethod               = "GET"
	messariHeader           = "x-messari-api-key"
)
