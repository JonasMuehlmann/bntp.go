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
	"regexp"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/helpers"
	"github.com/JonasMuehlmann/goaoi"
	"github.com/jmoiron/sqlx"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

type TagNode map[string]TagNode

func bFSTagPaths(node any, paths *[][]string, curPath *[]string) error {
	switch node.(type) {
	// Iterate over child nodes and BFS them
	case []any:
		for _, curNode := range node.([]any) {
			func(node any) {
				// Starting path
				if curPath == nil {
					newPath := make([]string, 0, 10)
					bFSTagPaths(node, paths, &newPath)
					// Continuing path
				} else {
					bFSTagPaths(node, paths, curPath)
				}
			}(curNode)
		}

		// Reached a parent node, add it's component to the current tag path
		// and pass children to BFS
	case map[string]any:
		for key, value := range node.(map[string]any) {
			*curPath = append(*curPath, key)
			bFSTagPaths(value, paths, curPath)
		}

		// Reached leaf node, add it's component to the current tag path
		// and the current path to the final paths list
	case string:
		*curPath = append(*curPath, node.(string))
		*paths = append(*paths, *curPath)
	}

	return nil
}

func buildInternalTagHierarchy(tags []string) (tagHierarchy TagNode, err error) {
	tagHierarchy = TagNode{"tags": TagNode{}}

	for _, tag := range tags {
		curNode := tagHierarchy["tags"]

		tagComponents := strings.Split(tag, "::")
		for _, component := range tagComponents {
			// component is not inserted yet, insert it, so we can add it's children later
			_, ok := curNode[component]
			if !ok {
				curNode[component] = TagNode{}
			}

			// component is added, process children in next iteration
			curNode = curNode[component]
		}
	}

	return
}

func DeserializeTagHierarchy(serializedTagHierarchy string) (tagHierarchy TagNode, err error) {
	var data map[string]any

	err = yaml.Unmarshal([]byte(serializedTagHierarchy), &data)
	if err != nil {
		err = helpers.DeserializationError{Inner: err}

		return
	}

	if _, ok := data["tags"]; !ok {
		err = helpers.DeserializationError{Inner: errors.New(`Top level tag should be "tags"`)}

		return
	}

	paths := make([][]string, 0, 200)

	// NOTE: This is terribly bad code but the unmarshalling logic is a PITA and I can't be bothered to fix this
	err = bFSTagPaths(data["tags"], &paths, nil)
	if err != nil {
		err = helpers.DeserializationError{Inner: err}

		return
	}

	if len(paths) == 0 {
		err = helpers.IneffectiveOperationError{Inner: errors.New("Empty tag hierarchy")}

		return
	}

	tags := make([]string, 0, len(paths))
	goaoi.ForeachSliceUnsafe(paths, func(path []string) { tags = append(tags, strings.Join(path, "::")) })

	tagHierarchy, err = buildInternalTagHierarchy(tags)

	return
}

func ImportTagHierarchy(dbConn *sqlx.DB, tagHierarchy TagNode) error {
	paths := make([][]string, 0, 200)

	err := bFSTagPaths(tagHierarchy["tags"], &paths, nil)
	if err != nil {
		if errors.As(err, &helpers.IneffectiveOperationError{}) {
			return helpers.IneffectiveOperationError{Inner: err}
		}

		return helpers.ImportError{Inner: err}
	}

	transaction, err := dbConn.Beginx()
	if err != nil {
		return helpers.ImportError{Inner: err}
	}

	for _, path := range paths {
		err = AddTag(dbConn, transaction, strings.Join(path, "::"))
		if err != nil {
			transaction.Rollback()

			return helpers.ImportError{Inner: err}
		}
	}

	err = transaction.Commit()
	if err != nil {
		return helpers.ImportError{Inner: err}
	}

	return nil
}

func SerializeTagHierarchy(tagHierarchy TagNode) (serializedTagHierarchy string, err error) {
	rootTag, okRootTag := tagHierarchy["tags"]
	if !okRootTag {
		err = helpers.SerializationError{Inner: errors.New("Trying to serialize tag hierarchy without root tag")}

		return
	}

	topLevelTags := maps.Keys(rootTag)
	if len(topLevelTags) == 0 {
		err = helpers.IneffectiveOperationError{Inner: errors.New("Tag hierarchy has no top level tags")}

		return
	}

	rawYAML, err := yaml.Marshal(tagHierarchy)
	if err != nil {
		err = helpers.SerializationError{Inner: err}

		return
	}

	serializedTagHierarchy = string(rawYAML)

	// TODO: Refactor using bufio and non regex string functions
	// NOTE: This is awful, but I can't seem to get it to work properly any other way
	regexEmptyMap := regexp.MustCompile(`: \{\}`)
	serializedTagHierarchy = regexEmptyMap.ReplaceAllString(serializedTagHierarchy, "")

	regexListItems := regexp.MustCompile(`(  )(\w)`)
	serializedTagHierarchy = regexListItems.ReplaceAllString(serializedTagHierarchy, "- $2")

	regexIndentation := regexp.MustCompile(`  `)
	serializedTagHierarchy = regexIndentation.ReplaceAllString(serializedTagHierarchy, " ")

	regexIndentation2 := regexp.MustCompile(`( )(-)`)
	serializedTagHierarchy = regexIndentation2.ReplaceAllString(serializedTagHierarchy, "$2")

	return
}

func ExportTagHierarchy(dbConn *sqlx.DB) (tagHierarchy TagNode, err error) {
	tags, err := ListTags(dbConn)
	if err != nil {
		err = helpers.ExportError{Inner: err}

		return
	}

	tagHierarchy, err = buildInternalTagHierarchy(tags)

	return
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

	tagsBuffer := make([]string, 0, 50)

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
