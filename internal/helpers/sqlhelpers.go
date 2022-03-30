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

// SqlExecuteNamed is a wrapper used to run a prepared statement stmt with the specified args.
func SqlExecuteNamed(dbConn *sqlx.DB, transaction *sqlx.Tx, stmt string, struct_ interface{}) (int64, int64, error) {
	var statement *sqlx.NamedStmt

	var err error

	if transaction != nil {
		statement, err = transaction.PrepareNamed(stmt)

		if err != nil {
			return 0, 0, err
		}
	} else {
		statement, err = dbConn.PrepareNamed(stmt)

		if err != nil {
			return 0, 0, err
		}
	}

	res, err := statement.Exec(struct_)
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

func GetDefaultDB() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "~/Documents/bntp/bookmarks.db?_foreign_keys=1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
