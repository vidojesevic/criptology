package datautil

import (
    "os"
    "io"
    "cryptology/logger"
    "encoding/json"
)

type ServerConfig struct {
    Port string `json:"port"`
    AppName string `json:"app"`
}

func GetConfig(name string) string {
    conf, err := os.Open("config/config.json")
    if err != nil {
        logger.WriteErrorLogFile("Cannot open config file")
        return ""
    }

    configByte, err := io.ReadAll(conf)
    if err != nil {
        logger.WriteErrorLogFile("Cannot read config file")
        return ""
    }
    
    var serverConfig []ServerConfig
    er := json.Unmarshal(configByte, &serverConfig)
    if er != nil {
        logger.WriteErrorLogFile("Cannot Unmarshal json")
        return ""
    }

    if len(serverConfig) > 0 {
        switch name {
            case "port":
                return serverConfig[0].Port
            case "app":
                return serverConfig[0].AppName
            default:
                logger.WriteErrorLogFile("Passing invalid config name")
        }
    }

    return ""
}
