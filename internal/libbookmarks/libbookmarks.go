package libbookmarks

import (
	"database/sql"
	"log"
	"os"
	"strings"
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

	for i := range rawBuffer {
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
	stmtRows := `
        SELECT
            *
        FROM
            Type;
    `
	stmtCount := `
        SELECT
            Count(*)
        FROM
            Type;
    `
	countRow := dbConn.QueryRow(stmtCount)
	var rowCount int
	countRow.Scan(&rowCount)

	rows, err := dbConn.Query(stmtRows)
	if err != nil {
		log.Fatal(err)
	}

	types := make([]string, 0, rowCount)

	i := 0
	for rows.Next() {
		rows.Scan(&types[i])
		i++
	}

	return types
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

func MarkAsRead(dbConn *sql.DB, transaction *sql.Tx, id int) {
	EditBookmark(dbConn, transaction, id, "IsRead", true)
}
func EditTitle(dbConn *sql.DB, transaction *sql.Tx, id int, newTitle string) {
	EditBookmark(dbConn, transaction, id, "Title", newTitle)
}
func EditUrl(dbConn *sql.DB, transaction *sql.Tx, id int, newUrl string) {
	EditBookmark(dbConn, transaction, id, "Url", newUrl)
}
func EditType(dbConn *sql.DB, transaction *sql.Tx, id int, newType string) {
	// TODO: Get id of type newType
	var typeId string
	EditBookmark(dbConn, transaction, id, "TypeId", typeId)
}
func EditIsCollection(dbConn *sql.DB, transaction *sql.Tx, id int, isCollection bool) {
	EditBookmark(dbConn, transaction, id, "IsCollection", isCollection)
}
func AddTag(dbConn *sql.DB, transaction *sql.Tx, id int, newTag string) {

}
func RemoveTag(dbConn *sql.DB, transaction *sql.Tx, id int, tag_ string) {

}
func ListBookmarks(dbConn *sql.DB, filters map[string]interface{}) []Bookmark {
	joinFragments := make([]string, 0, 10)
	whereFragments := make([]string, 0, 10)

	for key, value := range filters {
		switch key {
		case "IsRead":
			val, ok := value.(bool)

			if !ok {
				log.Println("filters[\"IsRead\"] should be type bool.")
			}

			var valConverted string
			if val {
				valConverted = "1"
			} else {
				valConverted = "0"
			}

			whereFragments = append(whereFragments, "WHERE IsRead = "+valConverted)
		case "IsCollection":
			val, ok := value.(bool)

			if !ok {
				log.Println("filters[\"IsCollection\"] should be type bool.")
			}

			var valConverted string
			if val {
				valConverted = "1"
			} else {
				valConverted = "0"
			}

			whereFragments = append(whereFragments, "WHERE IsCollection = "+valConverted)
		case "Title":
			value, ok := value.(string)

			if !ok {
				log.Println("filters[\"Title\"] should be type string.")
			}
			whereFragments = append(whereFragments, "WHERE Title = "+value)
		case "Url":
			value, ok := value.(string)

			if !ok {
				log.Println("filters[\"Url\"] should be type string.")
			}
			whereFragments = append(whereFragments, "WHERE Url = "+value)
		case "MaxAge":
			value, ok := value.(int)

			if !ok {
				log.Println("filters[\"MayAge\"] should be type string.")
			}

			// TODO: Compare age in SQL
		case "TimeAddedExact":
			_, ok := filters["TimeAddedExact"]
			if ok {
				log.Fatal("Filter \"TimeAddedExact\" can't be specified with \"TimeAddedRangeStart\"")
			}
			// TODO: Check if date is valid and convert it to db format

			whereFragments = append(whereFragments, "WHERE TimeAdded = "+value)

		case "TimeAddedRangeStart":
			_, ok := filters["TimeAddedExact"]
			if ok {
				log.Fatal("Filter \"TimeAddedExact\" can't be specified with \"TimeAddedRangeStart\"")
			}

			timeAddedRangeEnd, ok := filters["TimeAddedRangeEnd"]
			if !ok {
				log.Fatal("Filter \"TimeAddedStart\" specified without \"TimeAddedRangeEnd\"")
			}
		case "Types":
			types, okArray := value.([]string)
			if !okArray {
				type_, okString := value.(string)
				if !okString {
					log.Fatal("Filter filters[\"Types\"] should be []string or string.")
				} else {
					whereFragments = append(whereFragments, "WHERE Type = "+type_+")")
				}
			}
			whereFragments = append(whereFragments, "WHERE Type IN("+strings.Join(types, ", ")+")")

		case "Tags":
			joinFragments = append(joinFragments, "INNER JOIN Context ON BookmarkId = Bookmark.Id")

			tags, okArray := value.([]string)
			if !okArray {
				log.Fatal("filters[\"Tags\"] should be []string or string.")
				tag, okString := value.(string)
				if !okString {
					log.Fatal("Filter filters[\"Tags\"] should be []string or string.")
				} else {
					whereFragments = append(whereFragments, "WHERE Tag = "+tag+")")
				}
			}
			whereFragments = append(whereFragments, "WHERE Tag IN ("+strings.Join(tags, ", ")+")")
		default:
			log.Fatal("Encountered unrecognized filter.")
		}
	}

}

type Bookmark struct {
	Id           int
	IsRead       bool
	IsCollection bool
	Title        string
	Url          string
	TimeAdded    string
	Type         string
	Tags         []string
}
