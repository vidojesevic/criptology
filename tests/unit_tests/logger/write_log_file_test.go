package logger

import (
	"criptology/logger"
	"log"
	"os"
	"strings"
	"testing"
)

func createLogFile(path string) (*os.File, error) {
    return os.Create(path)
}

func readLogFile(path string) (string, error) {
    content, err := os.ReadFile(path)
    if err != nil {
        return "", err
    }
    return string(content), nil
}

func TestWriteLogFile(t *testing.T) {
    t.Run("Empty parameter", func(t *testing.T) {
        path := "test.log"
        message := ""
        defer os.Remove(path)

        logFile, err := createLogFile(path)
        if err != nil {
            t.Errorf("Error creating a file: %v", err)
            return
        }
        defer logFile.Close()

        log.SetOutput(logFile)
        defer log.SetOutput(os.Stderr)

        logger.WriteErrorLogFile(message)

        content, er := readLogFile(path)
        if er != nil {
            t.Errorf("Error reading log file content: %v", er)
            return
        }
        expected := "Error: there is no message for log file"

        if !strings.Contains(content,expected) {
            t.Errorf("Expected \"%v\", got \"%v\"", expected, content)
        }
    })
}
