package postgres

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type dbProcessor struct {
	db *sqlx.DB
}

var (
	errBeginTx  = errors.New("error while starting transaction")
	errCommitTx = errors.New("error while committing transaction")
)

func (d *dbProcessor) AddLink(link string) (int, error) {
	wrapErr := errors.New("error while inserting link to the database")
	var id int
	if err := d.db.Get(&id, getLinkId, link); err == nil {
		return id, nil
	}

	tx, err := d.db.Begin()
	if err != nil {
		return 0, errors.Join(wrapErr, errBeginTx, err)
	}
	defer tx.Rollback()

	if err = tx.QueryRow(addLink, link).Scan(&id); err != nil {
		return 0, errors.Join(wrapErr, err)
	}

	if err = tx.Commit(); err != nil {
		return 0, errors.Join(wrapErr, errCommitTx, err)
	}

	return id, nil
}

func (d *dbProcessor) GetLink(id int) (string, error) {
	var link string
	if err := d.db.Get(&link, getLink, id); err != nil {
		return "", errors.Join(errors.New("error while getting user from the database"), err)
	}

	return link, nil
}
