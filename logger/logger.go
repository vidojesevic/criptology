package logger

import (
    "log"
    "os"
    "strings"
    // "path/filepath"
)

func WriteErrorLogFile(message string) {
    logPath := "logger/error.log"

    logFile, err := OpenLogFile(logPath)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()

    logger := log.New(logFile, "[Error] ", log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
    if strings.Contains(message, "%v") {
        logger.Printf(message)
    } else {
        logger.Println(message)
    }
}

func WriteAccessLogFile(message string) {
    logPath := "logger/access.log"

    logFile, err := OpenLogFile(logPath)
    if err != nil {
        log.Fatal(err)
    }
    defer logFile.Close()

    logger := log.New(logFile, "[Access] ", log.LstdFlags | log.Lshortfile | log.Lmicroseconds)
    logger.Println(message)
}

// func getLogFilePath(logFileName string) string {
// 	dir, err := os.Getwd()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return filepath.Join(dir, "logger", logFileName)
// }

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

