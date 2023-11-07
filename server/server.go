package server

import (
    "fmt"
    "net/http"
    "os"
    "io"
    // "html"
    "log"
)

func Hello(message string) {
    fmt.Println(message)
}

func readFile(path string) string {
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
    // fmt.Println(string(respData))
    fmt.Printf("Type of response data is %T\n", respData)
    return respData
}

func Server() {
    head := readFile("html/head.html")
    data := GetDataFromApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo")
    http.HandleFunc("/cryptology", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte(head))
        w.Write([]byte("<h1 class='text-center'>Welcome to Cryptology!</h1>"))
        w.Write([]byte(data))
    })
    http.ListenAndServe(":9000", nil)

    //
    // http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    // })
    //
    // log.Fatal(http.ListenAndServe(":9000", nil))
}
