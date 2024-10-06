package dbconfig

import "fmt"

var PostgresDriver = "postgres"
var PostgresUser = "postgres"
var PostgresHost = "host.docker.internal"
var PostgresPort = "5432"
var PostgresPassword = "admin@123"
var PostgresDbName = "api-avulso"

var DataSourceName = fmt.Sprintf(
	"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	PostgresHost, PostgresPort, PostgresUser, PostgresPassword, PostgresDbName,
)
