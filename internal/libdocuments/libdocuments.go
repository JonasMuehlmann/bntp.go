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

// Package libdocuments implements functionality to work with documents in a database and file system context.
package libdocuments

import (
	"bufio"
	"errors"
	"os"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/jmoiron/sqlx"
)

// TODO: Allow passing string and id for document and tag

// AddTag adds a tag newTag to the document at documentPath.
// Passing a transaction is optional.
func AddTag(dbConn *sqlx.DB, transaction *sqlx.Tx, documentPath string, newTag string) error {
	stmt := `
        INSERT INTO
            DocumentContext(DocumentId, TagId)
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

	documentId, err := helpers.GetIdFromDocument(dbConn, transaction, documentPath)
	if err != nil {
		return err
	}

	result, err := statementContext.Exec(documentId, tagId)
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

// TODO: Allow passing string and id for document and tag

// RemoveTag removes a tag tag_ from the document at documentPath
// Passing a transaction is optional.
func RemoveTag(dbConn *sqlx.DB, transaction *sqlx.Tx, documentPath string, tag_ string) error {
	stmt := `
        DELETE FROM
            DocumentContext
        WHERE
            DocumentId = ?
            AND
            TagId = ?;
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

	documentId, err := helpers.GetIdFromDocument(dbConn, transaction, documentPath)
	if err != nil {
		return err
	}

	_, err = statement.Exec(documentId, tagId)
	if err != nil {
		return err
	}

	err = statement.Close()

	if err != nil {
		return err
	}

	return nil
}

// TODO: Allow passing string and id for document and tag

