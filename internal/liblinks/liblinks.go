// Package libtlinks implements functionality to work with links and backlinks in a database context.
package liblinks

import (
	"github.com/JonasMuehlmann/bntp.go/internal/sqlhelpers"
	"github.com/jmoiron/sqlx"
)

// AddLink adds a link from source to destination to the DB.
// Passing a transaction is optional.
func AddLink(dbConn *sqlx.DB, transaction *sqlx.Tx, source string, destination string) error {
	// TODO: Fix sql statement after db refactor
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

	return sqlhelpers.Execute(dbConn, transaction, stmt, source, destination)
}

// RemoveLink removes the link from source to destination from the db.
// Passing a transaction is optional.
func RemoveLink(dbConn *sqlx.DB, transaction *sqlx.Tx, source string, destination string) error {
	// TODO: Fix sql statement after db refactor
	stmt := `
        DELETE FROM
            Link
        WHERE
            Source = ?
            AND
            Destination = ?;
    `

	return sqlhelpers.Execute(dbConn, transaction, stmt, source, destination)
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

	linksBuffer := []string{}

	statementLinks, err := dbConn.Preparex(stmt)
	if err != nil {
		return nil, err
	}

	defer statementLinks.Close()

	err = dbConn.Select(linksBuffer, stmt)
	if err != nil {
		return nil, err
	}

	return linksBuffer, nil
}

// ListBacklinks lists all link sources from the DB.
func ListBacklinks(dbConn *sqlx.DB, destination string) ([]string, error) {
	// TODO: Fix sql statement after db refactor
	stmt := `
        SELECT
            Source
        FROM
            Link
        WHERE
            Destination = ?;
    `
	backlinksBuffer := []string{}

	statementLinks, err := dbConn.Preparex(stmt)
	if err != nil {
		return nil, err
	}

	defer statementLinks.Close()

	err = dbConn.Select(backlinksBuffer, stmt)
	if err != nil {
		return nil, err
	}

	return backlinksBuffer, nil
}
