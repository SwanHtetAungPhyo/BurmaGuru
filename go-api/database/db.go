package database

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/swanhtetaungphyo/burmaguru/utils"
    "log"
)

var SqlQuaries *utils.SqlQuaries

var DataBase *sql.DB

func ConnectDB() {
	var err error

	DataBase, err = sql.Open("mysql", "root:Swanhtet12@@tcp(localhost:3306)/BurmaGuru")
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	if err := DataBase.Ping(); err != nil {
		log.Fatalf("Database not reachable: %v", err)
	}
}

func InitSqlQuaries(filepathname string) {
	var err error
	SqlQuaries, err = utils.LoadSqlQuaries(filepathname)
	if err != nil {
		log.Printf("Fatal Error: %v", err)
		return
	}
}
