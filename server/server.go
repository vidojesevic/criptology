package server

import (
	"io"
	"log"
	"net/http"
	"os"
    // "sync"
)

// var logMutex sync.Mutex

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
    isDataEmpty := data == ""
    if isDataEmpty {
        WriteErrorLogFile("Passing empty string!")
    }

    _, err := w.Write([]byte(data))
    if err != nil {
        log.Fatalf("Cannot write data to connection: %v", err)
    }
}

func ServeCriprologyUint(w http.ResponseWriter, data []uint8) {
    isDataNil := data == nil
    if isDataNil {
        WriteErrorLogFile("Cannot pass nil value!")
    }

    // w.Write returns (int, error)
    _, err := w.Write([]byte(data))
    if err != nil {
        log.Fatalf("Cannot write data to connection: %v", err)
    }
}

func WriteAccessLogFile(message string) {
    logPath := "server/log/access.log"

    logFile, err := OpenLogFile(logPath)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()

    logger := log.New(logFile, "[Access]", log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
    logger.Println(message)
}

func WriteErrorLogFile(message string) {
    logPath := "server/log/error.log"

    logFile, err := OpenLogFile(logPath)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()

    logger := log.New(logFile, "[Error]", log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
    logger.Println(message)
}

func OpenLogFile(path string) (*os.File, error) {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        logFile, err := os.Create(path)
        if err != nil {
            return nil, err
        }
        return logFile, nil
    }

    logFile, err := os.OpenFile(path, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }
    return logFile, nil
}

func Server() {
    head := ReadFile("web/views/head.html")
    footer := ReadFile("web/views/footer.html")
    data := GetDataFromApi("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=CNY&apikey=demo")

    http.HandleFunc("/cryptology", func(w http.ResponseWriter, r *http.Request) {
        ServeCriprologyStr(w, head)
        ServeCriprologyStr(w, "<h1 class='text-center pt-3 pb-3'>Welcome to Cryptology</h1>")
        ServeCriprologyUint(w, data)
        ServeCriprologyStr(w, footer)
    })
    err := http.ListenAndServe(":9000", nil)
    if err != nil {
        log.Fatal(http.ListenAndServe(":9000", nil))
    }
}
