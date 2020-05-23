package main

import (
	"fmt"
	pr "github.com/tatrasoft/crypto-trader/pair_repo"
	"log"
	"net/http"
)

const (
	url = "https://api.eterbase.exchange/api/v1/"
	getMarkets = "markets"
	getTradingPair = "tickers/33/ticker"
)

func main()  {
	client := &http.Client{}

	pairs, err := pr.ParseSymbols(client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pairs)
	//
	//req, err := http.NewRequest("GET", fmt.Sprint(url, getMarkets), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("Accept", "application/json")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer res.Body.Close()
	//responseBody, err := ioutil.ReadAll(res.Body)
	//fmt.Println(res.StatusCode)
	//fmt.Println(string(responseBody))

	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("################################")
	//fmt.Println("################################")
	//fmt.Println("### WELCOME TO CRYPTO TRADER ###")
	//fmt.Println("################################")
	//fmt.Println("################################")
	//
	//fmt.Println("exit for exiting")
	//fmt.Println("get pairs - get list of pairs")
	//fmt.Println("get price {id} - get price of the pair for id")
	//fmt.Println("What to do?: ")
	//for  {
	//	text, err := reader.ReadString('\n')
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	switch text {
	//	case "exit\n":
	//		os.Exit(0)
	//	case "get pairs\n":
	//		fmt.Println("getting pairs")
	//	case "get price 33\n":
	//		fmt.Println("getting ETH/USDT")
	//	default:
	//		fmt.Println("Incorrect command try something else!")
	//	}
	//}
}
