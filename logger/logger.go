package logger

import (
    "log"
    "os"
    "fmt"
    "strings"
    "runtime"
    "path"
)

func WriteErrorLogFile(message string) {
    logPath := "logger/error.log"
    if message == "" {
        log.Println("Error: there is no message for log file")
        return
    }

    logFile, err := OpenLogFile(logPath)
    if err != nil {
        log.Printf("Error: %v", err)
        return
    }
    defer func() {
        if logFile != nil {
            logFile.Close()
        }
    }()

    _, file, line, ok := runtime.Caller(1)
    if !ok {
        file = "unknown"
        line = 0
    }
    shortFile := path.Base(file)

    logEntry := fmt.Sprintf("[%s:%d] %s", shortFile, line, message)

    logger := log.New(logFile, "[Error] ", log.LstdFlags | log.Lmicroseconds)
    if strings.Contains(message, "%v") {
        logger.Printf(logEntry)
    } else {
        logger.Println(logEntry)
    }
}

func WriteAccessLogFile(message string) {
    logPath := "logger/access.log"

    logFile, err := OpenLogFile(logPath)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()

    _, file, line, ok := runtime.Caller(1)
    if !ok {
        file = "unknown"
        line = 0
    }
    shortFile := path.Base(file)

    logEntry := fmt.Sprintf("[%s:%d] %s", shortFile, line, message)

    logger := log.New(logFile, "[Access] ", log.LstdFlags | log.Lmicroseconds)
    logger.Println(logEntry)
}

func OpenLogFile(path string) (*os.File, error) {
    if path == "" {
        return nil, nil
    }
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

