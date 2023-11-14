package datautil

import (
    "testing"
    "criptology/datautil"
)

func TestGetConfig(t *testing.T) {
    t.Run("Passing empty string", func(t *testing.T) {
        empty := ""
        configValue := datautil.GetConfig(empty)
        if configValue != "" {
            t.Errorf("Expected empty string for invalid configuration key \"%s\", but got \"%s\"", empty, configValue)
        }
    })
    t.Run("Passing random string", func(t *testing.T) {
        invalidConfigKey := "foo"
        configValue := datautil.GetConfig(invalidConfigKey)
        if configValue != "" {
            t.Errorf("Expected empty string for invalid configuration key \"%s\", but got \"%s\"", invalidConfigKey, configValue)
        }
    })
    t.Run("Passing port", func(t *testing.T) {
        portEx := ":9000"
        port := datautil.GetConfig("port")
        if port != portEx {
            t.Errorf("Expected port configuration key to be \"%s\", but got \"%s\"", portEx, port)
        }
    })
    t.Run("Passing app", func(t *testing.T) {
        appName := "Criptology"
        app := datautil.GetConfig("app")
        if app != appName {
            t.Errorf("Expected app configuration key to be \"%s\", but got \"%s\"", appName, app)
        }
    })
}
