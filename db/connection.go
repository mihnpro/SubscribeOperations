package db

import (
	"context"
	"fmt"
	"log"
	"test_task/db/migrations"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func Init() {

	var (
		ctx    = context.Background()
		dsn    = buildDsn()
		driver = "postgres"
	)

	if dsn == "" {
		log.Fatal("DSN environment variable is not set")
	}

	log.Println("Connecting to database...")

	db, err := sqlx.ConnectContext(ctx, driver, dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}
	log.Println("Connected to database successfully")

	DB = db

	if err := ApplyMigrations(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database initialization completed successfully!")

	DB.SetMaxOpenConns(5)
	DB.SetMaxIdleConns(2)
	DB.SetConnMaxLifetime(60 * time.Minute)
	DB.SetConnMaxIdleTime(30 * time.Minute)


}

func ApplyMigrations() error {
	log.Println("Applying migrations with separate connection...")

	dsn := buildDsn()
	migrationDB, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to connect for migrations: %v", err)
	}
	defer migrationDB.Close()

	migrator := migrations.MustNewMigrator()
	err = migrator.ApplyMigrations(migrationDB)
	if err != nil {
		return fmt.Errorf("error applying migrations: %v", err)
	}

	return nil
}

func GetDB() *sqlx.DB {
	return DB
}
