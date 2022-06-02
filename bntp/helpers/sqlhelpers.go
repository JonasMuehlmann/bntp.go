// Copyright © 2021-2022 Jonas Muehlmann
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the "Software"),
// to deal in the Software without restriction, including without limitation
// the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
// TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
// OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

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
	db, err := sqlx.Open("sqlite3", "~/Documents/pkg/bookmarks.db?_foreign_keys=1")
	if err != nil {
		return nil, err
	}

	return db, nil
}
