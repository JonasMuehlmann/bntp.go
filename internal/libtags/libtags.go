// libtags implements functionality to work with tags in a database context.
package libtags

import (
	"database/sql"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/yaml.v3"
)

// ImportYML reads a YML file at ymlPath  and imports it's tag structure into the db.
// The top level tag of the file is expected to be "tags".
func ImportYML(dbConn *sql.DB, ymlPath string) {
	file, err := os.ReadFile(ymlPath)
	if err != nil {
		log.Fatal(err)
	}

	var data interface{}

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
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

	transaction, err := dbConn.Begin()
	if err != nil {
		log.Fatal(err)
	}

	for path := range paths {
		AddTag(dbConn, transaction, strings.Join(path, "::"))
	}

	err = transaction.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func bFSTagPaths(node interface{}, paths chan []string, curPath []string) {
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
}

type tagNode map[string]tagNode

// ExportYML exports the DB's available into a YML encoded hierarchy at ymlPath.
func ExportYML(dbConn *sql.DB, ymlPath string) {
	tags := ListTags(dbConn)

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
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	yamlFile, err := yaml.Marshal(tagHierarchy)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}
}

// AddTag adds a new tag to the DB.
// Passing a transaction is optional.
func AddTag(dbConn *sql.DB, transaction *sql.Tx, tag string) {
	stmt := `
        INSERT INTO
            Tag(Tag)
        VALUES(?);
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

	_, err = statement.Exec(tag)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
}

// RenameTag renames the tag oldTag to newTag in the DB.
// Passing a transaction is optional.
func RenameTag(dbConn *sql.DB, transaction *sql.Tx, oldTag string, newTag string) {
	stmt := `
        UPDATE
            Tag
        SET
            Tag = '?'
        WHERE
            Tag = '?';
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

	_, err = statement.Exec(newTag, oldTag)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
}

// DeleteTag removes the tag tag from the DB.
// Passing a transaction is optional.
func DeleteTag(dbConn *sql.DB, transaction *sql.Tx, tag string) {
	stmt := `
        DELETE FROM
            Tag
        WHERE
            Tag = '?';
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

	_, err = statement.Exec(tag)
	if err != nil {
		log.Fatal(err)
	}

	err = statement.Close()

	if err != nil {
		log.Fatal(err)
	}
}

// TryShortenTag shortens tag as much as possible, while keeping it unambiguous.
// Components are removed from root to leaf.
// A::B::C can be shortened to C if C does not appear in any other tag(e.g. X::C::Y).
func TryShortenTag(dbConn *sql.DB, tag string) string {
	if IsLeafAmbiguous(dbConn, tag) {
		tagComponents := strings.Split(tag, "::")

		return tagComponents[len(tagComponents)-1]
	}

	return tag
}

// IsLeafAmbiguous checks if the leaf of the specified tag appears in any other tag.
func IsLeafAmbiguous(dbConn *sql.DB, tag string) bool {
	tags := ListTags(dbConn)

	tagComponents := strings.Split(tag, "::")

	leaf := tagComponents[len(tagComponents)-1]

	for _, tag := range tags {
		curTagComponents := strings.Split(tag, "::")
		curLeaf := curTagComponents[len(curTagComponents)-1]

		if curLeaf == leaf {
			return true
		}
	}

	return false
}

// ListTags lists all available from the DB.
// Tags are listed fully qualified, no components are removed.
func ListTags(dbConn *sql.DB) []string {
	stmt := `
        SELECT
            Tag
        FROM
            Tag;
    `

	tagRows, err := dbConn.Query(stmt)
	if err != nil {
		log.Fatal(err)
	}

	stmtCountTags := "SELECT COUNT(*) FROM  Tag;"

	tagsCountRow := dbConn.QueryRow(stmtCountTags)

	var rowCountTags int

	err = tagsCountRow.Scan(&rowCountTags)
	if err != nil {
		log.Fatal(err)
	}

	tagsBuffer := make([]string, rowCountTags)

	i := 0
	for tagRows.Next() {
		err := tagRows.Scan(&tagsBuffer[i])
		if err != nil {
			log.Fatal(err)
		}
		i++
	}

	return tagsBuffer
}

// ListTagsShortened lists all available from the DB.
// Tags are shortened as much as possible while being kept unambiguous.
func ListTagsShortened(dbConn *sql.DB) []string {
	tags := ListTags(dbConn)

	for i, tag := range tags {
		if IsLeafAmbiguous(dbConn, tag) {
			tagComponents := strings.Split(tag, "::")
			tags[i] = tagComponents[len(tagComponents)-1]
		}
	}

	return tags
}
