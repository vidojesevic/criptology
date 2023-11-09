package server

import (
	"bytes"
	"criptology/server"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestServeCriptologyStr(t *testing.T) {
    // For now, file is created when the new request is made
    // t.Run("When parameter is \"\"", func(t *testing.T) {
    //     data := ""
    //     w := &MockResponseWriter{}
    //     w.written = nil
    //
    //     want := "Passing empty string!"
    //     server.ServeCriprologyStr(w, data)
    //
    //     logger, err := readLog("server/log/error.log")
    //     if err != nil {
    //         log.Fatal("Something went wrong!")
    //     }
    //
    //     if strings.Contains(logger, want) {
    //         t.Errorf("Expected %v in log file, got %v\n", want, logger)
    //     }
    //
    // })

    t.Run("When parameter is not \"\"", func(t *testing.T) {
        want := "<h1>Hello, World!</h1>"
        w := &MockResponseWriter{}
        w.written = nil

        var logBuffer bytes.Buffer
        log.SetOutput(&logBuffer)

        server.ServeCriprologyStr(w, want)

        log.SetOutput(os.Stderr)

        logOut := logBuffer.String()
        if logOut != "" {
            t.Errorf("Expected nmo log message, but got: %v", logOut)
        }

        got := string(w.written)
        if  want != got {
            t.Errorf("Expected output: %v, but got: %v\n", want, got)
        }

    })
}

// func readLog(path string) (string, error) {
//     file, err := os.Open(path)
//     if err != nil {
//         return "", err
//     }
//     defer file.Close()
//
//     fileRead, err := io.ReadAll(file)
//     if err != nil {
//         return "", err
//     }
//
//     return string(fileRead), nil
// }

type MockResponseWriter struct {
    written []byte
}

func (w *MockResponseWriter) Header() http.Header {
    return nil
}

func (w *MockResponseWriter) Write(p []byte) (n int, err error) {
    w.written = append(w.written, p...)
    return len(p), nil
}

func (w *MockResponseWriter) WriteHeader(statusCode int) {
    // Do nothing
}
