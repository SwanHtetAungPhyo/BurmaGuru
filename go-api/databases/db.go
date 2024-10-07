package databases

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Ensure you import the SQLite driver
)
var DataBase *sql.DB

func InitDB(DataSourceName string) error {
    var err error
    DataBase, err = sql.Open("sqlite3", DataSourceName)
    if err != nil {
        log.Printf("Error opening database: %v", err)
        return err
    }
    if err = DataBase.Ping(); err != nil {
        log.Printf("Error connecting to database: %v", err)
        return err
    }
    return nil
}