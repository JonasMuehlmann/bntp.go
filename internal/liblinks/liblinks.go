package liblinks

import (
	"database/sql"
	"log"
)

func AddLink(dbConn *sql.DB, transaction *sql.Tx, source string, destination string) {
	stmt := `
        INSERT INTO
            Link(
                Source,
                Destination
            )
        VALUES(
            '?',
            '?'
        );
    `
	var statement *sql.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Prepare(stmt)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		statement, err = dbConn.Prepare(stmt)

		if err != nil {
			log.Fatal(err)
		}
	}
	defer func() {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = statement.Exec(source, destination)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func RemoveLink(dbConn *sql.DB, transaction *sql.Tx, source string, destination string) {
	stmt := `
        DELETE FROM
            Link
        WHERE
            Source = '?'
            AND
            Destination = '?';
    `
	var statement *sql.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Prepare(stmt)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		statement, err = dbConn.Prepare(stmt)

		if err != nil {
			log.Fatal(err)
		}
	}
	defer func() {
		err := statement.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = statement.Exec(source, destination)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func ListLinks(dbConn *sql.DB, source string) []string {
	stmtLinks := `
        SELECT
            Destination
        FROM
            Link
        WHERE
            Source = '?';
    `

	statementLinks, err := dbConn.Prepare(stmtLinks)
	if err != nil {
		log.Fatal(err)
	}

	linkRows, err := statementLinks.Query(source)
	if err != nil {
		log.Fatal(err)
	}

	stmtCountLinks := "SELECT COUNT(*) FROM Link WHERE Source = '?';"

	statementLinksCount, err := dbConn.Prepare(stmtCountLinks)
	if err != nil {
		log.Fatal(err)
	}
	linksCountRow := statementLinksCount.QueryRow(source)

	var rowCountLinks int

	err = linksCountRow.Scan(&rowCountLinks)
	if err != nil {
		log.Fatal(err)
	}
	linksBuffer := make([]string, rowCountLinks)

	i := 0
	for linkRows.Next() {
		err := linkRows.Scan(&linksBuffer[i])
		if err != nil {
			log.Fatal(err)
		}
		i++
	}

	return linksBuffer
}

func ListBacklinks(dbConn *sql.DB, destination string) []string {
	stmtBacklinks := `
        SELECT
            Source
        FROM
            Link
        WHERE
            Destination = '?';
    `

	statementBacklinks, err := dbConn.Prepare(stmtBacklinks)
	if err != nil {
		log.Fatal(err)
	}

	backlinkRows, err := statementBacklinks.Query(destination)
	if err != nil {
		log.Fatal(err)
	}
	stmtCountBacklinks := "SELECT COUNT(*) FROM Link  WHERE Destination = '?';"

	statementBacklinksCount, err := dbConn.Prepare(stmtCountBacklinks)
	if err != nil {
		log.Fatal(err)
	}
	backLinksCountRow := statementBacklinksCount.QueryRow(destination)

	var rowCountLinks int

	err = backLinksCountRow.Scan(&rowCountLinks)
	if err != nil {
		log.Fatal(err)
	}
	backlinksBuffer := make([]string, rowCountLinks)

	i := 0
	for backlinkRows.Next() {
		err := backlinkRows.Scan(&backlinksBuffer[i])
		if err != nil {
			log.Fatal(err)
		}
		i++
	}

	return backlinksBuffer
}
