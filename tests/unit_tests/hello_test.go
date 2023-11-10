package server

import (
    "testing"
    "cryptology/server"
)

func TestHello(t *testing.T) {
    got := server.Hello("Hello, World!")
    want := "Hello, World!"

    if got != want {
        t.Errorf("Expected output: %v, but got: %v\n", want, got)
    }

}
