package logger

import (
	"criptology/logger"
	"os"
	"testing"
)

func createTestLogFile(path string) (*os.File, error) {
	return os.Create(path)
}

func TestOpenLogFile(t *testing.T) {
    t.Run("Empty string", func(t *testing.T) {
        logPath := ""
        file, err := logger.OpenLogFile(logPath)
        if file != nil && err != nil {
            t.Errorf("File and error should be nil!")
        }
    })
    t.Run("Valid file path", func(t *testing.T) {
        logPath := "test.log"
		file, err := createTestLogFile(logPath)
		if err != nil {
			t.Errorf("Error creating test log file: %v", err)
		}
		defer file.Close()

        file, err = logger.OpenLogFile(logPath)
		if file == nil || err != nil {
			t.Errorf("Expected non-nil file and nil error, but got file: %v, error: %v", file, err)
		}
    })
}
