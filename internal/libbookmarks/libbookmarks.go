package libbookmarks

import (
	"database/sql"
	"log"
	"os"
	"time"

	"encoding/csv"

	_ "github.com/mattn/go-sqlite3"
)

func AddFromCSV(dbConn *sql.DB, csvPath string) {
	file, err := os.Open(csvPath)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	reader.Comma = ';'

	bookmarks, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, bookmark := range bookmarks {
		AddBookmark(dbConn, bookmark[0], bookmark[1], 1, false)
	}
}

func AddBookmark(dbConn *sql.DB, title string, url string, type_ int, isCollection bool) {
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
	statement, err := dbConn.Prepare(stmt)
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec(title, url, time.Now().Format(time.RFC822), type_, isCollection)
	statement.Close()
}
