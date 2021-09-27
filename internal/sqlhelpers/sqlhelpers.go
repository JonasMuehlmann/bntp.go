package sqlhelpers

import (
	"github.com/jmoiron/sqlx"
)

// Execute is a wrapper used to run a prepared statement stmt with the specified args.
func Execute(dbConn *sqlx.DB, transaction *sqlx.Tx, stmt string, args ...interface{}) error {
	var statement *sqlx.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statement.Exec(args)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}
