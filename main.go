package main

import (
    "fmt"
    "criptology/server"
    "criptology/logger"
    "criptology/datautil"
)

func main() {

    mes := fmt.Sprintf("%v's server succedully started!\n", datautil.GetConfig("app"))
    fmt.Print(server.Hello(mes))
    port := datautil.GetConfig("port")
    fmt.Printf("Listening port: %s\n", port)

    // data := server.GetDataFromApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo")
    logger.WriteAccessLogFile("Server succesfully started")

    server.Server()
}
