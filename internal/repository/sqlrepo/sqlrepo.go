package sqlrepo

import (
	"database/sql"

	"github.com/amledigital/arcxp-circulations/internal/config"
)

type SqlRepo struct {
	App  *config.AppConfig
	Conn *sql.DB
}

func NewSQLRepo(a *config.AppConfig, conn *sql.DB) *SqlRepo {
	return &SqlRepo{
		App:  a,
		Conn: conn,
	}
}

func (sqlrepo *SqlRepo) TestPrint() string {
	return ""
}
