// Package libdocuments implements functionality to work with documents in a database and file system context.
package libdocuments

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/jmoiron/sqlx"
)

// AddTag adds a tag to the tag line of the document at documentPath.
func AddTag(documentPath string, tag string) error {
	if tag == "" {
		return errors.New("Can't add empty tag")
	}

	tagsLineNumber, _, err := FindTagsLine(documentPath)
	if err != nil {
		return err
	}

	fileBuffer, err := os.ReadFile(documentPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(fileBuffer), "\n")

	// If the tag line is the last line, we append a new one.
	if len(lines) == tagsLineNumber {
		lines = append(lines, tag)
	} else {
		lines[tagsLineNumber] += "," + tag
	}

	err = os.WriteFile(documentPath, []byte(strings.Join(lines, "\n")), 0o644)
	if err != nil {
		return err
	}

	return nil
}

// RemoveTag removes a tag from the tag line of the document at documentPath.
func RemoveTag(documentPath string, tag string) error {
	if tag == "" {
		return errors.New("Can't remove empty tag")
	}

	tagsLineNumber, tags, err := FindTagsLine(documentPath)
	if err != nil {
		return err
	}

	tags = strings.Replace(tags, tag, "", -1)
	tags = strings.Replace(tags, ",,", ",", -1)

	fileBuffer, err := os.ReadFile(documentPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(fileBuffer), "\n")
	lines[tagsLineNumber] = tags

	err = os.WriteFile(documentPath, []byte(strings.Join(lines, "\n")), 0o644)
	if err != nil {
		return err
	}

	return nil
}

// Rename renames a tag oldTag to newTag in the tag line of the doucment at documentPath.
// This method preserves the order of all tags in the doucment.
func RenameTag(documentPath string, oldTag string, newTag string) error {
	if oldTag == "" {
		return errors.New("Can't rename from empty Tag")
	}

	if newTag == "" {
		return errors.New("Can't rename tp empty Tag")
	}

	tagsLineNumber, tags, err := FindTagsLine(documentPath)
	if err != nil {
		return err
	}

	tags = strings.Replace(tags, oldTag, newTag, -1)

	fileBuffer, err := os.ReadFile(documentPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(fileBuffer), "\n")
	lines[tagsLineNumber] = tags

	err = os.WriteFile(documentPath, []byte(strings.Join(lines, "\n")), 0o644)
	if err != nil {
		return err
	}

	return nil
}

// GetTags returns all tags contained in the document at documentPath.
func GetTags(documentPath string) ([]string, error) {
	_, tags, err := FindTagsLine(documentPath)
	if err != nil {
		return nil, err
	}

	return strings.Split(tags, ","), nil
}

// FindTagsLine finds the line in documentPath which contains it's tags.
// It returns the 0 based line lumber of the tags line as well as the line itself.
func FindTagsLine(documentPath string) (int, string, error) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)
	if err != nil {
		return 0, "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0

	for scanner.Scan() {
		if scanner.Text() == "# Tags" {
			scanner.Scan()

			return i + 1, scanner.Text(), nil
		}
		i++
	}

	return 0, "", errors.New("Could not find tags line")
}

// HasTags checks if the document at documentPath has all specified tags.
func HasTags(documentPath string, tags []string) (bool, error) {
	documentTags, err := GetTags(documentPath)
	if err != nil {
		return false, err
	}

	for _, tag := range tags {
		for _, documentTag := range documentTags {
			if tag == documentTag {
				continue
			}

			return false, nil
		}
	}

	return true, nil
}

// TODO: Refactor to search in DB not FS
// FindDocumentsWithTags returns all paths to doucments which have all specified tags.
func FindDocumentsWithTags(rootDir string, tags []string) ([]string, error) {
	filesWithTags := make([]string, 0, 100)

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			hasTags, err := HasTags(path, tags)
			if err != nil {
				return err
			}
			if !hasTags {
				return nil
			}
			filesWithTags = append(filesWithTags, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return filesWithTags, nil
}

// FindLinksLines finds the lines in documentPath in which links to other documents are listed.
// It returns the range of line numbers containing links as well as the lines themselves.
func FindLinksLines(documentPath string) (int, int, []string, error) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)
	if err != nil {
		return 0, 0, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumberFirstLink := -1
	lineNumberLastLink := -1
	links := make([]string, 0, 10)

	i := 0

	for scanner.Scan() {
		if scanner.Text() == "# Links" {
			lineNumberFirstLink = i + 1

			break
		}
		i++
	}

	for scanner.Scan() && strings.HasPrefix(scanner.Text(), "- ") {
		links[i-lineNumberFirstLink] = scanner.Text()
		i++
	}

	lineNumberLastLink = i

	return lineNumberFirstLink, lineNumberLastLink, links, nil
}

// FindBacklinksLines finds the lines in documentPath in which backlinks to other documents are listed.
// It returns the range of line numbers containing backlinks as well as the lines themselves.
func FindBacklinksLines(documentPath string) (int, int, []string, error) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)
	if err != nil {
		return 0, 0, nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNumberFirstLink := -1
	lineNumberLastLink := -1
	links := make([]string, 0, 10)

	i := 0

	for scanner.Scan() {
		if scanner.Text() == "# Backlinks" {
			lineNumberFirstLink = i + 1

			break
		}
		i++
	}

	for scanner.Scan() && strings.HasPrefix(scanner.Text(), "- ") {
		links[i-lineNumberFirstLink] = scanner.Text()
		i++
	}

	lineNumberLastLink = i

	return lineNumberFirstLink, lineNumberLastLink, links, nil
}

