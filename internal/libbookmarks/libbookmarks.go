package libbookmarks

import (
	"database/sql"
	"log"
	"os"
	"strconv"
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

func ExportCSV(dbConn *sql.DB, csvPath string, filters map[string]interface{}) {
	// ####################################
	// # Filter validation and processing #
	// ####################################
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
		case "MaxAge":
			val, ok := value.(int)

			if !ok {
				log.Fatal("filters[\"MayAge\"] should be type string.")
			}
			whereFragments = append(whereFragments, "WHERE timeAdded BETWEEN DATE('now') AND datetime(DATE('now'),'-'"+strconv.Itoa(val)+" days')")
		case "Types":
			types, okArray := value.([]string)
			if !okArray {
				type_, okString := value.(string)
				if !okString {
					log.Fatal("Filter filters[\"Types\"] should be []string or string.")
				} else {
					whereFragments = append(whereFragments, "WHERE Type = '"+type_+"'")
				}
			} else {
				whereFragments = append(whereFragments, "WHERE Type IN('"+strings.Join(types, "', '")+"')")
			}
		case "Tags":
			joinFragments = append(joinFragments, "INNER JOIN Context ON Context.BookmarkId = Bookmark.Id INNER JOIN Tag ON Tag.Id = Context.TagId")

			tags, okArray := value.([]string)
			if !okArray {
				tag, okString := value.(string)
				if !okString {
					log.Fatal("Filter filters[\"Tags\"] should be []string or string.")
				} else {
					whereFragments = append(whereFragments, "WHERE Tag = '"+tag+"'")
				}
			} else {
				whereFragments = append(whereFragments, "WHERE Tag IN ('"+strings.Join(tags, "', '")+"')")
			}
		default:
			log.Fatal("Encountered unrecognized filter.")
		}
	}
	// ###############
	// # SELECT ROWS #
	// ###############
	stmtBookmarks := `
        SELECT
            Bookmark.Id,
            Bookmark.IsRead,
            Bookmark.Title,
            Bookmark.Url,
            Bookmark.TimeAdded,
            Type.Type,
            Bookmark.IsCollection
        FROM
            Bookmark
        INNER JOIN Type ON
            Type.Id = Bookmark.TypeId
    `
	stmtTags := `
        SELECT
            Tag.Tag
        FROM
            Tag
        INNER JOIN Context ON
            Context.TagId = Tag.Id
        WHERE Context.BookmarkId = ?;`

	joinFragment := strings.Join(joinFragments, " AND ")
	whereFragment := strings.Join(whereFragments, " AND ")

	stmtBookmarks += joinFragment + " " + whereFragment + ";"

	statementTags, err := dbConn.Prepare(stmtTags)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = statementTags.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()
	bookmarkRows, err := dbConn.Query(stmtBookmarks)
	if err != nil {
		log.Fatal(err)
	}
	// #####################################
	// # Prepare file write or STDOUT print #
	// #####################################
	var writer *csv.Writer
	if csvPath != "" { // 0664 UNIX Permission code
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
		writer = csv.NewWriter(file)
	} else { // 0664 UNIX Permission code
		writer = csv.NewWriter(os.Stdout)
	}

	writer.Comma = ';'

	csvHeader, err := bookmarkRows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	csvHeader = append(csvHeader, "Tags")

	err = writer.Write(csvHeader)
	if err != nil {
		log.Fatal(err)
	}

	// ############
	// # Get Data #
	// ############
	stmtCountBookmarks := "SELECT COUNT(*) FROM Bookmark INNER JOIN Type ON Bookmark.TypeId = Type.Id" + joinFragment + " " + whereFragment + ";"
	stmtCountTags := "SELECT COUNT(*) FROM  Context WHERE BookmarkId = ?;"

	statementTagsCount, err := dbConn.Prepare(stmtCountTags)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = statementTagsCount.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()
	countRowBookmarks := dbConn.QueryRow(stmtCountBookmarks)

	var rowCountBookmarks int

	err = countRowBookmarks.Scan(&rowCountBookmarks)
	if err != nil {
		log.Fatal(err)
	}
	// row -> destBuffer -> rawBuffer -> rowsBuffer -> CSV
	rowsBuffer := make([][]string, rowCountBookmarks)
	destBuffer := make([]interface{}, len(csvHeader)-1)
	rawBuffer := make([][]byte, len(csvHeader)-1)

	for i := range rowsBuffer {
		rowsBuffer[i] = make([]string, len(csvHeader))
	}

	for i := range rawBuffer {
		destBuffer[i] = &rawBuffer[i]
	}

	// TODO: Try to optimize this into a single joined query on the Context and tag table
	// and then match the result to the bookmarks
	// https://turriate.com/articles/making-sqlite-faster-in-go
	i := 0
	for bookmarkRows.Next() {
		err := bookmarkRows.Scan(destBuffer...)
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
		// ############
		// # Add tags #
		// ############
		var rowCountTags int

		bookmarkIdAsInt, err := strconv.Atoi(rowsBuffer[i][0])
		if err != nil {
			log.Fatal(err)
		}
		countRowTags := statementTagsCount.QueryRow(bookmarkIdAsInt)

		err = countRowTags.Scan(&rowCountTags)
		if err != nil {
			log.Fatal(err)
		}
		tagsBuffer := make([]string, rowCountTags)

		tagRows, err := statementTags.Query(bookmarkIdAsInt)
		if err != nil {
			log.Fatal(err)
		}

		j := 0
		for tagRows.Next() {
			err := tagRows.Scan(&tagsBuffer[j])
			if err != nil {
				log.Fatal(err)
			}
			j++
		}
		rowsBuffer[i][len(csvHeader)-1] = strings.Join(tagsBuffer, ",")

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
                Type
            )
        VALUES(
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

	_, err = statement.Exec(title, url, time.Now().Format("2006-01-02"), type_, isCollection)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
}

