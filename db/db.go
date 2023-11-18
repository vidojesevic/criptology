package db

import (
    "fmt"
    "database/sql"
    "criptology/logger"
    _ "github.com/go-sql-driver/mysql"
)

func EstablishDBConnection() {
    conf := GetDB()
    // fmt.Printf("db: %v, host: %v, user: %v, pass: %v, dbName: %v\n", conf.db, conf.host, conf.user, conf.pass, conf.name)
    db, err := sql.Open(conf.db, conf.user + ":<" + conf.pass + ">@tcp(" + conf.host + ")/" + conf.name)
    // db, err := sql.Open(conf.db, fmt.Sprintf("%s%s%s", conf.name, conf.user, conf.pass))
    if err != nil {
        message := fmt.Sprintf("Error with database connection: %v", err)
        logger.WriteErrorLogFile(message)
    }
    // defer db.Close()
    logger.WriteAccessLogFile("Successfully connected to database")

    createTable, er := db.Query("CREATE TABLE IF NOT EXIST test (id INT(11) NOT NULL PRIMARY KEY, name VARCHAR(32) NOT NULL);")
    if er != nil {
        message := fmt.Sprintf("Error with creation of table: %v", er)
        logger.WriteErrorLogFile(message)
    }
    // defer createTable.Close()
    fmt.Printf("Table was successfully created: %T", createTable)
}
