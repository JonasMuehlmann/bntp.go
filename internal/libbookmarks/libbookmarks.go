// Package libbookmarks implements functionality to work with bookmarks in a database context.
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

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
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

	if len(bookmarks) < 2 {
		return errors.New("CSV does not contain at least one entry")
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

	transaction, err := dbConn.Beginx()
	if err != nil {
		return err
	}

	for _, bookmark := range bookmarks[1:] {
		err := AddBookmark(dbConn, transaction, bookmark[titleColumn], bookmark[linkColumn], sql.NullInt32{Valid: false})

		if err != nil {
			return err
		}
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

	if len(bookmarks) == 0 {
		return errors.New("No bookmarks to export")
	}

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

	rowsBuffer := make([][]string, len(bookmarks))

	for i, bookmark := range bookmarks {
		var isCollectionOut string

		if bookmark.IsCollection.Valid {
			isCollectionOut = strconv.FormatBool(bookmark.IsCollection.Bool)
		} else {
			isCollectionOut = "NULL"
		}

		rowsBuffer[i] = []string{
			bookmark.Title.String,
			bookmark.Url,
			bookmark.TimeAdded,
			bookmark.Type.String,
			strings.Join(bookmark.Tags, ","),
			strconv.Itoa(bookmark.Id),
			strconv.FormatBool(bookmark.IsRead),
			isCollectionOut,
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
            Bookmark.Id AS Id,
            Bookmark.IsRead AS IsRead,
            Bookmark.Title AS Title,
            Bookmark.Url AS Url,
            Bookmark.TimeAdded AS TimeAdded,
            BookmarkType.Type AS Type,
            Bookmark.IsCollection AS IsCollection
        FROM
            Bookmark
        LEFT JOIN BookmarkType ON
            BookmarkType.Id = Bookmark.BookmarkTypeId
    `

	stmtBookmarks = ApplyBookmarkFilters(stmtBookmarks, filter)

	stmtTags := `
        SELECT
            Tag.Tag
        FROM
            Tag
        INNER JOIN BookmarkContext ON
            BookmarkContext.TagId = Tag.Id
        WHERE
            BookmarkContext.BookmarkId = ?;
    `

	bookmarksBuffer := []Bookmark{}

	err := dbConn.Select(&bookmarksBuffer, stmtBookmarks)

	if err != nil {
		return nil, err
	}

	for _, bookmark := range bookmarksBuffer {
		bookmark.Tags = []string{}

		err = dbConn.Select(&bookmark.Tags, stmtTags, bookmark.Id)

		if err != nil {
			return nil, err
		}
	}

	return bookmarksBuffer, nil
}

// AddType makes a new BookmarkType available for use in the DB.
// Passing a transaction is optional.
func AddType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) error {
	if type_ == "" {
		return errors.New("Can't add empty tag")
	}

	stmt := `
        INSERT INTO
            BookmarkType(
                Type
            )
        VALUES(
            ?
        );
    `
	var statement *sqlx.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

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
func RemoveType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) error {
	stmt := `
        DELETE FROM
            BookmarkType
        WHERE
            Type = ?;
    `
	var statement *sqlx.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	result, err := statement.Exec(type_)
	if err != nil {
		return err
	}

	numAffectedRows, err := result.RowsAffected()
	if numAffectedRows == 0 || err != nil {
		return errors.New("Type to be deleted does not exist")
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// ListTypes lists all BookmarkTypes available in the DB.
func ListTypes(dbConn *sqlx.DB) ([]string, error) {
	stmtTypes := `
        SELECT
            Type
        FROM
            BookmarkType;
    `

	stmtNumTypes := `
        SELECT
            Count(*)
        FROM
            BookmarkType;
    `

	var numTypes int

	err := dbConn.Get(&numTypes, stmtNumTypes)
	if err != nil {
		return nil, err
	}

	types := make([]string, 0, numTypes)

	err = dbConn.Select(&types, stmtTypes)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// TODO: Allow passing string for type_

// AddBookmark adds a new bookmark to the DB.
// Passing a transaction is optional.
func AddBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, title string, url string, type_ sql.NullInt32) error {
	stmt := `
        INSERT INTO
            Bookmark(
                Title,
                Url,
                TimeAdded,
                BookmarkTypeId
            )
        VALUES(
            ?,
            ?,
            ?,
            ?
        );
    `
	if title == "" || url == "" {
		return errors.New("Entry is missing a column")
	}

	var statement *sqlx.Stmt

	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	_, err = statement.Exec(title, url, time.Now().Format("2006-01-02"), type_)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// TODO: Implement
func RemoveBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, ID int) error {
	return nil
}

func EditBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, newData Bookmark) error {
	stmt := `
        UPDATE
            Bookmark
        SET
            Title = :Title,
            Url = :Url
            TimeAdded = :TimeAdded
            Type = :Type
            Id = :Id
            IsRead = :IsRead
            IsCollection = :IsCollection
        WHERE
            Id = ?;
    `

	_, _, err := helpers.SqlExecuteNamed(dbConn, transaction, stmt, newData)

	return err
}

// editBookmarkField sets column to newVal for the bookmark with the specified id.
// Passing a transaction is optional.
func editBookmarkField(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, column string, newVal interface{}) error {
	stmt := `
        UPDATE
            Bookmark
        SET
            ` + column + ` = ?
        WHERE
            Id = ?;
    `

	_, _, err := helpers.SqlExecute(dbConn, transaction, stmt, newVal, id)

	return err
}

// EditIsRead sets IsRead to true for the bookmark with the specified id.
// Passing a transaction is optional.
func EditIsRead(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, isRead bool) error {
	return editBookmarkField(dbConn, transaction, id, "IsRead", isRead)
}

// EditTitle sets Title to newTile for the bookmark with the specified id.
// Passing a transaction is optional.
func EditTitle(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, newTitle string) error {
	return editBookmarkField(dbConn, transaction, id, "Title", newTitle)
}

// EditUrl sets Url to newUrl for the bookmark with the specified id.
// Passing a transaction is optional.
func EditUrl(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, newUrl string) error {
	return editBookmarkField(dbConn, transaction, id, "Url", newUrl)
}

// EditType sets Type to newType for the bookmark with the specified id.
// Passing a transaction is optional.
func EditType(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, newType string) error {
	typeId, err := helpers.GetIdFromType(dbConn, transaction, newType)
	if err != nil {
		return err
	}
	return editBookmarkField(dbConn, transaction, id, "BookmarkTypeId", typeId)
}

// EditIsCollection sets isCollection to isCollection for the bookmark with the specified id.
// Passing a transaction is optional.
func EditIsCollection(dbConn *sqlx.DB, transaction *sqlx.Tx, id int, isCollection bool) error {
	return editBookmarkField(dbConn, transaction, id, "IsCollection", isCollection)
}

// AddTag adds a tag newTag to the bookmark with bookmarkId.
// Passing a transaction is optional.
func AddTag(dbConn *sqlx.DB, transaction *sqlx.Tx, bookmarkId int, newTag string) error {
	stmt := `
        INSERT INTO
            BookmarkContext(BookmarkId, TagId)
        VALUES(
            ?,
            ?
    );
    `

	var statementContext *sqlx.Stmt
	var err error

	if transaction != nil {
		statementContext, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statementContext, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	tagId, err := helpers.GetIdFromTag(dbConn, transaction, newTag)
	if err != nil {
		return err
	}

	result, err := statementContext.Exec(bookmarkId, tagId)
	if err != nil {
		return err
	}

	numAffectedRows, err := result.RowsAffected()
	if numAffectedRows == 0 || err != nil {
		return errors.New("Type to be deleted does not exist")
	}

	err = statementContext.Close()

	if err != nil {
		return err
	}

	return nil
}

// RemoveTag removes a tag tag_ from the bookmark with bookmarkId.
// Passing a transaction is optional.
func RemoveTag(dbConn *sqlx.DB, transaction *sqlx.Tx, bookmarkId int, tag_ string) error {
	stmt := `
        DELETE FROM
            BookmarkContext
        WHERE
            BookmarkId = ?
            AND
            TagId = ?;
    );
    `

	var statement *sqlx.Stmt
	var err error

	if transaction != nil {
		statement, err = transaction.Preparex(stmt)

		if err != nil {
			return err
		}
	} else {
		statement, err = dbConn.Preparex(stmt)

		if err != nil {
			return err
		}
	}

	tagId, err := helpers.GetIdFromTag(dbConn, transaction, tag_)
	if err != nil {
		return err
	}

	_, err = statement.Exec(bookmarkId, tagId)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}
