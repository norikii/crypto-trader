package pair_repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url = "https://api.eterbase.exchange/api/v1/"
	getMarkets = "markets"
	getTradingPair = "tickers/33/ticker"
)

//type symbols struct {
//	Symbols []pair
//}

type pair struct {
	ID int `json:"id"`
	Symbol string `json:"symbol"`
	Base string `json:"base"`
	Quote string `json:"quote"`

}

func getSymbols(client *http.Client) (*[]pair, error) {
	var symbols []pair

	req, err := http.NewRequest("GET", fmt.Sprint(url, getMarkets), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(responseBody, &symbols)
	if err != nil {
		return nil, err
	}

	return &symbols, nil
}

func ParseSymbols(client *http.Client) ([]string, error) {
	symbols, err := getSymbols(client)
	if err != nil {
		return nil, err
	}

	var pairs []string
	for _, symbol := range *symbols {
		pairs = append(pairs, fmt.Sprintf("%v - %s, ", symbol.ID, symbol.Symbol))
	}

	return pairs, nil
}

func FindSymbol(searchItem string, input []string) []string {
	var foundSymbols []string

	for _, symbol := range input {
		if symbol == searchItem {
			foundSymbols = append(foundSymbols, symbol)
		}
	}

	return foundSymbols
}

