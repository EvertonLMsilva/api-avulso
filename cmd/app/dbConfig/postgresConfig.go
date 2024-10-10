package dbConfig

import "fmt"

var PostgresDriver = "postgres"
var PostgresUser = "users"
var PostgresHost = "localhost"
var PostgresPort = "5432"
var PostgresPassword = "admin@123"
var PostgresDbName = "api_avulso"

var DataSourceName = fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	PostgresHost, PostgresPort, PostgresUser, PostgresPassword, PostgresDbName,
)
