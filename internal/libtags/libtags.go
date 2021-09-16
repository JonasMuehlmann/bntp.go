package libtags

import (
	"database/sql"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

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
				//func(node interface{}) {
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

func ExportYML(dbConn *sql.DB, ymlPath string) {
	// 0664 UNIX Permission code
	file, err := os.OpenFile(ymlPath, os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
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

func RenameTag(dbConn *sql.DB, transaction *sql.Tx, oldTag string, newTag string) {
}

func DeleteTag(dbConn *sql.DB, transaction *sql.Tx, tag string) {
}

func ListChildTags(dbConn *sql.DB, tag string) {
}

func TryShortenTag(dbConn *sql.DB, tag string) string {
}

func IsLeafAmbiguous(dbConn *sql.DB, tag string) bool {
}

func ListTags(dbConn *sql.DB) [][]string {
}

func ListTagsShortened(dbConn *sql.DB) [][]string {
}
