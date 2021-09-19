package main

import (
	"database/sql"
	"log"
	"os/user"
	"path/filepath"

	"github.com/JonasMuehlmann/productivity.go/internal/libbookmarks"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	user, err := user.Current()
	dbConn, err := sql.Open("sqlite3", filepath.Join(user.HomeDir, "Documents/productivity/bookmarks.db"))

	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	// libbookmarks.ImportMinimalCSV(dbConn, filepath.Join(user.HomeDir, "scripts/deepstash_ideas.csv.uniq"))
	// libbookmarks.ExportCSV(dbConn, filepath.Join(user.HomeDir, "Documents/productivity/out.csv"), nil)
	libbookmarks.ExportCSV(dbConn, "", map[string]interface{}{"Tags": []string{"software_development::hardware", "compiler_dev"}})
}
