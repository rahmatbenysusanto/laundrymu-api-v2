package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"laundrymu-api/internal/config"
	"time"
)

func MysqlConn() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.LoadConfig().DBUsername,
		config.LoadConfig().DBPassword,
		config.LoadConfig().DBHost,
		config.LoadConfig().DBPort,
		config.LoadConfig().DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("failed to connect to mysql")
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
