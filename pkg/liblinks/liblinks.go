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

// Package libtlinks implements functionality to work with links and backlinks in a database context.
package liblinks

import (
	"errors"

	"github.com/JonasMuehlmann/bntp.go/pkg/helpers"
	"github.com/jmoiron/sqlx"
)

// TODO: Allow passing string and id for document

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

	if sourceId == -1 {
		return errors.New("Could not retrieve DestinationId")
	}

	destinationId, err := helpers.GetIdFromDocument(dbConn, transaction, destination)
	if err != nil {
		return err
	}

	if destinationId == -1 {
		return errors.New("Could not retrieve DestinationId")
	}

	_, _, err = helpers.SqlExecute(dbConn, transaction, stmt, sourceId, destinationId)

	return err
}

// TODO: Allow passing string and id for document

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

	if sourceId == -1 {
		return errors.New("Could not retrieve DestinationId")
	}

	destinationId, err := helpers.GetIdFromDocument(dbConn, transaction, destination)
	if err != nil {
		return err
	}

	if destinationId == -1 {
		return errors.New("Could not retrieve DestinationId")
	}

	_, _, err = helpers.SqlExecute(dbConn, transaction, stmt, sourceId, destinationId)

	return err
}

// TODO: Allow passing string and id for document

// ListLinks lists all link destionations from the DB.
func ListLinks(dbConn *sqlx.DB, source string) ([]string, error) {
	stmt := `
        SELECT
            Path
        FROM
            Document
        LEFT  JOIN Link ON
            Link.DestinationId = Document.Id
        WHERE
            SourceId = ?;
    `

	sourceId, err := helpers.GetIdFromDocument(dbConn, nil, source)
	if err != nil {
		return nil, err
	}

	if sourceId == -1 {
		return nil, errors.New("Could not retrieve DestinationId")
	}

	linksBuffer := []string{}

	statementLinks, err := dbConn.Preparex(stmt)
	if err != nil {
		return nil, err
	}

	defer statementLinks.Close()

	err = dbConn.Select(&linksBuffer, stmt, sourceId)
	if err != nil {
		return nil, err
	}

	return linksBuffer, nil
}

// TODO: Allow passing string and id for document

// ListBacklinks lists all link sources from the DB.
func ListBacklinks(dbConn *sqlx.DB, destination string) ([]string, error) {
	stmt := `
        SELECT
            Path
        FROM
            Document
        LEFT  JOIN Link ON
            Link.SourceId = Document.Id
        WHERE
            DestinationId = ?;
    `

	destinationId, err := helpers.GetIdFromDocument(dbConn, nil, destination)
	if err != nil {
		return nil, err
	}

	if destinationId == -1 {
		return nil, errors.New("Could not retrieve DestinationId")
	}

	backlinksBuffer := []string{}

	statementLinks, err := dbConn.Preparex(stmt)
	if err != nil {
		return nil, err
	}

	defer statementLinks.Close()

	err = dbConn.Select(&backlinksBuffer, stmt, destinationId)
	if err != nil {
		return nil, err
	}

	return backlinksBuffer, nil
}
