package postgres

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type Db struct {
	MySQL *sqlx.DB
}

func NewDb(cfg *Config) (*Db, error) {
	dataSource := fmt.Sprintf("user=%s password= %s host=%s port=%d dbname=%s sslmode=disable",
		cfg.Dsn.User, cfg.Dsn.Password, cfg.Dsn.Host, cfg.Dsn.Port, cfg.Dsn.DB,
	)

	conn, err := sqlx.Connect("postgres", dataSource)
	if err != nil {
		return nil, errors.New("failed to connect to postgres" + err.Error())
	}

	err = conn.Ping()
	if err != nil {
		return nil, errors.New("failed to ping postgres" + err.Error())
	}
	return &Db{
		MySQL: conn,
	}, nil
}
