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

// Package libtags implements functionality to work with tags in a database context.
package libtags

import (
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v3"
)

// ImportYML reads a YML file at ymlPath  and imports it's tag structure into the db.
// The top level tag of the file is expected to be "tags".
func ImportYML(dbConn *sqlx.DB, ymlPath string) error {
	file, err := os.ReadFile(ymlPath)
	if err != nil {
		return err
	}

	// TODO: Simplify and document
	var data interface{}

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	var topLevelTags []interface{}

	switch data.(type) {
	case map[string]interface{}:
		if val, ok := data.(map[string]interface{})["tags"]; ok {
			topLevelTags, ok = val.([]interface{})

			if !ok {
				return errors.New("Top level tag does not have children")
			}
		} else {
			return errors.New("Could not recognize top level tag(should be 'tags'")
		}
	case []interface{}:
		topLevelTags = data.([]interface{})
	default:
		return errors.New("Could not parse YML tag file")
	}

	// REFACTOR: YAML parsing and import should be split
	paths := make(chan []string, 200)

	err = bFSTagPaths(topLevelTags, paths, nil)
	if err != nil {
		return err
	}

	close(paths)

	transaction, err := dbConn.Beginx()
	if err != nil {
		return err
	}

	for path := range paths {
		err = AddTag(dbConn, transaction, strings.Join(path, "::"))
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

func bFSTagPaths(node interface{}, paths chan []string, curPath []string) error {
	switch node.(type) {
	// Iterate over child nodes and BFS them in parallel
	case []interface{}:
		for _, curNode := range node.([]interface{}) {
			func(node interface{}) {
				// Starting path
				if curPath == nil {
					bFSTagPaths(node, paths, make([]string, 0, 10))
					// Continuing path
				} else {
					newPath := make([]string, len(curPath), 10)
					copy(newPath, curPath)

					bFSTagPaths(node, paths, newPath)
				}
			}(curNode)
		}

		// Reached a parent node, add it's component to the current tag path
		// and pass children to BFS
	case map[string]interface{}:
		for key, value := range node.(map[string]interface{}) {
			curPath = append(curPath, key)
			bFSTagPaths(value, paths, curPath)
		}

		// Reached leaf node, add it's component to the current tag path
		// and the current path to the final paths list
	case string:
		curPath = append(curPath, node.(string))
		paths <- curPath
	}

	return nil
}

type tagNode map[string]tagNode

// ExportYML exports the DB's available into a YML encoded hierarchy at ymlPath.
func ExportYML(dbConn *sqlx.DB, ymlPath string) error {
	tags, err := ListTags(dbConn)
	if err != nil {
		return err
	}

	tagHierarchy := tagNode{"tags": tagNode{}}

	for _, tag := range tags {
		curNode := tagHierarchy["tags"]

		tagComponents := strings.Split(tag, "::")
		for _, component := range tagComponents {
			// component is not inserted yet, insert it, so we can add it's children later
			_, ok := curNode[component]
			if !ok {
				curNode[component] = tagNode{}
			}

			// component is added, descend to child
			curNode = curNode[component]
		}
	}

	// TODO: The following code should be extracted to a separate function (split building YML structure and writing it)
	// 0664 UNIX Permission code
	file, err := os.OpenFile(ymlPath, os.O_CREATE|os.O_WRONLY, 0o664)
	if err != nil {
		return err
	}

	defer file.Close()

	yamlFile, err := yaml.Marshal(tagHierarchy)
	if err != nil {
		return err
	}

	fileString := string(yamlFile)
	// TODO: Refactor using bufio and non regex string functions
	// NOTE: This is awful, but I can't seem to get it to work properly any other way
	regexEmptyMap := regexp.MustCompile(`: \{\}`)
	fileString = regexEmptyMap.ReplaceAllString(fileString, "")

	regexListItems := regexp.MustCompile(`(  )(\w)`)
	fileString = regexListItems.ReplaceAllString(fileString, "- $2")

	regexIndentation := regexp.MustCompile(`  `)
	fileString = regexIndentation.ReplaceAllString(fileString, " ")

	regexIndentation2 := regexp.MustCompile(`( )(-)`)
	fileString = regexIndentation2.ReplaceAllString(fileString, "$2")

	_, err = file.Write([]byte(fileString))
	if err != nil {
		return err
	}

	return nil
}

// TODO: Allow passing string and id for tag

// AddTag adds a new tag to the DB.
// Passing a transaction is optional.
func AddTag(dbConn *sqlx.DB, transaction *sqlx.Tx, tag string) error {
	stmt := `
        INSERT INTO
            Tag(Tag)
        VALUES(?);
    `

	_, _, err := helpers.SqlExecute(dbConn, transaction, stmt, tag)

	return err
}

// TODO: Allow passing string and id for tag

// RenameTag renames the tag oldTag to newTag in the DB.
// Passing a transaction is optional.
func RenameTag(dbConn *sqlx.DB, transaction *sqlx.Tx, oldTag string, newTag string) error {
	stmt := `
        UPDATE
            Tag
        SET
            Tag = ?
        WHERE
            Tag = ?;
    `

	_, _, err := helpers.SqlExecute(dbConn, transaction, stmt, oldTag, newTag)

	return err
}

// TODO: Allow passing string and id for tag

// DeleteTag removes the tag tag from the DB.
// Passing a transaction is optional.
func DeleteTag(dbConn *sqlx.DB, transaction *sqlx.Tx, tag string) error {
	stmt := `
        DELETE FROM
            Tag
        WHERE
            Tag = ?;
    `

	_, _, err := helpers.SqlExecute(dbConn, transaction, stmt, tag)

	return err
}

// TODO: Allow passing string and id for tag

// FindAmbiguousTagComponent finds the index (root = 0) of an ambiguous component.
func FindAmbiguousTagComponent(dbConn *sqlx.DB, tag string) (ambiguousIndex int, ambiguousComponent string, err error) {
	stmt := `
        SELECT
            Tag
        FROM
            Tag
        WHERE
            INSTR(Tag, ?) > 0
            AND
            Tag != ?;
    `
	inputTagComponents := strings.Split(tag, "::")
	inputLeaf := inputTagComponents[len(inputTagComponents)-1]

	var ambiguousTag string

	statement, err := dbConn.Preparex(stmt)
	if err != nil {
		return -1, "", err
	}

	defer statement.Close()

	_ = statement.Get(&ambiguousTag, inputLeaf, tag)

	ambiguousTagComponents := strings.Split(ambiguousTag, "::")

	i := len(ambiguousTagComponents) - 1

	// REFACTOR: This could probably be simplified with goaoi
	// Find where input tag's leaf appears in ambiguous tag
	for ; i > 0; i-- {
		if ambiguousTagComponents[i] == inputLeaf {
			break
		}
	}

	// Find index of first differing tag component (traversal from leaf to root)
	ambiguousIndex = len(inputTagComponents) - 1

	for i > 0 && ambiguousIndex > 0 && ambiguousTagComponents[i] == inputTagComponents[ambiguousIndex] {
		i--
		ambiguousIndex--
	}

	ambiguousComponent = inputTagComponents[ambiguousIndex]

	return
}

// TODO: Allow passing string and id for tag

// TryShortenTag shortens tag as much as possible, while keeping it unambiguous.
// Components are removed from root to leaf.
// A::B::C can be shortened to C if C does not appear in any other tag(e.g. X::C::Y).
func TryShortenTag(dbConn *sqlx.DB, tag string) (string, error) {
	tagComponents := strings.Split(tag, "::")

	isAmbiguous, err := IsLeafAmbiguous(dbConn, tag)
	if err != nil {
		return "", err
	}

	if !isAmbiguous {
		return tagComponents[len(tagComponents)-1], nil
	}

	i, _, err := FindAmbiguousTagComponent(dbConn, tag)
	if err != nil {
		return "", err
	}

	return strings.Join(tagComponents[i:], "::"), nil
}

// TODO: Allow passing string and id for tag

// IsLeafAmbiguous checks if the leaf of the specified tag appears in any other tag.
func IsLeafAmbiguous(dbConn *sqlx.DB, tag string) (bool, error) {
	tags, err := ListTags(dbConn)
	if err != nil {
		return false, err
	}

	tagComponents := strings.Split(tag, "::")
	leaf := tagComponents[len(tagComponents)-1]
	seenOnce := false

	for _, tag := range tags {
		curTagComponents := strings.Split(tag, "::")
		curLeaf := curTagComponents[len(curTagComponents)-1]

		if curLeaf == leaf && seenOnce {
			return true, nil
		}

		seenOnce = true
	}

	return false, nil
}

// ListTags lists all available from the DB.
// Tags are listed fully qualified, no components are removed.
func ListTags(dbConn *sqlx.DB) ([]string, error) {
	stmt := `
        SELECT
            Tag
        FROM
            Tag;
    `

	tagsBuffer := []string{}

	err := dbConn.Select(&tagsBuffer, stmt)
	if err != nil {
		return nil, err
	}

	return tagsBuffer, nil
}

// ListTagsShortened lists all available from the DB.
// Tags are shortened as much as possible while being kept unambiguous.
func ListTagsShortened(dbConn *sqlx.DB) ([]string, error) {
	tags, err := ListTags(dbConn)
	if err != nil {
		return nil, err
	}

	for i, tag := range tags {
		shortenedTag, err := TryShortenTag(dbConn, tag)
		if err != nil {
			return nil, err
		}

		tags[i] = shortenedTag
	}

	return tags, nil
}
