package coingeko

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Coingeko struct {
	baseUrl string
}

func New(baseUrl string) Coingeko {
	return Coingeko{
		baseUrl,
	}
}

func (cg Coingeko) CoinsListBasic(queryVals url.Values) ([]CoinsListBasic, error) {
	res, err := cg.makeRequest(basicListEndpoint, queryVals)
	if err != nil {
		return []CoinsListBasic{}, err
	}
	var data []CoinsListBasic
	err = json.Unmarshal(res, &data)
	return data, err
}

func (cg Coingeko) CoinsList(queryVals url.Values) ([]CoinsList, error) {
	res, err := cg.makeRequest(listEndpoint, queryVals)
	if err != nil {
		return nil, err
	}

	var data []CoinsList
	err = json.Unmarshal(res, &data)
	return data, err
}

func (cg Coingeko) CoinInfo(queryVals url.Values, id string) (*CoinInfo, error) {
	res, err := cg.makeRequest(fmt.Sprintf(coinInfoEndpoint, id), queryVals)
	if err != nil {
		return nil, err
	}
	var data CoinInfo
	err = json.Unmarshal(res, &data)
	return &data, err
}

func (cg Coingeko) makeRequest(endpoint string, queryVals url.Values) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(getMethod, cg.baseUrl+endpoint, nil)
	if err != nil {
		return nil, err
	}
	if queryVals != nil {
		req.URL.RawQuery = queryVals.Encode()
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}
	body, err := ioutil.ReadAll(res.Body)
	return body, err
}

const (
	basicListEndpoint = "/coins/list"
	listEndpoint      = "/coins/markets"
	coinInfoEndpoint  = "/coins/%s"
	getMethod         = "GET"
)
