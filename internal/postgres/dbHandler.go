package postgres

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type DBHandler interface {
	AddLink(link string) (int, error)
	GetLink(id int) (string, error)
}

type Config struct {
	Host, Port, User, Password, Name, Driver string
}

func (c *Config) Dsn() string {
	return fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", c.Driver, c.User, c.Password, c.Host, c.Port, c.Name)

}

func Get(db *sql.DB, override bool) (DBHandler, error) {
	if override {
		if err := overrideDB(db); err != nil {
			return nil, err
		}
	} else {
		if err := createTables(db); err != nil {
			return nil, err
		}
	}

	return &dbProcessor{sqlx.NewDb(db, "postgres")}, nil
}
