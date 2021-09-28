// Package libtags implements functionality to work with tags in a database context.
package libtags

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/JonasMuehlmann/bntp.go/internal/sqlhelpers"
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

	var data interface{}

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return err
	}

	var topLevelTags []interface{}

	switch data.(type) {
	case map[string]interface{}:
		if val, ok := data.(map[string]interface{})["tags"]; ok {
			topLevelTags = val.([]interface{})
		} else {
			log.Fatal("Could not recognize top level tag(should be 'tags'")
		}
	case []interface{}:
		topLevelTags = data.([]interface{})
	default:
		log.Fatal("Could not parse YML tag file.")
	}

	paths := make(chan []string, 200)

	bFSTagPaths(topLevelTags, paths, nil)

	close(paths)

	transaction, err := dbConn.Beginx()
	if err != nil {
		return err
	}

	for path := range paths {
		AddTag(dbConn, transaction, strings.Join(path, "::"))
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
				// func(node interface{}) {
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
		tagComponents := strings.Split(tag, "::")

		curNode := tagHierarchy["tags"]

		for _, component := range tagComponents {
			_, ok := curNode[component]

			if !ok {
				curNode[component] = tagNode{}
			}

			curNode = curNode[component]
		}
	}

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

// AddTag adds a new tag to the DB.
// Passing a transaction is optional.
func AddTag(dbConn *sqlx.DB, transaction *sqlx.Tx, tag string) error {
	stmt := `
        INSERT INTO
            Tag(Tag)
        VALUES(?);
    `

	return sqlhelpers.Execute(dbConn, transaction, stmt, tag)
}

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

	return sqlhelpers.Execute(dbConn, transaction, stmt, oldTag, newTag)
}

// DeleteTag removes the tag tag from the DB.
// Passing a transaction is optional.
func DeleteTag(dbConn *sqlx.DB, transaction *sqlx.Tx, tag string) error {
	stmt := `
        DELETE FROM
            Tag
        WHERE
            Tag = ?;
    `

	return sqlhelpers.Execute(dbConn, transaction, stmt, tag)
}

// FindAmbiguousTagComponent finds the index (root = 0) of an ambiguous component.
func FindAmbiguousTagComponent(dbConn *sqlx.DB, tag string) (int, error) {
	stmt := `
        SELECT
            Tag
        FROM
            Tag
        WHERE
            INSTR(Tag, ?) > 0;
    `
	inputTagComponents := strings.Split(tag, "::")
	leaf := inputTagComponents[len(inputTagComponents)-1]

	var tagWithAmbiguousComponent string

	statement, err := dbConn.Preparex(stmt)
	if err != nil {
		return -1, err
	}

	defer statement.Close()

	err = statement.Get(&tagWithAmbiguousComponent, leaf)
	if err != nil {
		return -1, err
	}

	i := strings.Index(leaf, tagWithAmbiguousComponent)
	if i == -1 {
		return -1, err
	}

	return i, nil
}

// TryShortenTag shortens tag as much as possible, while keeping it unambiguous.
// Components are removed from root to leaf.
// A::B::C can be shortened to C if C does not appear in any other tag(e.g. X::C::Y).
func TryShortenTag(dbConn *sqlx.DB, tag string) (string, error) {
	isAmbiguous, err := IsLeafAmbiguous(dbConn, tag)
	if err != nil {
		return "", err
	}

	// FIX: This should find the ambiguous component
	// and return a tag containing itself and it's children
	if isAmbiguous {
		i, err := FindAmbiguousTagComponent(dbConn, tag)
		if err != nil {
			return "", err
		}
		tagComponents := strings.Split(tag, "::")

		return strings.Join(tagComponents[i:], "::"), nil
	}

	return tag, nil
}

// IsLeafAmbiguous checks if the leaf of the specified tag appears in any other tag.
func IsLeafAmbiguous(dbConn *sqlx.DB, tag string) (bool, error) {
	tags, err := ListTags(dbConn)
	if err != nil {
		return false, err
	}
	tagComponents := strings.Split(tag, "::")

	leaf := tagComponents[len(tagComponents)-1]

	for _, tag := range tags {
		curTagComponents := strings.Split(tag, "::")
		curLeaf := curTagComponents[len(curTagComponents)-1]

		if curLeaf == leaf {
			return true, nil
		}
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
