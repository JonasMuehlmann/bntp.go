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

// Package libbookmarks implements functionality to work with bookmarks in a database context.
package libbookmarks

import (
	"errors"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/optional.go"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	"github.com/gocarina/gocsv"
)

func DeserializeBookmarks(serializedBookmarkList string) (bookmarkList []Bookmark, err error) {
	bookmarkList = make([]Bookmark, 0, 100)

	err = gocsv.UnmarshalString(serializedBookmarkList, &bookmarkList)
	if err != nil {
		err = helpers.DeserializationError{Inner: err}

		return
	}

	if len(bookmarkList) == 0 {
		err = helpers.IneffectiveOperationError{Inner: errors.New("Empty bookmark list")}

		return
	}

	return
}

func ImportBookmarks(dbConn *sqlx.DB, bookmarkList []Bookmark) error {
	transaction, err := dbConn.Beginx()
	if err != nil {
		return err
	}

	for _, bookmark := range bookmarkList {
		err := AddBookmark(nil, transaction, bookmark)

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

func SerializeBookmarks(bookmarkList []Bookmark) (serializedBookmarks string, err error) {
	if len(bookmarkList) == 0 {
		err = helpers.IneffectiveOperationError{Inner: errors.New("Empty bookmark list")}

		return
	}

	serializedBookmarks, err = gocsv.MarshalString(bookmarkList)

	return
}

func ExportBookmarks(dbConn *sqlx.DB) (bookmarkList []Bookmark, err error) {
	return GetBookmarks(dbConn, BookmarkFilter{})
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

	// REFACTOR: Extract function
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

	bookmarksBuffer := make([]Bookmark, 0, 200)

	err := dbConn.Select(&bookmarksBuffer, stmtBookmarks)

	if err != nil {
		return nil, err
	}

	for _, bookmark := range bookmarksBuffer {
		bookmark.Tags = make([]string, 0, 5)

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
		// REFACTOR: Extract AddingEmptyItemError
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

	// REFACTOR: Simplify with helpers
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
func RemoveType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ helpers.NameOrId) error {
	stmtName := `
        DELETE FROM
            BookmarkType
        WHERE
            Type = ?;
    `
	stmtID := `
        DELETE FROM
            BookmarkType
        WHERE
            Id = ?;
    `
	command, err := helpers.GetCommandFromIdentifier(type_, stmtName, stmtID)
	if err != nil {
		return err
	}

	statement, err := helpers.GetStatement(dbConn, transaction, command)
	if err != nil {
		return err
	}

	numAffectedRows, _, err := helpers.SqlExecuteStatement(statement, type_)
	if err != nil {
		return err
	}

	if numAffectedRows == 0 {
		return helpers.IneffectiveOperationError{errors.New("Type to delete does not exist")}
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

	// TODO: Benchmark to see if this is worth it
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

// AddBookmark adds a new bookmark to the DB.
// Passing a transaction is optional.
func AddBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, bookmark Bookmark) error {
	if strings.TrimSpace(bookmark.Url) == "" {
		return errors.New("Can't add bookmark without url")
	}

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
	var statement *sqlx.Stmt

	var err error

	// REFACTOR: Simplify with helpers
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

	// TODO: Handle importing tags
	var bookmarkTypeId int
	if bookmark.Type.HasValue {
		bookmarkTypeId, err = helpers.GetIdFromBookmarkType(dbConn, transaction, bookmark.Type.Wrappee)
		_, err = statement.Exec(bookmark.Title, bookmark.Url, bookmark.TimeAdded, bookmarkTypeId)
	} else {
		_, err = statement.Exec(bookmark.Title, bookmark.Url, bookmark.TimeAdded, nil)
	}

	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

func RemoveBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, bookmark helpers.NameOrId) error {
	stmtName := `
        DELETE FROM
            Bookmark
        WHERE
            Title = ?;
    `
	stmtID := `
        DELETE FROM
            Bookmark
        WHERE
            ID = ?;
    `
	command, err := helpers.GetCommandFromIdentifier(bookmark, stmtName, stmtID)
	if err != nil {
		return err
	}

	statement, err := helpers.GetStatement(dbConn, transaction, command)
	if err != nil {
		return err
	}

	numAffectedRows, _, err := helpers.SqlExecuteStatement(statement, bookmark)
	if err != nil {
		return err
	}

	if numAffectedRows == 0 {
		return helpers.IneffectiveOperationError{errors.New("Bookmark to delete does not exist")}
	}

	return nil
}

func EditBookmark(dbConn *sqlx.DB, transaction *sqlx.Tx, newData Bookmark) error {
	// TODO: Handle change of tags
	stmt := `
        UPDATE
            Bookmark
        SET
            Title = ?,
            Url = ?,
            TimeAdded = ?,
            BookmarkTypeId = ?,
            IsRead = ?,
            IsCollection = ?
        WHERE
            Id = ?;
    `

	var typeID optional.Optional[int]
	if newData.Type.HasValue {
		typeIDUnwrapped, err := helpers.GetIdFromBookmarkType(dbConn, transaction, newData.Type.Wrappee)
		if err != nil {
			return err
		}

		typeID.Wrappee = typeIDUnwrapped
	}

	_, _, err := helpers.SqlExecute(dbConn, transaction, stmt, newData.Title, newData.Url, newData.TimeAdded, typeID, newData.IsRead, newData.IsCollection, newData.Id)

	return err
}

// editBookmarkField sets column to newVal for the bookmark with the specified id.
// Passing a transaction is optional.
func editBookmarkField(dbConn *sqlx.DB, transaction *sqlx.Tx, identifier helpers.NameOrId, column string, newVal interface{}) error {
	stmtID := `
        UPDATE
            Bookmark
        SET
            ` + column + ` = ?
        WHERE
            Id = ?;
    `

	stmtName := `
        UPDATE
            Bookmark
        SET
            ` + column + ` = ?
        WHERE
            Title = ?;
    `

	command, err := helpers.GetCommandFromIdentifier(identifier, stmtName, stmtID)
	if err != nil {
		return err
	}

	statement, err := helpers.GetStatement(dbConn, transaction, command)
	if err != nil {
		return err
	}

	numAffectedRows, _, err := helpers.SqlExecuteStatement(statement, identifier, newVal)
	if err != nil {
		return err
	}

	if numAffectedRows == 0 {
		return helpers.IneffectiveOperationError{errors.New("Bookmark to edit does not exist")}
	}

	return nil
}

// EditIsRead sets IsRead to true for the bookmark with the specified id.
// Passing a transaction is optional.
func EditIsRead(dbConn *sqlx.DB, transaction *sqlx.Tx, id helpers.NameOrId, isRead bool) error {
	return editBookmarkField(dbConn, transaction, id, "IsRead", isRead)
}

// EditTitle sets Title to newTile for the bookmark with the specified id.
// Passing a transaction is optional.
func EditTitle(dbConn *sqlx.DB, transaction *sqlx.Tx, id helpers.NameOrId, newTitle string) error {
	return editBookmarkField(dbConn, transaction, id, "Title", newTitle)
}

// EditUrl sets Url to newUrl for the bookmark with the specified id.
// Passing a transaction is optional.
func EditUrl(dbConn *sqlx.DB, transaction *sqlx.Tx, id helpers.NameOrId, newUrl string) error {
	return editBookmarkField(dbConn, transaction, id, "Url", newUrl)
}

// EditType sets Type to newType for the bookmark with the specified id.
// Passing a transaction is optional.
func EditType(dbConn *sqlx.DB, transaction *sqlx.Tx, id helpers.NameOrId, newType string) error {
	typeId, err := helpers.GetIdFromType(dbConn, transaction, newType)
	if err != nil {
		return err
	}
	return editBookmarkField(dbConn, transaction, id, "BookmarkTypeId", typeId)
}

// EditIsCollection sets isCollection to isCollection for the bookmark with the specified id.
// Passing a transaction is optional.
func EditIsCollection(dbConn *sqlx.DB, transaction *sqlx.Tx, id helpers.NameOrId, isCollection bool) error {
	return editBookmarkField(dbConn, transaction, id, "IsCollection", isCollection)
}

// AddTag adds a tag newTag to the bookmark with bookmarkId.
// Passing a transaction is optional.
func AddTag(dbConn *sqlx.DB, transaction *sqlx.Tx, bookmark helpers.NameOrId, newTag helpers.NameOrId) error {
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

	// REFACTOR: Simplify with helpers
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

	var tagId int
	var bookmarkId int
	switch newTag.(type) {
	case string:
		tagId, err = helpers.GetIdFromTag(dbConn, transaction, newTag.(string))
	case int:
		tagId = newTag.(int)
	default:
		err = helpers.InvalidIdentifierError{newTag}
	}
	if err != nil {
		return err
	}

	switch bookmark.(type) {
	case string:
		bookmarkId, err = helpers.GetIdFromBookmark(dbConn, transaction, bookmark.(string))
	case int:
		bookmarkId = bookmark.(int)
	default:
		err = helpers.InvalidIdentifierError{bookmark}
	}
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
func RemoveTag(dbConn *sqlx.DB, transaction *sqlx.Tx, bookmark helpers.NameOrId, tag helpers.NameOrId) error {
	stmt := `
        DELETE FROM
            BookmarkContext
        WHERE
            BookmarkId = ?
            AND
            TagId = ?;
    );
    `

	var err error

	statement, err := helpers.GetStatement(dbConn, transaction, stmt)
	if err != nil {
		return err
	}

	var tagId int
	var bookmarkId int
	switch tag.(type) {
	case string:
		tagId, err = helpers.GetIdFromTag(dbConn, transaction, tag.(string))
	case int:
		tagId = tag.(int)
	default:
		err = helpers.InvalidIdentifierError{tag}
	}
	if err != nil {
		return err
	}

	switch bookmark.(type) {
	case string:
		bookmarkId, err = helpers.GetIdFromBookmark(dbConn, transaction, bookmark.(string))
	case int:
		bookmarkId = bookmark.(int)
	default:
		err = helpers.InvalidIdentifierError{bookmark}
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
