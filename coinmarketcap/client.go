package coinmarketcap

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Coinmarketcap struct {
	baseUrl   string
	apiKey    string
	apiHeader string
}

func New(baseUrl string, apiKey string) Coinmarketcap {
	return Coinmarketcap{
		baseUrl:   baseUrl,
		apiKey:    apiKey,
		apiHeader: cmcHeader,
	}
}

func (cmc Coinmarketcap) MapData(queryVals url.Values) (*_map, error) {
	res, err := cmc.makeRequest(mapEndpoint, queryVals)
	if err != nil {
		return nil, err
	}
	var data _map
	err = json.Unmarshal(res, &data)
	return &data, err
}

func (cmc Coinmarketcap) MarketData(queryVals url.Values) (*market, error) {
	res, err := cmc.makeRequest(marketEndpoint, queryVals)
	if err != nil {
		return nil, err
	}
	var data market
	err = json.Unmarshal(res, &data)
	return &data, err
}

func (cmc Coinmarketcap) Metadata(queryVals url.Values) (*metadata, error) {
	res, err := cmc.makeRequest(metadataEndpoint, queryVals)
	if err != nil {
		return nil, err
	}
	var data metadata
	err = json.Unmarshal(res, &data)
	return &data, err
}

func (cmc Coinmarketcap) makeRequest(endpoint string, queryVals url.Values) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(getMethod, cmc.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add(cmc.apiHeader, cmc.apiKey)
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
	mapEndpoint      = "/v1/cryptocurrency/map"
	marketEndpoint   = "/v1/cryptocurrency/listings/latest"
	metadataEndpoint = "/v2/cryptocurrency/info"
	getMethod        = "GET"
	cmcHeader        = "X-CMC_PRO_API_KEY"
)
