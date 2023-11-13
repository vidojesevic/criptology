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
    // configByte, err := io.ReadAll(conf)
    // if err != nil {
    //     logger.WriteErrorLogFile("Cannot read config file")
    //     return ""
    // }
    
    fmt.Println("Config file contents: ", decoder)
    fmt.Printf("Config type is : %T\n", decoder)

    var serverConfig []ServerConfig

    // er := json.Unmarshal(configByte, &serverConfig)
    if er := decoder.Decode(&serverConfig); er != nil {
        var message = fmt.Sprintf("Cannot Unmarshal json: %v", er)
        logger.WriteErrorLogFile(message)
        return ""
    }
    fmt.Printf("Port: %v", serverConfig[0])

    if len(serverConfig) > 0 {
        switch name {
            case "port":
                fmt.Println("Opa")
                return serverConfig[0].Server.Port
            case "app":
                return serverConfig[0].Server.AppName
            default:
                logger.WriteErrorLogFile("Passing invalid config name")
                return "Opa"
        }
    }

    return ""
}
