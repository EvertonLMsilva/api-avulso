package dbConfig

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/EvertonLMsilva/api-avulso/internal/infra/environments"
	_ "github.com/lib/pq"
)

func ConnectDb() (db *sql.DB) {
	var dataSourceName = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		environments.Env.DatabaseHost,
		environments.Env.DatabasePort,
		environments.Env.DatabaseUsername,
		environments.Env.DatabasePassword,
		environments.Env.DatabaseDBName,
	)

	db, err := sql.Open(environments.Env.DatabaseDrive, dataSourceName)

	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Println("Db Connected error!", sql.ErrConnDone)
	} else {
		log.Printf("Db Connected!")
	}

	schema := environments.Env.DatabaseSchema
	db.Exec(`set search_path=%s`, schema)
	return db
}
