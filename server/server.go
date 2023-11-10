package server

import (
	"io"
	"log"
	"net/http"
	"os"
    "cryptology/logger"
    "cryptology/datautil"
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
        message := `"Unable to get data from source: %v", err`
        logger.WriteErrorLogFile(message)
    }

    respData, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Unable to read response data: %v", err)
    }
    // fmt.Printf("Type of response data is %T\n", respData)
    return respData
}

func ServeCriprologyStr(w http.ResponseWriter, data string) {
    isDataEmpty := data == ""
    if isDataEmpty {
        logger.WriteErrorLogFile("Passing empty string!")
    }

    _, err := w.Write([]byte(data))
    if err != nil {
        log.Fatalf("Cannot write data to connection: %v", err)
    }
}

func ServeCriprologyUint(w http.ResponseWriter, data []uint8) {
    isDataNil := data == nil
    if isDataNil {
        logger.WriteErrorLogFile("Cannot pass nil value!")
    }
    encoded := datautil.ParseJsonFromAlpha(data)
    log.Printf("Data type: %T\n", encoded)

    // fix tomorrow
    // for _, item := range encoded[0] {
        // _, err := w.Write([]byte(item))
        _, err := w.Write([]byte(data))
        if err != nil {
            log.Fatalf("Cannot write data to connection: %v", err)
        }
    // }
}

func Server() {
    port := ":9000"
    head := ReadFile("web/views/head.html")
    footer := ReadFile("web/views/footer.html")
    // data := GetDataFromApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo")
    data := GetDataFromApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo")
    // wiki := GetDataFromApi("https://en.wikipedia.org/w/api.php?action=parse&page=Bitcoin&format=json")

    http.HandleFunc("/cryptology", func(w http.ResponseWriter, r *http.Request) {
        ServeCriprologyStr(w, head)
        ServeCriprologyStr(w, "<h1 class='text-center pt-3 pb-3'>Welcome to Cryptology</h1>")
        ServeCriprologyUint(w, data)
        // ServeCriprologyUint(w, wiki)
        ServeCriprologyStr(w, footer)
    })
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal(http.ListenAndServe(port, nil))
    }
}
