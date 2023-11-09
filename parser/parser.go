package parser

import (
	"encoding/json"
	"fmt"
)

type ExchangeRate struct {
    RealtimeCurrencyExchangeRate struct {
        FromCurrencyCode string `json:"1. From_Currency Code"`
        FromCurrencyName string `json:"2. From_Currency Name"`
        ToCurrencyCode     string `json:"3. To_Currency Code"`
		ToCurrencyName     string `json:"4. To_Currency Name"`
		ExchangeRate       string `json:"5. Exchange Rate"`
		LastRefreshed      string `json:"6. Last Refreshed"`
		TimeZone           string `json:"7. Time Zone"`
		BidPrice           string `json:"8. Bid Price"`
		AskPrice           string `json:"9. Ask Price"`
    } `json:"Realtime Currency Exchange Rate"`
}

func ParseJsonFromAlpha(data []uint8) (ExchangeRate) {
    fmt.Println("Parsing some JSON")

    var exRate ExchangeRate
    err := json.Unmarshal([]byte(data), &exRate)
    if err != nil {
        fmt.Print("Opa")
    }

    return exRate
}
