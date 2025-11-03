package db

import (
	"os"
)

func buildDsn() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbTimeZone := os.Getenv("DB_TIMEZONE")
	dbSslMode := os.Getenv("DB_SSLMODE")
	dsn := "host=" + dbHost +
		" port=" + dbPort +
		" user=" + dbUser +
		" password=" + dbPassword +
		" dbname=" + dbName +
		" sslmode=" + dbSslMode +
		" TimeZone=" + dbTimeZone

	return dsn
}
