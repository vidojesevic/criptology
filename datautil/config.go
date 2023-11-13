package datautil

import (
    "os"
    // "io"
    "fmt"
    "cryptology/logger"
    "encoding/json"
)

type ServerConfig struct {
    Server struct {
        Port string `json:"port"`
        AppName string `json:"app"`
    } `json:"server"`
}

func GetConfig(name string) string {
    conf, err := os.Open("config/config.json")
    if err != nil {
        var message = fmt.Sprintf("Cannot open json file: %v", err)
        logger.WriteErrorLogFile(message)
        return ""
    }
    defer conf.Close()

    decoder := json.NewDecoder(conf)
    var serverConfig []ServerConfig

    if er := decoder.Decode(&serverConfig); er != nil {
        var message = fmt.Sprintf("Cannot Unmarshal json: %v", er)
        logger.WriteErrorLogFile(message)
        return ""
    }

    if len(serverConfig) > 0 {
        switch name {
            case "port":
                return serverConfig[0].Server.Port
            case "app":
                return serverConfig[0].Server.AppName
            default:
                logger.WriteErrorLogFile("Passing invalid config name")
        }
    }

    return ""
}
