package main

import (
	"fmt"
	"io/ioutil"
)

func readAPIKeyFromFile() string {
	data, err := ioutil.ReadFile("apikey")
	if err != nil {
		fmt.Println("File reading error", err)
	}
	return string(data)
}

func getCurrencyExchangeURL(fromCurrency string, toCurrency string) string {
	url := "https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=" + fromCurrency + "&to_currency=" + toCurrency + "&apikey=" + readAPIKeyFromFile()
	return url
}