// AddLink adds a link to documentPathDestination into the document at documentPathSource.
func AddLink(documentPathSource string, documentPathDestination string) error {
	// lineNumberFirstLink, lineNumberLastLink, links, err := FindLinksLines(documentPathSource)
	lineNumberFirstLink, _, links, err := FindLinksLines(documentPathSource)
	if err != nil {
		return err
	}

	links = append(links, documentPathDestination)

	file, err := os.OpenFile(documentPathSource, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}

	defer file.Close()

	offset, err := file.Seek(int64(lineNumberFirstLink), io.SeekStart)
	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(strings.Join(links, "\n")), offset)

	if err != nil {
		return err
	}

	return nil
}

// RemoveLink removes the link to documentPathDestination from the document at documentPathSource.
func RemoveLink(documentPathSource string, documentPathDestination string) error {
	// lineNumberFirstLink, lineNumberLastLink, linksOrig, err := FindLinksLines(documentPathSource)
	lineNumberFirstLink, _, linksOrig, err := FindLinksLines(documentPathSource)
	if err != nil {
		return err
	}

	iLinkToDelete := -1

	for i, link := range linksOrig {
		if link == documentPathSource {
			iLinkToDelete = i
		}
	}

	links := make([]string, 0, 10)

	links = append(links, linksOrig[:iLinkToDelete]...)
	links = append(links, linksOrig[iLinkToDelete+1:]...)

	file, err := os.OpenFile(documentPathDestination, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	offset, err := file.Seek(int64(lineNumberFirstLink), io.SeekStart)
	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(strings.Join(links, "\n")), offset)

	if err != nil {
		return err
	}

	return err
}

// AddBacklink adds a Backlink to documentPathSource into the document at documentPathDestination.
func AddBacklink(documentPathDestination string, documentPathSource string) error {
	// lineNumberFirstLink, lineNumberLastLink, links, err := FindBacklinksLines(documentPathSource)
	lineNumberFirstLink, _, links, err := FindBacklinksLines(documentPathSource)
	if err != nil {
		return err
	}

	links = append(links, documentPathSource)

	file, err := os.OpenFile(documentPathDestination, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}

	defer file.Close()

	offset, err := file.Seek(int64(lineNumberFirstLink), io.SeekStart)
	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(strings.Join(links, "\n")), offset)

	if err != nil {
		return err
	}

	return nil
}

// RemoveBacklink removes the backlink to documentPathSource from the document at documentPathDestination.
func RemoveBacklink(documentPathDestination string, documentPathSource string) error {
	// lineNumberFirstLink, lineNumberLastLink, linksOrig, err := FindBacklinksLines(documentPathSource)
	lineNumberFirstLink, _, linksOrig, err := FindBacklinksLines(documentPathSource)
	if err != nil {
		return err
	}

	iLinkToDelete := -1

	for i, link := range linksOrig {
		if link == documentPathSource {
			iLinkToDelete = i
		}
	}

	links := make([]string, 0, 10)

	links = append(links, linksOrig[:iLinkToDelete]...)
	links = append(links, linksOrig[iLinkToDelete+1:]...)

	file, err := os.OpenFile(documentPathDestination, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	offset, err := file.Seek(int64(lineNumberFirstLink), io.SeekStart)
	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(strings.Join(links, "\n")), offset)

	if err != nil {
		return err
	}

	return err
}

// AddDocument adds a new document located at documentPath to the DB.
func AddDocument(dbConn *sqlx.DB, transaction *sqlx.Tx, documentPath string, documentType string) error {
	stmt := `
        INSERT INTO
            Document(
                Path,
                DocumentTypeId
            )
        VALUES(
            ?,
            ?
        );
    `

	documentTypeId, err := helpers.GetIdFromDocumentType(dbConn, transaction, documentType)
	if err != nil {
		return err
	}

	return helpers.SqlExecute(dbConn, transaction, stmt, documentPath, documentTypeId)
}

//  RemoveDocument removes  a document located at documentPath from the DB.
func RemoveDocument(dbConn *sqlx.DB, transaction *sqlx.Tx, documentPath string) error {
	stmt := `
        DELETE FROM
            Document
        WHERE
            Path = ?;
    `

	return helpers.SqlExecute(dbConn, transaction, stmt, documentPath)
}

// RenameDocument moves a document located at documentPathOld to documentPathNew.
func RenameDocument(dbConn *sqlx.DB, transaction *sqlx.Tx, documentPathOld string, documentPathNew string) error {
	stmt := `
        UPDATE
            Document
        SET
            Path = ?
        WHERE
            Path = ?;
    `

	return helpers.SqlExecute(dbConn, transaction, stmt, documentPathNew, documentPathOld)
}

// ChangeDocumentType changes the type of the document located documentPath to documentTypeNew.
func ChangeDocumentType(dbConn *sqlx.DB, transaction *sqlx.Tx, documentPath string, documentTypeNew string) error {
	stmt := `
        UPDATE
            Document
        SET
            DocumentTypeId = ?
        WHERE
            Path = ?;
    `

	documentTypeId, err := helpers.GetIdFromDocumentType(dbConn, transaction, documentTypeNew)
	if err != nil {
		return err
	}
	if documentTypeId == -1 {
		return errors.New("Could not retrieve DocumentTypeId")
	}

	return helpers.SqlExecute(dbConn, transaction, stmt, documentTypeId, documentPath)
}

// AddType makes a new DocumentType available for use in the DB.
// Passing a transaction is optional.
func AddType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) error {
	if type_ == "" {
		return errors.New("Can't add empty tag")
	}

	stmt := `
        INSERT INTO
            DocumentType(
                DocumentType
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

// RemoveType removes an available DocumentType from the DB.
// Passing a transaction is optional.
func RemoveType(dbConn *sqlx.DB, transaction *sqlx.Tx, type_ string) error {
	stmt := `
        DELETE FROM
            DocumentType
        WHERE
            DocumentType = ?;
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
