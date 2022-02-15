package helpers

import (
	"github.com/jmoiron/sqlx"
)

// SqlExecute is a wrapper used to run a prepared statement stmt with the specified args.
func SqlExecute(dbConn *sqlx.DB, transaction *sqlx.Tx, stmt string, args ...interface{}) (int64, int64, error) {
	var statement *sqlx.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return 0, 0, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return 0, 0, err
		}
	}

	res, err := statement.Exec(args...)
	if err != nil {
		return 0, 0, err
	}

	err = statement.Close()

	if err != nil {
		return 0, 0, err
	}

	numAffectedRows, err := res.RowsAffected()
	if err != nil {
		return 0, 0, err
	}
	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		return 0, 0, err
	}
	return lastInsertedId, numAffectedRows, nil
}
