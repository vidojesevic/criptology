package main

import (
    "fmt"
    "criptology/server"
)

func main() {
    fmt.Println("Hello, World!")

    server.Hello("Hello from server")

    // data := server.GetDataFromApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo")


    server.Server()
}
