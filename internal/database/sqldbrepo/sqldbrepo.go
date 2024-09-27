package sqldbrepo

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/amledigital/arcxp-circulations/internal/config"
)

type sqlService struct {
	conn *sql.DB
}

func (sqls *sqlService) PrintMsg() {
	fmt.Println("hello world")
}

func NewSQLConn(app *config.AppConfig) (*sqlService, error) {

	db, err := sql.Open("mysql", app.DSN)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(30)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(time.Second * 30)
	db.SetConnMaxLifetime(time.Minute * 30)

	return &sqlService{
		conn: db,
	}, nil

}
