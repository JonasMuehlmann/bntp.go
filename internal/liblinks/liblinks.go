// Package libtlinks implements functionality to work with links and backlinks in a database context.
package liblinks

import (
	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/jmoiron/sqlx"
)

// AddLink adds a link from source to destination to the DB.
// Passing a transaction is optional.
func AddLink(dbConn *sqlx.DB, transaction *sqlx.Tx, source string, destination string) error {
	stmt := `
        INSERT INTO
            Link(
                SourceId,
                DestinationId
            )
        VALUES(
            ?,
            ?
        );
    `

	sourceId, err := helpers.GetIdFromDocument(dbConn, transaction, source)
	if err != nil {
		return err
	}

	destinationId, err := helpers.GetIdFromDocument(dbConn, transaction, destination)
	if err != nil {
		return err
	}

	return helpers.SqlExecute(dbConn, transaction, stmt, sourceId, destinationId)
}

// RemoveLink removes the link from source to destination from the db.
// Passing a transaction is optional.
func RemoveLink(dbConn *sqlx.DB, transaction *sqlx.Tx, source string, destination string) error {
	stmt := `
        DELETE FROM
            Link
        WHERE
            SourceId = ?
            AND
            DestinationId = ?;
    `
	sourceId, err := helpers.GetIdFromDocument(dbConn, transaction, source)
	if err != nil {
		return err
	}

	destinationId, err := helpers.GetIdFromDocument(dbConn, transaction, destination)
	if err != nil {
		return err
	}
	return helpers.SqlExecute(dbConn, transaction, stmt, sourceId, destinationId)
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
            SourceId = ?;
    `

	sourceId, err := helpers.GetIdFromDocument(dbConn, nil, source)
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
            DestinationId = ?;
    `

	destinationId, err := helpers.GetIdFromDocument(dbConn, nil, destination)
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
