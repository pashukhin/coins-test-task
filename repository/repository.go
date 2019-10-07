// Package repository contains entity repositories.
package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

// ErrStorageRowsAffected is error which throws where sql query affects not exactly one row.
var ErrStorageRowsAffected = errors.New(" count of affected rows is not 1")

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
		return ErrStorageRowsAffected
	}
	return nil
}
