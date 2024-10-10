package dbConfig

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDb() (db *sql.DB) {
	db, err := sql.Open(PostgresDriver, DataSourceName)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Println("Db Connected error!", sql.ErrConnDone)
	} else {
		log.Printf("Db Connected!")
	}

	schema := "profile"
	db.Exec(`set search_path=%s`, schema)
	return db
}
