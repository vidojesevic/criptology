package main

import (
    "fmt"
    "cryptology/server"
    "cryptology/logger"
    "cryptology/datautil"
)

func main() {

    fmt.Print(server.Hello("Server successfylly started\n"))
    port := datautil.GetConfig("port")
    fmt.Printf("Listening port: %v\n", port)

    // data := server.GetDataFromApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo")
    logger.WriteAccessLogFile("Server succesfully started")

    server.Server()
}
