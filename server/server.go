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
    Name string 
    FromCurrencyName string
    ToCurrencyCode string
    ToCurrencyName string
    ExchangeRate string
    LastRefreshed string
    TimeZone string
    Price string
    AskPrice string
}

type WikiData struct {
    Paragraph1 string
    Paragraph2 string
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

func DataCriprologyUint(data []uint8, field string) (*string, error) {
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
    dataStr, greska := DataCriprologyUint(data, tip)
    if greska != nil {
        mes := fmt.Sprintf("Cannot get data from API: %v", greska)
        logger.WriteErrorLogFile(mes)
        return
    }
    price, danger := DataCriprologyUint(data, "BidPrice") 
    if danger != nil {
        mes := fmt.Sprintf("Cannot get data from API: %v", danger)
        logger.WriteErrorLogFile(mes)
        return
    }
    view := ViewData {
        Title: caption,
        Name: *dataStr,
        Price: *price,
    }
    // fmt.Printf("view: %v, %v, %v", view.Title, view.Content, view.Price)

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
    index, err := os.ReadFile("public/index.html")
    if err != nil {
        mess := fmt.Sprintf("Cannot load %v: %v", index, err)
        logger.WriteErrorLogFile(mess)
    }
    w.Header().Set("Content-type", "text/html")
    w.WriteHeader(http.StatusOK)
    _, er := w.Write(index)
    if er != nil {
        mess := fmt.Sprintf("Cannot serve content: %v", er)
        logger.WriteErrorLogFile(mess)
    }
}

func handlePage(w http.ResponseWriter, r *http.Request, page string) {
    path := "public/views" + page + ".html"
    view := ReadFile(path)
    ServeCriprologyStr(w, view)
}

func handleMessage(w http.ResponseWriter, r *http.Request) {
    insert := ReadFile("public/views/insert.html")
    ServeCriprologyStr(w, insert)
}

func handleImage(w http.ResponseWriter, r *http.Request, img string) {
    image, err := os.ReadFile(img)
    if err != nil {
        mess := fmt.Sprintf("Cannot load %v: %v", image, err)
        logger.WriteErrorLogFile(mess)
    }
    w.Header().Set("Content-type", "image/webp")
    w.WriteHeader(http.StatusOK)
    _, er := w.Write(image)
    if er != nil {
        mess := fmt.Sprintf("Cannot serve content: %v", er)
        logger.WriteErrorLogFile(mess)
    }
}

func loadStatic(w http.ResponseWriter, r *http.Request, path string) {
    http.ServeFile(w, r, path)
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
            injectDataIntoView(w, link, "FromCurrencyCode", "Naslov", "public/views/data.html")
        case "/footer":
            handlePage(w, r, "/footer")
        case "/message":
            handleMessage(w, r)
        case "/login":
            handlePage(w, r, "/login")
        case "/images":
            handleImage(w, r, "/images/cover.webp")
        case "/public/css/style.css":
            loadStatic(w, r, "public/css/style.css")
        case "/public/scripts/main.js":
            loadStatic(w, r, "public/scripts/main.js")
        case "/public/images/cover.webp":
            handleImage(w, r, "public/images/cover.webp")
        default:
            errorHandler(w, r, 404)
    }
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "Desi poso sine!")
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
