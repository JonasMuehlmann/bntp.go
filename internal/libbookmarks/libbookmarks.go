package libbookmarks

import (
	"database/sql"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

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

	err := countRow.Scan(&rowCount)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := dbConn.Query(stmtRows)
	if err != nil {
		log.Fatal(err)
	}

	types := make([]string, 0, rowCount)

	i := 0
	for rows.Next() {
		err := rows.Scan(&types[i])
		if err != nil {
			log.Fatal(err)
		}
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

	err = typeRow.Scan(&typeId)

	if err != nil {
		log.Fatal(err)
	}

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

	err = tagRow.Scan(&tagId)

	if err != nil {
		log.Fatal(err)
	}

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

	err = tagRow.Scan(&tagId)

	if err != nil {
		log.Fatal(err)
	}

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
