package repository

import (
	"database/sql"
	"errors"
)

func SkipNotFound(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return nil
	}
	return err
}
