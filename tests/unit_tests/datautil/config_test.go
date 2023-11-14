package datautil

import (
    "testing"
    "cryptology/datautil"
)

func TestGetConfig(t *testing.T) {
    portEx := ":9000"
    port := datautil.GetConfig("port")
    if port != portEx {
        t.Errorf("Expected port configuration key to be \"%s\", but got \"%s\"", portEx, port)
    }

    appName := "Cryptology"
    app := datautil.GetConfig("app")
    if app != appName {
        t.Errorf("Expected app configuration key to be \"%s\", but got \"%s\"", appName, app)
    }

    invalidConfigKey := "foo"
    configValue := datautil.GetConfig(invalidConfigKey)
    if configValue != "" {
        t.Errorf("Expected empty string for invalid configuration key \"%s\", but got \"%s\"", invalidConfigKey, configValue)
    }
}
