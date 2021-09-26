package libbookmarks

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

// ImportMinimalCSV reads a csv file at csvPath and imports it into the bookmark DB.
// The csv file is expected to have the columns "Text" and "Url" ONLY.
func ImportMinimalCSV(dbConn *sqlx.DB, csvPath string) error {
	file, err := os.Open(csvPath)
	if err != nil {
		return err
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	bookmarks, err := reader.ReadAll()
	if err != nil {
		return err
	}

	header := bookmarks[0]

	if len(header) != 2 {
		return errors.New("CSV Header does not have correct number of fields. It should have 2.")
	}

	if !(header[0] == "Title" || header[1] == "Title") || !(header[0] == "Url" || header[1] == "Url") || header[0] == header[1] {
		return errors.New("CSV Header does not have necessary fields 'Title' and 'Url.'")
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
		return err
	}

	for _, bookmark := range bookmarks {
		AddBookmark(dbConn, transaction, bookmark[titleColumn], bookmark[linkColumn], 1, false)
	}

	err = transaction.Commit()

	if err != nil {
		return err
	}

	return nil
}

// ExportCSV exports an array of bookmarks to a CSV file at csvPath.
func ExportCSV(bookmarks []Bookmark, csvPath string) error {
	var writer *csv.Writer

	if csvPath != "" { // 0664 UNIX Permission code
		file, err := os.OpenFile(csvPath, os.O_CREATE|os.O_WRONLY, 0o664)
		if err != nil {
			return err
		}

		defer file.Close()

		writer = csv.NewWriter(file)
	} else { // 0664 UNIX Permission code
		writer = csv.NewWriter(os.Stdout)
	}

	writer.Comma = ';'

	csvHeader := make([]string, 0, 10)

	tempBookmark := &Bookmark{}
	bookmarkReflected := reflect.ValueOf(tempBookmark).Elem()

	for i := 0; i < bookmarkReflected.NumField(); i++ {
		csvHeader = append(csvHeader, bookmarkReflected.Type().Field(i).Name)
	}

	err := writer.Write(csvHeader)
	if err != nil {
		return err
	}

	rowsBuffer := make([][]string, 0, len(bookmarks))
	for i := range rowsBuffer {
		rowsBuffer[i] = make([]string, 0, 10)
	}

	for i, bookmark := range bookmarks {
		rowsBuffer[i] = []string{
			strconv.Itoa(bookmark.Id),
			bookmark.Title,
			bookmark.Url,
			strconv.FormatBool(bookmark.IsCollection),
			bookmark.Type,
			strings.Join(bookmark.Tags, ","),
			bookmark.TimeAdded,
			strconv.FormatBool(bookmark.IsRead),
		}
	}

	err = writer.WriteAll(rowsBuffer)
	if err != nil {
		return err
	}

	return nil
}

// GetBookmarks returns all bookmarks stored in the DB which satisfy the given filter.
func GetBookmarks(dbConn *sqlx.DB, filter BookmarkFilter) ([]Bookmark, error) {
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

	stmtBookmarks = ApplyBookmarkFilters(stmtBookmarks, filter)

	stmtNumberOfBookmarks := "SELECT COUNT(*) FROM Bookmark INNER JOIN Type ON Bookmark.TypeId = Type.Id"

	stmtNumberOfBookmarks = ApplyBookmarkFilters(stmtNumberOfBookmarks, filter)

	stmtTags := `
        SELECT
            Tag.Tag
        FROM
            Tag
        INNER JOIN Context ON
            Context.TagId = Tag.Id
        WHERE Context.BookmarkId = ?;`

	stmtNumberOfTags := "SELECT COUNT(*) FROM  Context WHERE BookmarkId = ?;"

	var numberOfBookmarks int

	err := dbConn.Get(numberOfBookmarks, stmtNumberOfBookmarks, nil)
	if err != nil {
		return nil, errors.New("Could not count number of bookmarks")
	}

	bookmarksBuffer := make([]Bookmark, 0, numberOfBookmarks)

	err = dbConn.Select(bookmarksBuffer, stmtBookmarks)

	if err != nil {
		return nil, errors.New("Could not select bookmarks")
	}

	var numberOfTags int

	for _, bookmark := range bookmarksBuffer {
		err := dbConn.Get(numberOfTags, stmtNumberOfTags, bookmark.Id)
		if err != nil {
			return nil, errors.New("Could not read bookmark")
		}

		bookmark.Tags = make([]string, 0, 10)

		err = dbConn.Select(bookmark.Tags, stmtTags, bookmark.Id)

		if err != nil {
			return nil, errors.New("Could not read bookmark's tags")
		}
	}

	return bookmarksBuffer, nil
}

// AddType makes a new BookmarkType available for use in the DB.
// Passing a transaction is optional.
func AddType(dbConn *sqlx.DB, transaction *sql.Tx, type_ string) error {
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
			return err
		}
	} else {
		statement, err = dbConn.Prepare(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statement.Exec(type_)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// RemoveType removes an available BookmarkType from the DB.
// Passing a transaction is optional.
func RemoveType(dbConn *sqlx.DB, transaction *sql.Tx, type_ string) error {
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
			return err
		}
	} else {
		statement, err = dbConn.Prepare(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statement.Exec(type_)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// ListTypes lists all BookmarkTypes available in the DB.
func ListTypes(dbConn *sqlx.DB) ([]string, error) {
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

	err := countRow.Scan(&rowCount)
	if err != nil {
		return nil, err
	}

	rows, err := dbConn.Query(stmtRows)
	if err != nil {
		return nil, err
	}

	types := make([]string, 0, rowCount)

	i := 0
	for rows.Next() {
		err := rows.Scan(&types[i])
		if err != nil {
			return nil, err
		}
		i++
	}

	return types, nil
}

// TODO: Allow passing string for type_
// AddBookmark adds a new bookmark to the DB.
// Passing a transaction is optional.
func AddBookmark(dbConn *sqlx.DB, transaction *sql.Tx, title string, url string, type_ int, isCollection bool) error {
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
			return err
		}
	} else {
		statement, err = dbConn.Prepare(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statement.Exec(title, url, time.Now().Format("2006-01-02"), type_, isCollection)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// EditBookmark sets column to newVal for the bookmark with the specified id.
// Passing a transaction is optional.
func EditBookmark(dbConn *sqlx.DB, transaction *sql.Tx, id int, column string, newVal interface{}) error {
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
			return err
		}
	} else {
		statement, err = dbConn.Prepare(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statement.Exec(column, newVal, id)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// MarkAsRead sets IsRead to true for the bookmark with the specified id.
// Passing a transaction is optional.
func MarkAsRead(dbConn *sqlx.DB, transaction *sql.Tx, id int) error {
	return EditBookmark(dbConn, transaction, id, "IsRead", true)
}

// EditTitle sets Title to newTile for the bookmark with the specified id.
// Passing a transaction is optional.
func EditTitle(dbConn *sqlx.DB, transaction *sql.Tx, id int, newTitle string) error {
	return EditBookmark(dbConn, transaction, id, "Title", newTitle)
}

// EditUrl sets Url to newUrl for the bookmark with the specified id.
// Passing a transaction is optional.
func EditUrl(dbConn *sqlx.DB, transaction *sql.Tx, id int, newUrl string) error {
	return EditBookmark(dbConn, transaction, id, "Url", newUrl)
}

// EditType sets Type to newType for the bookmark with the specified id.
// Passing a transaction is optional.
func EditType(dbConn *sqlx.DB, transaction *sql.Tx, id int, newType string) error {
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
			return err
		}
	} else {
		statement, err = dbConn.Prepare(stmt)

		if err != nil {
			return err
		}
	}

	typeRow := statement.QueryRow(newType)

	if err != nil {
		return err
	}

	err = typeRow.Scan(&typeId)

	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return EditBookmark(dbConn, transaction, id, "TypeId", typeId)
}

// EditIsCollection sets isCollection to isCollection for the bookmark with the specified id.
// Passing a transaction is optional.
func EditIsCollection(dbConn *sqlx.DB, transaction *sql.Tx, id int, isCollection bool) error {
	return EditBookmark(dbConn, transaction, id, "IsCollection", isCollection)
}

// AddTag adds a tag newTag to the bookmark with bookmarkId.
// Passing a transaction is optional.
func AddTag(dbConn *sqlx.DB, transaction *sql.Tx, bookmarkId int, newTag string) error {
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
			return err
		}
	} else {
		statementTag, err = dbConn.Prepare(stmtTag)

		if err != nil {
			return err
		}
	}

	tagRow := statementTag.QueryRow(newTag)

	if err != nil {
		return err
	}

	err = tagRow.Scan(&tagId)

	if err != nil {
		return err
	}

	err = statementTag.Close()

	if err != nil {
		return err
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
			return err
		}
	} else {
		statementContext, err = dbConn.Prepare(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statementContext.Exec(bookmarkId, tagId)
	if err != nil {
		return err
	}

	err = statementContext.Close()

	if err != nil {
		return err
	}

	return nil
}

// RemoveTag removes a tag tag_ from the bookmark with bookmarkId.
// Passing a transaction is optional.
func RemoveTag(dbConn *sqlx.DB, transaction *sql.Tx, bookmarkId int, tag_ string) error {
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
			return err
		}
	} else {
		statementTag, err = dbConn.Prepare(stmtTag)

		if err != nil {
			return err
		}
	}

	tagRow := statementTag.QueryRow(tag_)

	if err != nil {
		return err
	}

	err = tagRow.Scan(&tagId)

	if err != nil {
		return err
	}

	err = statementTag.Close()

	if err != nil {
		return err
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
			return err
		}
	} else {
		statementContext, err = dbConn.Prepare(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statementContext.Exec(bookmarkId, tagId)
	if err != nil {
		return err
	}

	err = statementContext.Close()

	if err != nil {
		return err
	}

	return nil
}
