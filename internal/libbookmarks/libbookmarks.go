package libbookmarks

import (
	"database/sql"
	"log"
	"os"
	"time"

	"encoding/csv"

	_ "github.com/mattn/go-sqlite3"
)

func ImportMinimalCSV(dbConn *sql.DB, csvPath string) {
	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	bookmarks, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	header := bookmarks[0]

	if len(header) != 2 {
		log.Fatal("CSV Header does not have correct number of fields. It should have 2.")
	}

	if !(header[0] == "Title" || header[1] == "Title") || !(header[0] == "Url" || header[1] == "Url") || header[0] == header[1] {
		log.Fatal("CSV Header does not have necessary fields 'Title' and 'Url.'")
	}

	var titleColumn, linkColumn int

	if header[0] == "Title" {
		titleColumn = 0
		linkColumn = 1
	} else {
		titleColumn = 1
		linkColumn = 0
	}

	transaction, err := dbConn.Begin()
	if err != nil {
		log.Fatal(err)
	}
	for _, bookmark := range bookmarks {
		AddBookmark(dbConn, transaction, bookmark[titleColumn], bookmark[linkColumn], 1, false)
	}
	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func ExportFullCSV(dbConn *sql.DB, csvPath string) {
	// ###############
	// # SELECT ROWS #
	// ###############
	stmtRows := `
        SELECT
            Id,
            IsRead,
            Title,
            Url,
            TimeAdded,
            TypeId,
            IsCollection
        FROM
            Bookmark;
    `

	statementRows, err := dbConn.Prepare(stmtRows)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := statementRows.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	rows, err := statementRows.Query()
	if err != nil {
		log.Fatal(err)
	}
	// #################
	// # SELECT  COUNT #
	// #################
	stmtCount := "SELECT COUNT(*) FROM Bookmark;"

	countRow := dbConn.QueryRow(stmtCount)

	var rowCount int

	err = countRow.Scan(&rowCount)
	if err != nil {
		log.Fatal(err)
	}
	// ##############
	// # WRITE FILE #
	// ##############
	// 0664 UNIX Permission code
	file, err := os.OpenFile(csvPath, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	writer := csv.NewWriter(file)
	writer.Comma = ';'

	csvHeader, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Write(csvHeader)
	if err != nil {
		log.Fatal(err)
	}

	// row -> destBuffer -> rawBuffer -> rowsBuffer -> CSV
	rowsBuffer := make([][]string, rowCount)
	destBuffer := make([]interface{}, len(csvHeader))
	rawBuffer := make([][]byte, len(csvHeader))

	for i := range rowsBuffer {
		rowsBuffer[i] = make([]string, len(csvHeader))
	}

	for i, _ := range rawBuffer {
		destBuffer[i] = &rawBuffer[i]
	}

	i := 0
	for rows.Next() {
		err := rows.Scan(destBuffer...)
		if err != nil {
			log.Fatal(err)
		}

		for j, raw := range rawBuffer {
			if raw == nil {
				rowsBuffer[i][j] = "\n"
			} else {
				rowsBuffer[i][j] = string(raw)
			}
		}
		i++
	}

	err = writer.WriteAll(rowsBuffer)
	if err != nil {
		log.Fatal(err)
	}
}

func AddType(dbConn *sql.DB, transaction *sql.Tx, type_ string) {
	stmt := `
        INSERT INTO
            Type(
                Type,
            )
        VALUES(
            ?,
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

	_, err = statement.Exec(type_)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
}
func RemoveType(dbConn *sql.DB, transaction *sql.Tx, type_ string) {
	stmt := `
        DELETE FROM
            Type
        WHERE
            Type = ?;
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

	_, err = statement.Exec(type_)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func ListTypes(dbConn *sql.DB) []string {

}

func AddBookmark(dbConn *sql.DB, transaction *sql.Tx, title string, url string, type_ int, isCollection bool) {
	stmt := `
        INSERT INTO
            Bookmark(
                Title,
                Url,
                TimeAdded,
                TypeId,
                IsCollection
            )
        VALUES(
            ?,
            ?,
            ?,
            ?,
            ?
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

	_, err = statement.Exec(title, url, time.Now().Format(time.RFC822), type_, isCollection)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func EditBookmark(dbConn *sql.DB, transaction *sql.Tx, id int, column string, newVal interface{}) {
}
