package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

	menu()

}

func CalculateCurrency(rate float64, ammount int) float64 {

	floatammount := float64(ammount)

	return floatammount * rate
}

func menu() {

	fmt.Println("#### Currency Excange ####")
	fmt.Print("Give Currency you want excange from : ")
	fromCurrency := readInput()
	fmt.Print("Give Currency you want excange to : ")
	toCurrency := readInput()
	fmt.Print("Ammount to excange : ")

	fmt.Println("#### Currency Excange Rate ####")
	fmt.Print(fromCurrency, toCurrency)

}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	return input
}

func loadData() {
	jsonFile, err := os.Open("data.json")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Opened data.json")

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var exchange ExchangeData

	json.Unmarshal(byteValue, &exchange)

	fmt.Println(exchange.ExchangeData.CurrencyFromCode)
	fmt.Println(exchange.ExchangeData.CurrencyToCode)
	exchangeFloat, _ := strconv.ParseFloat(exchange.ExchangeData.ExchangeRate, 64)
	fmt.Println(exchangeFloat)
	result := CalculateCurrency(exchangeFloat, 100)
	fmt.Println(result)
}
