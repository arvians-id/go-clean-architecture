package config

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

func NewPostgresSQLGorm(configuration Config) (*gorm.DB, error) {
	username := configuration.Get("DB_USERNAME")
	password := configuration.Get("DB_PASSWORD")
	host := configuration.Get("DB_HOST")
	port := configuration.Get("DB_PORT")
	database := configuration.Get("DB_DATABASE")
	sslMode := configuration.Get("DB_SSL_MODE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, database, sslMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	dbPool, err := db.DB()
	if err != nil {
		return nil, err
	}
	_, err = databasePooling(configuration, dbPool)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func databasePooling(configuration Config, db *sql.DB) (*sql.DB, error) {
	setMaxIdleConns, err := strconv.Atoi(configuration.Get("DB_POOL_MIN"))
	if err != nil {
		return nil, err
	}
	setMaxOpenConns, err := strconv.Atoi(configuration.Get("DB_POOL_MAX"))
	if err != nil {
		return nil, err
	}
	setConnMaxIdleTime, err := strconv.Atoi(configuration.Get("DB_MAX_IDLE_TIME"))
	if err != nil {
		return nil, err
	}
	setConnMaxLifetime, err := strconv.Atoi(configuration.Get("DB_MAX_LIFE_TIME"))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(setMaxIdleConns)                                    // minimal connection
	db.SetMaxOpenConns(setMaxOpenConns)                                    // maximal connection
	db.SetConnMaxLifetime(time.Duration(setConnMaxIdleTime) * time.Second) // unused connections will be deleted
	db.SetConnMaxIdleTime(time.Duration(setConnMaxLifetime) * time.Second)

	return db, nil
}

func NewPostgresSQLGormTest(configuration Config) (*gorm.DB, error) {
	username := configuration.Get("DB_USERNAME_TEST")
	password := configuration.Get("DB_PASSWORD_TEST")
	host := configuration.Get("DB_HOST_TEST")
	port := configuration.Get("DB_PORT_TEST")
	database := configuration.Get("DB_DATABASE_TEST")
	sslMode := configuration.Get("DB_SSL_MODE_TEST")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, username, password, database, sslMode)
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	dbPool, err := db.DB()
	if err != nil {
		return nil, err
	}

	_, err = databasePoolingTest(configuration, dbPool)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func databasePoolingTest(configuration Config, db *sql.DB) (*sql.DB, error) {
	setMaxIdleConns, err := strconv.Atoi(configuration.Get("DB_POOL_MIN_TEST"))
	if err != nil {
		return nil, err
	}
	setMaxOpenConns, err := strconv.Atoi(configuration.Get("DB_POOL_MAX_TEST"))
	if err != nil {
		return nil, err
	}
	setConnMaxIdleTime, err := strconv.Atoi(configuration.Get("DB_MAX_IDLE_TIME_TEST"))
	if err != nil {
		return nil, err
	}
	setConnMaxLifetime, err := strconv.Atoi(configuration.Get("DB_MAX_LIFE_TIME_TEST"))
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(setMaxIdleConns)                                    // minimal connection
	db.SetMaxOpenConns(setMaxOpenConns)                                    // maximal connection
	db.SetConnMaxLifetime(time.Duration(setConnMaxIdleTime) * time.Second) // unused connections will be deleted
	db.SetConnMaxIdleTime(time.Duration(setConnMaxLifetime) * time.Second)

	return db, nil
}

func TearDownDBTest(db *gorm.DB) error {
	return nil
}
