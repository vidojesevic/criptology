package main

import (
    "fmt"
    "criptology/server"
)

func main() {
    fmt.Println("Hello, World!")

    fmt.Print(server.Hello("Hello from server\n"))

    // data := server.GetDataFromApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo")


    server.Server()
}
