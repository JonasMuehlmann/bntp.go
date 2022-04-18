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

func GetIdFromDocumentType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) (int, error) {
	stmt := `
        SELECT
            Id
        FROM
            DocumentType
        WHERE
            DocumentType = ?;
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	}

	if err != nil {
		return -1, err
	}

	typeId := -1

	_ = statement.Get(&typeId, type_)

	err = statement.Close()

	if err != nil {
		return -1, err
	}

	return typeId, nil
}

func GetIdFromBookmarkType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) (int, error) {
	stmt := `
        SELECT
            Id
        FROM
            BookmarkType
        WHERE
            Type = ?;
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	}

	if err != nil {
		return -1, err
	}

	typeId := -1

	_ = statement.Get(&typeId, type_)

	err = statement.Close()

	if err != nil {
		return -1, err
	}

	return typeId, nil
}

func GetIdFromBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, url string) (int, error) {
	stmt := `
        SELECT
            Id
        FROM
            Bookmark
        WHERE
            Url = ?;
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	}

	if err != nil {
		return -1, err
	}

	bookmarkId := -1

	_ = statement.Get(&bookmarkId, url)

	err = statement.Close()

	if err != nil {
		return -1, err
	}

	return bookmarkId, nil
}

func GetIdFromDocument(dbConn *sqlx.DB, transaction *sqlx.Tx, documentPath string) (int, error) {
	stmt := `
        SELECT
            Id
        FROM
            Document
        WHERE
            Path = ?;
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return -1, err
		}
	}

	if err != nil {
		return -1, err
	}

	documentId := -1

	_ = statement.Get(&documentId, documentPath)

	err = statement.Close()

	if err != nil {
		return -1, err
	}

	return documentId, nil
}

func GetIdFromTag(dbConn *sqlx.DB, transaction *sqlx.Tx, tag string) (int, error) {
	stmt := `
        SELECT
            Id
        FROM
            Tag
        WHERE
            Tag = ?;
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return 0, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	var tagId int

	err = statement.Get(&tagId, tag)

	if err != nil {
		return 0, err
	}

	err = statement.Close()

	if err != nil {
		return 0, err
	}

	return tagId, nil
}

func GetIdFromType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) (int, error) {
	stmt := `
        SELECT
            Id
        FROM
            BookmarkType
        WHERE
            Type = ?;
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return 0, err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return 0, err
		}
	}

	if err != nil {
		return 0, err
	}

	var typeId int

	err = statement.Get(&typeId, type_)

	if err != nil {
		return 0, err
	}

	err = statement.Close()

	if err != nil {
		return 0, err
	}

	return typeId, nil
}

// OnError executes handler(err) if err != nil.
func OnError(err error, handler func(args ...interface{})) {
	if err != nil {
		handler(err)
	}
}
