package server

import (
	"bytes"
	"criptology/server"
	"log"
	"net/http"
	"os"
    "strings"
	"testing"
)

func TestServeCriptologyStr(t *testing.T) {
    t.Run("When parameter is \"\"", func(t *testing.T) {
        data := ""
        w := &MockResponseWriter{}
        w.written = nil

        var logBuffer bytes.Buffer
        log.SetOutput(&logBuffer)

        server.ServeCriprologyStr(w, data)

        log.SetOutput(os.Stderr)

        want := "Passing empty string!"
        logOut := logBuffer.String()
        if !strings.Contains(logOut, want) {
            t.Errorf("Expected output: %v, but got: %v\n", want, logOut)
        }

    })

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