// AddTagToFile adds a tag to the tag line of the document at documentPath.
func AddTagToFile(documentPath string, tag string) error {
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

// TODO: Allow passing string and id for document and tag

// RemoveTagFromFile removes a tag from the tag line of the document at documentPath.
func RemoveTagFromFile(documentPath string, tag string) error {
	if tag == "" {
		return errors.New("Can't remove empty tag")
	}

	tagsLineNumber, tags, err := FindTagsLine(documentPath)
	if err != nil {
		return err
	}

	// PERF: This does unnecessary work
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

// TODO: Allow passing string and id for document and tag

// RenameTagInFile renames a tag oldTag to newTag in the tag line of the doucment at documentPath.
// This method preserves the order of all tags in the doucment.
func RenameTagInFile(documentPath string, oldTag string, newTag string) error {
	if oldTag == "" {
		return errors.New("Can't rename from empty Tag")
	}

	if newTag == "" {
		return errors.New("Can't rename to empty Tag")
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

// TODO: Allow passing string and id for document

// GetTags returns all tags contained in the document at documentPath.
func GetTags(documentPath string) ([]string, error) {
	_, tags, err := FindTagsLine(documentPath)
	if err != nil {
		return nil, err
	}

	tagsBuffer := strings.Split(tags, ",")

	if len(tagsBuffer) == 1 && tagsBuffer[0] == "" {
		return nil, errors.New("No tags found")
	}

	return tagsBuffer, nil
}

// TODO: Allow passing string and id for document

// FindTagsLine finds the line in documentPath which contains it's tags.
// It returns the 0 based line lumber of the tags line as well as the line itself.
func FindTagsLine(documentPath string) (lineNumber int, tagsLine string, err error) {
	file, err := os.OpenFile(documentPath, os.O_RDONLY, 0o644)
	if err != nil {
		return 0, "", err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	i := 0

	// NOTE: Are the named parameters handled correctly?
	// REFACTOR: This could probably be simplified with goaoi
	for scanner.Scan() {
		if scanner.Text() == "# Tags" {
			scanner.Scan()

			return i + 1, scanner.Text(), nil
		}
		i++
	}

	return 0, "", errors.New("Could not find tags line")
}

// TODO: Allow passing string and id for document and tag

// HasTags checks if the document at documentPath has all specified tags.
func HasTags(documentPath string, tags []string) (bool, error) {
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == "") {
		return false, errors.New("No input tags")
	}

	documentTags, err := GetTags(documentPath)
	if err != nil {
		return false, err
	}

	// REFACTOR: This could probably be simplified with goaoi
	for _, tag := range tags {
		hasTag := false

		for _, documentTag := range documentTags {
			if tag == documentTag {
				hasTag = true

				break
			}
		}

		if !hasTag {
			return false, nil
		}
	}

	return true, nil
}

// TODO: Allow passing string and id for tag

// FindDocumentsWithTags returns all paths to doucments which have all specified tags.
func FindDocumentsWithTags(dbConn *sqlx.DB, tags []string) ([]string, error) {
	if len(tags) == 0 || (len(tags) == 1 && tags[0] == "") {
		return nil, errors.New("No input tags")
	}
	stmtTags := `
       SELECT
           Path
       FROM
           Document
       LEFT JOIN DocumentContext ON
           DocumentId = Document.Id
       WHERE
           TagId IN (SELECT Id FROM Tag WHERE Tag IN(?))
       GROUP BY
           Document.Id
       HAVING
           COUNT(*) =`

	stmtNumPaths := `
       SELECT
           COUNT(*)
       FROM
           Document
       LEFT JOIN DocumentContext ON
           DocumentId = Document.Id
       WHERE
           TagId IN (SELECT Id FROM Tag WHERE Tag IN(?))
       GROUP BY
           Document.Id
       HAVING
           COUNT(*) =`

	var numPaths int

	stmtTagsIn, args, err := sqlx.In(stmtTags, tags)
	if err != nil {
		return nil, err
	}

	stmtNumPathsIn, args, err := sqlx.In(stmtNumPaths, tags)
	if err != nil {
		return nil, err
	}

	args = append(args, len(tags))

	stmtTagsIn = dbConn.Rebind(stmtTagsIn)
	stmtNumPathsIn = dbConn.Rebind(stmtNumPathsIn)

	stmtTagsIn += " ?;"
	stmtNumPathsIn += " ?;"

	err = dbConn.Get(&numPaths, stmtNumPathsIn, args...)
	if err != nil {
		return nil, err
	}

	paths := make([]string, 0, numPaths)

	err = dbConn.Select(&paths, stmtTagsIn, args...)
	if err != nil {
		return nil, err
	}

	return paths, nil
}

// TODO: Allow passing string and id for document

// FindLinksLines finds the lines in documentPath in which links to other documents are listed.
// It returns the range of line numbers containing links as well as the lines themselves.
func FindLinksLines(documentPath string) (lineNumberFirstLink int, lineNumberLastLink int, linksLines []string, err error) {
	file, err := os.ReadFile(documentPath)
	if err != nil {
		return 0, 0, nil, err
	}

	lines := strings.Split(string(file), "\n")

	lineNumberFirstLink = -1
	lineNumberLastLink = -1
	links := make([]string, 0, 10)

	var line string
	i := 0

	// REFACTOR: This could probably be simplified with goaoi
	for _, line = range lines {
		if line == "# Links" {
			lineNumberFirstLink = i + 1

			break
		}

		i++
	}

	if lineNumberFirstLink == -1 {
		return 0, 0, nil, errors.New("Could not find links")
	}

	for _, line := range lines[lineNumberFirstLink:] {
		if strings.HasPrefix(line, "- ") {
			links = append(links, line)

			i++
		} else if line != "" {
			return 0, 0, nil, errors.New("Invalid links list")
		}
	}

	lineNumberLastLink = i

	return lineNumberFirstLink, lineNumberLastLink, links, nil
}

// TODO: Allow passing string and id for document

// FindBacklinksLines finds the lines in documentPath in which backlinks to other documents are listed.
// It returns the range of line numbers containing backlinks as well as the lines themselves.
func FindBacklinksLines(documentPath string) (lineNumberFirstBacklink int, lineNumberLastBacklink int, backlinksLines []string, err error) {
	file, err := os.ReadFile(documentPath)
	if err != nil {
		return 0, 0, nil, err
	}

	lines := strings.Split(string(file), "\n")

	lineNumberFirstBacklink = -1
	lineNumberLastBacklink = -1
	backlinks := make([]string, 0, 10)

	var line string
	i := 0

	// REFACTOR: This could probably be simplified with goaoi
	for _, line = range lines {
		if line == "# Backlinks" {
			lineNumberFirstBacklink = i + 1

			break
		}
		i++
	}

	if lineNumberFirstBacklink == -1 {
		return 0, 0, nil, errors.New("Could not find links")
	}

	for _, line := range lines[lineNumberFirstBacklink:] {
		if strings.HasPrefix(line, "- ") {
			backlinks = append(backlinks, line)

			i++
		} else if line != "" {
			return 0, 0, nil, errors.New("Invalid backlinks list")
		}
	}

	lineNumberLastBacklink = i

	return lineNumberFirstBacklink, lineNumberLastBacklink, backlinks, nil
}

// TODO: Allow passing string and id for document

// AddLink adds a link to documentPathDestination into the document at documentPathSource.
func AddLink(documentPathSource string, documentPathDestination string) error {
	_, lineNumberlastLink, _, err := FindLinksLines(documentPathSource)
	if err != nil {
		return err
	}

	file, err := os.ReadFile(documentPathSource)
	if err != nil {
		return err
	}

	linesOld := strings.Split(string(file), "\n")

	linesNew := append(linesOld[:lineNumberlastLink], "- ()["+documentPathDestination+"]")
	linesNew = append(linesNew, linesOld[lineNumberlastLink+1:]...)

	err = os.WriteFile(documentPathSource, []byte(strings.Join(linesNew, "\n")), 0o644)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Allow passing string and id for document

// RemoveLink removes the link to documentPathDestination from the document at documentPathSource.
func RemoveLink(documentPathSource string, documentPathDestination string) error {
	lineNumberFirstLink, _, linksOrig, err := FindLinksLines(documentPathSource)
	if err != nil {
		return err
	}

	iLinkToDelete := -1

	// REFACTOR: This could probably be simplified with goaoi
	for i, link := range linksOrig {
		if strings.Contains(link, documentPathDestination) {
			iLinkToDelete = i + lineNumberFirstLink
		}
	}

	file, err := os.ReadFile(documentPathSource)
	if err != nil {
		return err
	}

	linesOld := strings.Split(string(file), "\n")

	linesNew := append(linesOld[:iLinkToDelete-1], linesOld[:iLinkToDelete]...)

	err = os.WriteFile(documentPathSource, []byte(strings.Join(linesNew, "\n")), 0o644)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Allow passing string and id for document

// AddBacklink adds a Backlink to documentPathSource into the document at documentPathDestination.
func AddBacklink(documentPathDestination string, documentPathSource string) error {
	_, lineNumberlastBacklink, _, err := FindBacklinksLines(documentPathDestination)
	if err != nil {
		return err
	}

	file, err := os.ReadFile(documentPathDestination)
	if err != nil {
		return err
	}

	linesOld := strings.Split(string(file), "\n")

	linesNew := append(linesOld[:lineNumberlastBacklink], "- ()["+documentPathSource+"]")
	linesNew = append(linesNew, linesOld[lineNumberlastBacklink+1:]...)

	err = os.WriteFile(documentPathDestination, []byte(strings.Join(linesNew, "\n")), 0o644)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Allow passing string and id for document

// RemoveBacklink removes the backlink to documentPathSource from the document at documentPathDestination.
func RemoveBacklink(documentPathDestination string, documentPathSource string) error {
	lineNumberFirstBacklink, _, backlinksOrig, err := FindBacklinksLines(documentPathDestination)
	if err != nil {
		return err
	}

	iBacklinkToDelete := -1

	// REFACTOR: This could probably be simplified with goaoi
	for i, link := range backlinksOrig {
		if strings.Contains(link, documentPathSource) {
			iBacklinkToDelete = i + lineNumberFirstBacklink
		}
	}

	file, err := os.ReadFile(documentPathDestination)
	if err != nil {
		return err
	}

	linesOld := strings.Split(string(file), "\n")

	linesNew := append(linesOld[:iBacklinkToDelete-1], linesOld[:iBacklinkToDelete]...)

	err = os.WriteFile(documentPathDestination, []byte(strings.Join(linesNew, "\n")), 0o644)
	if err != nil {
		return err
	}

	return nil
}

// TODO: Allow passing string and id for type

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

	_, _, err = helpers.SqlExecute(dbConn, transaction, stmt, documentPath, documentTypeId)

	return err
}

// TODO: Allow passing string and id for document

// RemoveDocument removes a document located at documentPath from the DB.
func RemoveDocument(dbConn *sqlx.DB, transaction *sqlx.Tx, documentPath string) error {
	stmt := `
        DELETE FROM
            Document
        WHERE
            Path = ?;
    `

	_, numAffectedRows, err := helpers.SqlExecute(dbConn, transaction, stmt, documentPath)

	if numAffectedRows == 0 {
		return errors.New("documentPathOld does not exist")
	}

	return err
}

// TODO: Allow passing string and id for old document

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

	_, numAffectedRows, err := helpers.SqlExecute(dbConn, transaction, stmt, documentPathNew, documentPathOld)

	if numAffectedRows == 0 {
		return errors.New("documentPathOld does not exist")
	}

	return err
}

// TODO: Allow passing string and id for document and type

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

	_, numAffectedRows, err := helpers.SqlExecute(dbConn, transaction, stmt, documentTypeId, documentPath)

	if numAffectedRows == 0 {
		return errors.New("DocumentPath does not exist")
	}

	return err
}

// TODO: Allow passing string and id for document and type

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

// TODO: Allow passing string and id for document and type

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
