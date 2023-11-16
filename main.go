package main

import (
    "fmt"
    "criptology/server"
    "criptology/logger"
    "criptology/datautil"
)

func main() {

    mes := fmt.Sprintf("%v's server succedully started!\n", datautil.GetConfig("app"))
    fmt.Print(server.Hello(mes))
    port := datautil.GetConfig("port")
    fmt.Printf("Listening port: %s\n", port)
    logger.WriteAccessLogFile("Server succesfully started")

    // serving html
    server.Server("/")
    server.Server("/scripts")

    server.Server("/footer")
}
