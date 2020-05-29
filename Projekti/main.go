package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ExchangeData struct {
	ExchangeData Currency `json:"Realtime Currency Exchange Rate"`
}

type Currency struct {
	CurrencyFromCode string `json:"1. From_Currency Code"`
	CurrencyToCode   string `json:"3. To_Currency Code"`
	CurrencyFromName string `json:"2. From_Currency Name"`
	CurrencyToName   string `json:"4. To_Currency Name"`
	ExchangeRate     string `json:"5. Exchange Rate"`
}

func main() {

	fmt.Print("Currency to exchange from: ")
	currencyFrom := readInput()
	fmt.Print("Currency to exchange to: ")
	currencyTo := readInput()
	fmt.Print("Currency ammount to exchange: ")
	ammount := readInput()

	loadJSON(currencyFrom, currencyTo, ammount)
}

//Helper function to read user input
func readInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func loadJSON(fromCurrency string, toCurrency string, ammount string) {
	url := getCurrencyExchangeURL(fromCurrency, toCurrency)
	var exchange ExchangeData

	//just for test to see api url generated
	//fmt.Println(url)

	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {

		}
		//Just for testing to see what i get
		//bodyString := string(bodyBytes)
		//fmt.Println(bodyString)

		json.Unmarshal(bodyBytes, &exchange)
		floatammount, _ := strconv.ParseFloat(ammount, 64)
		exchangeFloat, _ := strconv.ParseFloat(exchange.ExchangeData.ExchangeRate, 64)

		result := floatammount * exchangeFloat
		fmt.Println("### Result ###")
		fmt.Println(result, exchange.ExchangeData.CurrencyToName)
	}

}
