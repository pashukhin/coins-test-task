package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

var errStorageRowsAffected = errors.New(" count of affected rows is not 1")

type repository struct {
	db *sqlx.DB
}

func (r repository) ExecForOne(sql string, args ...interface{}) error {
	result, err := r.db.Exec(sql, args...)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return errStorageRowsAffected
	}
	return nil
}