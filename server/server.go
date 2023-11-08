package server

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Hello(message string) string {
    return message
}

func ReadFile(path string) string {
    body, err := os.ReadFile(path)
    if err != nil {
        log.Fatalf("Unable to read file: %v", err)
    }

    return string(body)
}

func GetDataFromApi(url string) []uint8 {
    resp, err := http.Get(url)

    if err != nil {
        log.Fatalf("Unable to get data ftom source: %v", err)
    }

    respData, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Unable to read response data: %v", err)
    }
    // fmt.Printf("Type of response data is %T\n", respData)
    return respData
}

func ServeCriprologyStr(w http.ResponseWriter, data string) {
    // w.Write returns (int, error)
    if data == "" {
        log.Print("Passing empty string!")
    }

    _, err := w.Write([]byte(data))
    if err != nil {
        log.Fatalf("Cannot write data to connection: %v", err)
    }
}

func ServeCriprologyUint(w http.ResponseWriter, data []uint8) {
    // w.Write returns (int, error)
    if data == nil {
        log.Print("Error: Cannot pass nil value!")
    }

    _, err := w.Write([]byte(data))
    if err != nil {
        log.Fatalf("Cannot write data to connection: %v", err)
    }
}

func Server() {
    head := ReadFile("html/head.html")
    footer := ReadFile("html/footer.html")
    data := GetDataFromApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo")

    http.HandleFunc("/cryptology", func(w http.ResponseWriter, r *http.Request) {
        ServeCriprologyStr(w, head)
        ServeCriprologyStr(w, "")
        ServeCriprologyStr(w, "<h1 class='text-center pt-3 pb-3'>Welcome to Cryptology</h1>")
        ServeCriprologyUint(w, data)
        ServeCriprologyStr(w, footer)
    })
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal(http.ListenAndServe(":9000", nil))
    }
}
