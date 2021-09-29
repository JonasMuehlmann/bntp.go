// Package libtlinks implements functionality to work with links and backlinks in a database context.
package liblinks

import (
	"github.com/JonasMuehlmann/bntp.go/internal/sqlhelpers"
	"github.com/jmoiron/sqlx"
)

func GetIdFromType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) (int, error) {
	stmt := `
        SELECT
            Id
        FROM
            DocumentType
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

	err = statement.Get(typeId, type_)

	if err != nil {
		return 0, err
	}

	err = statement.Close()

	if err != nil {
		return 0, err
	}

	return typeId, nil
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

	var documentId int

	err = statement.Get(documentId, documentPath)

	if err != nil {
		return 0, err
	}

	err = statement.Close()

	if err != nil {
		return 0, err
	}

	return documentId, nil
}

// AddLink adds a link from source to destination to the DB.
// Passing a transaction is optional.
func AddLink(dbConn *sqlx.DB, transaction *sqlx.Tx, source string, destination string) error {
	stmt := `
        INSERT INTO
            Link(
                Source,
                Destination
            )
        VALUES(
            ?,
            ?
        );
    `

	sourceId, err := GetIdFromDocument(dbConn, transaction, source)
	if err != nil {
		return err
	}

	destinationId, err := GetIdFromDocument(dbConn, transaction, destination)
	if err != nil {
		return err
	}

	return sqlhelpers.Execute(dbConn, transaction, stmt, sourceId, destinationId)
}

// RemoveLink removes the link from source to destination from the db.
// Passing a transaction is optional.
func RemoveLink(dbConn *sqlx.DB, transaction *sqlx.Tx, source string, destination string) error {
	stmt := `
        DELETE FROM
            Link
        WHERE
            Source = ?
            AND
            Destination = ?;
    `
	sourceId, err := GetIdFromDocument(dbConn, transaction, source)
	if err != nil {
		return err
	}

	destinationId, err := GetIdFromDocument(dbConn, transaction, destination)
	if err != nil {
		return err
	}
	return sqlhelpers.Execute(dbConn, transaction, stmt, sourceId, destinationId)
}

// ListLinks lists all link destionations from the DB.
func ListLinks(dbConn *sqlx.DB, source string) ([]string, error) {
	// TODO: Fix sql statement after db refactor
	stmt := `
        SELECT
            Destination
        FROM
            Link
        WHERE
            Source = ?;
    `

	sourceId, err := GetIdFromDocument(dbConn, nil, source)
	if err != nil {
		return nil, err
	}

	linksBuffer := []string{}

	statementLinks, err := dbConn.Preparex(stmt)
	if err != nil {
		return nil, err
	}

	defer statementLinks.Close()

	err = dbConn.Select(linksBuffer, stmt, sourceId)
	if err != nil {
		return nil, err
	}

	return linksBuffer, nil
}

// ListBacklinks lists all link sources from the DB.
func ListBacklinks(dbConn *sqlx.DB, destination string) ([]string, error) {
	stmt := `
        SELECT
            Source
        FROM
            Link
        WHERE
            Destination = ?;
    `

	destinationId, err := GetIdFromDocument(dbConn, nil, destination)
	if err != nil {
		return nil, err
	}
	backlinksBuffer := []string{}

	statementLinks, err := dbConn.Preparex(stmt)
	if err != nil {
		return nil, err
	}

	defer statementLinks.Close()

	err = dbConn.Select(backlinksBuffer, stmt, destinationId)
	if err != nil {
		return nil, err
	}

	return backlinksBuffer, nil
}
