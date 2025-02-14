package postgres

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// TO-DO: move to cfg
const (
	host     = "postgres"
	port     = "5432"
	user     = "user"
	dbname   = "pingme"
	password = "password"

	driver = "pgx"

	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

func NewPsqlDB() (*sqlx.DB, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := sqlx.Connect(driver, conn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
