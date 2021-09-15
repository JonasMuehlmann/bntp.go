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

	libbookmarks.AddFromCSV(dbConn, filepath.Join(user.HomeDir, "scripts/deepstash_ideas.csv"))
}
