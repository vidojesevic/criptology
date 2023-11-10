package server

import (
    "testing"
    "os"
    "cryptology/server"
)

func TestReadFile(t *testing.T) {
    testFile := "testFile.html"
    want := "<h1>Hello, World!</h1>"
    err := os.WriteFile(testFile, []byte(want), 0644)
    if err != nil {
        t.Fatalf("Cannot write file: %v", err)
    }
    defer os.Remove(testFile)

    got := server.ReadFile(testFile)

    if want != got {
        t.Errorf("Expected output: %v, but got: %v\n", want, got)
    }
}
