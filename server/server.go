package server

import (
	"criptology/datautil"
	"criptology/logger"
    "html/template"
	"fmt"
	"io"
    "reflect"
	"log"
	"net/http"
	"os"
)

type ViewData struct {
    Title string
    Content string
}

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

func GetDataFromApi(url string) ([]uint8, error) {
    resp, err := http.Get(url)

    if err != nil {
        message := `"Unable to get data from source: %v", err`
        logger.WriteErrorLogFile(message)
        return nil, err
    }

    respData, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalf("Unable to read response data: %v", err)
        return nil, err
    }
    return respData, nil
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

func ServeCriprologyUint(data []uint8, field string) (*string, error) {
    isDataNil := data == nil
    if isDataNil {
        logger.WriteErrorLogFile("Cannot pass nil value!")
    }

    dataUint, err := datautil.ParseJsonFromAlpha(data)
    if err != nil {
        message := fmt.Sprintf("Cannot parse JSON: %v\n", err)
        logger.WriteErrorLogFile(message)
        return nil, err
    }
    reflectData := reflect.ValueOf(dataUint.RealtimeCurrencyExchangeRate).FieldByName(field)

    if !reflectData.IsValid() || reflectData.Interface() == "" {
        mess := fmt.Sprintf("Reflected data %v is not valid or \"\"", field)
        logger.WriteErrorLogFile(mess)
        return nil, err
    }

    resultData := fmt.Sprintf("%v", reflectData.Interface())

    return &resultData, nil
}

func injectDataIntoView(w http.ResponseWriter, link string, tip string, caption string, page string) {
    data, er := GetDataFromApi(link)
    if er != nil {
        mes := fmt.Sprintf("Cannot get data from API: %v", er)
        logger.WriteErrorLogFile(mes)
    }
    dataStr, greska := ServeCriprologyUint(data, tip)
    if greska != nil {
        mes := fmt.Sprintf("Cannot get data from API: %v", greska)
        logger.WriteErrorLogFile(mes)
        return
    }
    view := ViewData {
        Title: caption,
        Content: *dataStr,
    }
    // fmt.Printf("view: %v, %v", view.Title, view.Content)

    tmpl, err := template.ParseFiles(page)
    if err != nil {
        mes := fmt.Sprintf("Cannot parse file %v: %v", page, err)
        logger.WriteErrorLogFile(mes)
        return
    }
    // fmt.Printf("Type of tmpl: %T\n", tmpl)

    erro := tmpl.Execute(w, view)
    if erro != nil {
        mess := fmt.Sprintf("Cannot inject data into view: %v", erro)
        logger.WriteErrorLogFile(mess)
    }
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
    index := ReadFile("web/index.html")
    ServeCriprologyStr(w, index)
}

func handleFooter(w http.ResponseWriter, r *http.Request) {
    footer := ReadFile("web/views/footer.html")
    ServeCriprologyStr(w, footer)
}

func handler(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Path
    switch url {
        case "":
            errorHandler(w, r, http.StatusNotFound)
        case "/":
            handleIndex(w, r)
        case "/data":
            link := datautil.GetLink("crypto")
            injectDataIntoView(w, link, "FromCurrencyCode", "Naslov", "web/views/data.html")
        case "/footer":
            handleFooter(w, r)
    }
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}

func Server(url string) {
    port := datautil.GetConfig("port")
    http.HandleFunc(url, handler)
    err := http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatal(http.ListenAndServe(port, nil))
    }
}
