package config

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func NewInitializedDatabase(configuration Config) (*sql.DB, error) {
	db, err := NewPostgresSQL(configuration)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func NewPostgresSQL(configuration Config) (*sql.DB, error) {
	username := configuration.Get("DB_USERNAME")
	password := configuration.Get("DB_PASSWORD")
	host := configuration.Get("DB_HOST")
	port := configuration.Get("DB_PORT")
	database := configuration.Get("DB_DATABASE")
	sslMode := configuration.Get("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, database, sslMode)
	db, err := sql.Open(configuration.Get("DB_CONNECTION"), dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	err = databasePooling(configuration, db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func databasePooling(configuration Config, db *sql.DB) error {
	// Limit connection with db pooling
	setMaxIdleConns, err := strconv.Atoi(configuration.Get("DATABASE_POOL_MAX_IDLE"))
	if err != nil {
		return err
	}
	setMaxOpenConns, err := strconv.Atoi(configuration.Get("DATABASE_POOL_MAX_OPEN"))
	if err != nil {
		return err
	}
	setConnMaxIdleTime, err := strconv.Atoi(configuration.Get("DATABASE_MAX_IDLE_TIME_SECOND"))
	if err != nil {
		return err
	}
	setConnMaxLifetime, err := strconv.Atoi(configuration.Get("DATABASE_MAX_LIFE_TIME_SECOND"))
	if err != nil {
		return err
	}

	db.SetMaxIdleConns(setMaxIdleConns)                                    // maximal idle connection
	db.SetMaxOpenConns(setMaxOpenConns)                                    // maximal open connection
	db.SetConnMaxLifetime(time.Duration(setConnMaxIdleTime) * time.Minute) // unused connections will be deleted
	db.SetConnMaxIdleTime(time.Duration(setConnMaxLifetime) * time.Minute) // connection that can be used

	return nil
}