func EditBookmark(dbConn *sql.DB, transaction *sql.Tx, id int, column string, newVal interface{}) {
	stmt := `
        UPDATE
            Bookmark
        SET
            ? = ?
        WHERE Id =
    `

	_, ok := newVal.(string)

	if ok {
		stmt += "'?';"
	} else {
		stmt += "?;"
	}

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

	_, err = statement.Exec(column, newVal, id)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
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
	// TODO: Refactor getting Type.Id from Type.Type
	var typeId int

	stmt := `
        SELECT
            Id
        FROM
            Type
        WHERE
            Type = '?';
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

	typeRow := statement.QueryRow(newType)
	if err != nil {
		log.Fatal(err)
	}

	typeRow.Scan(&typeId)

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}

	EditBookmark(dbConn, transaction, id, "TypeId", typeId)
}
func EditIsCollection(dbConn *sql.DB, transaction *sql.Tx, id int, isCollection bool) {
	EditBookmark(dbConn, transaction, id, "IsCollection", isCollection)
}
func AddTag(dbConn *sql.DB, transaction *sql.Tx, bookmarkId int, newTag string) {
	// TODO: Refactor getting Tag.Id from Tag.Tag
	var tagId int

	stmtTag := `
        SELECT
            Id
        FROM
            Tag
        WHERE
            Tag = '?';
    `
	var statementTag *sql.Stmt
	var err error

	if transaction != nil {
		statementTag, err = transaction.Prepare(stmtTag)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		statementTag, err = dbConn.Prepare(stmtTag)

		if err != nil {
			log.Fatal(err)
		}
	}

	tagRow := statementTag.QueryRow(newTag)
	if err != nil {
		log.Fatal(err)
	}

	tagRow.Scan(&tagId)

	err = statementTag.Close()

	if err != nil {
		log.Fatal(err)
	}

	stmt := `
        INSERT INTO
            Context(BookmarkId, TagId)
        VALUES(
            ?,
            ?
    );
    `
	var statementContext *sql.Stmt

	if transaction != nil {
		statementContext, err = transaction.Prepare(stmt)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		statementContext, err = dbConn.Prepare(stmt)

		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = statementContext.Exec(bookmarkId, tagId)
	if err != nil {
		log.Fatal(err)
	}

	err = statementContext.Close()

	if err != nil {
		log.Fatal(err)
	}
}
func RemoveTag(dbConn *sql.DB, transaction *sql.Tx, bookmarkId int, tag_ string) {
	// TODO: Refactor getting Tag.Id from Tag.Tag
	var tagId int

	stmtTag := `
        SELECT
            Id
        FROM
            Tag
        WHERE
            Tag = '?';
    `
	var statementTag *sql.Stmt
	var err error

	if transaction != nil {
		statementTag, err = transaction.Prepare(stmtTag)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		statementTag, err = dbConn.Prepare(stmtTag)

		if err != nil {
			log.Fatal(err)
		}
	}

	tagRow := statementTag.QueryRow(tag_)
	if err != nil {
		log.Fatal(err)
	}

	tagRow.Scan(&tagId)

	err = statementTag.Close()

	if err != nil {
		log.Fatal(err)
	}

	stmt := `
        DELETE FROM
            Context
        WHERE
            BookmarkId = ?
            AND
            TagId = ?;
    );
    `
	var statementContext *sql.Stmt

	if transaction != nil {
		statementContext, err = transaction.Prepare(stmt)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		statementContext, err = dbConn.Prepare(stmt)

		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = statementContext.Exec(bookmarkId, tagId)
	if err != nil {
		log.Fatal(err)
	}

	err = statementContext.Close()

	if err != nil {
		log.Fatal(err)
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
