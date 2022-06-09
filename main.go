package main

//go:generate go run ./tools/schema_converter
//go:generate go run ./tools/generate_config_key_constants
//go:generate go run ./tools/generate_external_cli_documentation
//go:generate go run ./tools/generate_domain_models
//go:generate go run ./tools/generate_repository_interfaces
//go:generate go run ./tools/generate_sql_repositories
//go:generate go run ./tools/generate_sql_repository_filter_converters
//go:generate go run ./tools/generate_sql_repository_model_converters
//go:generate go run ./tools/generate_sql_repository_updater_converters

import (
	"database/sql"

	"github.com/JonasMuehlmann/bntp.go/bntp"
	"github.com/JonasMuehlmann/bntp.go/bntp/backend"
	"github.com/JonasMuehlmann/bntp.go/bntp/libbookmarks"
	"github.com/JonasMuehlmann/bntp.go/bntp/libdocuments"
	"github.com/JonasMuehlmann/bntp.go/bntp/libtags"
	"github.com/JonasMuehlmann/bntp.go/cmd"
	"github.com/JonasMuehlmann/bntp.go/model/domain"
	"github.com/spf13/afero"

	fsRepository "github.com/JonasMuehlmann/bntp.go/model/repository/fs"
	// mssqlRepository "github.com/JonasMuehlmann/bntp.go/model/repository/mssql"
	// mysqlRepository "github.com/JonasMuehlmann/bntp.go/model/repository/mysql"
	// psqlRepository "github.com/JonasMuehlmann/bntp.go/model/repository/psql".
	sqlite3Repository "github.com/JonasMuehlmann/bntp.go/model/repository/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	backend := new(backend.Backend)

	db, err := sql.Open("sqlite3", "bntp_tests.db")
	if err != nil {
		panic(err)
	}

	fs := afero.NewOsFs()

	// ***********************    Set up hooks    ***********************//

	bookmarkHooks := new(bntp.Hooks[domain.Bookmark])
	tagsHooks := new(bntp.Hooks[domain.Tag])
	documentHooks := new(bntp.Hooks[domain.Document])
	documentContentHooks := new(bntp.Hooks[string])

	// ********************    Set up repositories    *******************//

	bookmarkRepository := new(sqlite3Repository.Sqlite3BookmarkRepository)

	bookmarkRepositoryAbstract, err := bookmarkRepository.New(sqlite3Repository.Sqlite3BookmarkRepositoryConstructorArgs{DB: db})
	if err != nil {
		panic(err)
	}

	bookmarkRepository, _ = bookmarkRepositoryAbstract.(*sqlite3Repository.Sqlite3BookmarkRepository)

	tagsRepository := new(sqlite3Repository.Sqlite3TagRepository)

	tagsRepositoryAbstract, err := tagsRepository.New(sqlite3Repository.Sqlite3TagRepositoryConstructorArgs{DB: db})
	if err != nil {
		panic(err)
	}

	tagsRepository, _ = tagsRepositoryAbstract.(*sqlite3Repository.Sqlite3TagRepository)

	documentRepository := new(sqlite3Repository.Sqlite3DocumentRepository)

	documentRepositoryAbstract, err := documentRepository.New(sqlite3Repository.Sqlite3DocumentRepositoryConstructorArgs{DB: db})
	if err != nil {
		panic(err)
	}

	documentRepository, _ = documentRepositoryAbstract.(*sqlite3Repository.Sqlite3DocumentRepository)

	documentContentRepository := new(fsRepository.FSDocumentContentRepository)

	documentContentRepositoryAbstract, err := documentContentRepository.New(fs)
	if err != nil {
		panic(err)
	}

	documentContentRepository, _ = documentContentRepositoryAbstract.(*fsRepository.FSDocumentContentRepository)

	// **********************    Set up Managers    *********************//

	bookmarkManager, err := libbookmarks.NewBookmarkManager(bookmarkHooks, bookmarkRepository)
	if err != nil {
		panic(err)
	}

	tagsManager, err := libtags.NewTagmanager(tagsHooks, tagsRepository)
	if err != nil {
		panic(err)
	}

	documentManager, err := libdocuments.NewDocumentManager(documentHooks, documentRepository)
	if err != nil {
		panic(err)
	}

	documentContentManager, err := libdocuments.NewDocumentContentRepository(documentContentHooks, documentContentRepository)
	if err != nil {
		panic(err)
	}

	// **********************    Set up backend    **********************//

	backend.BookmarkManager = bookmarkManager
	backend.TagManager = tagsManager
	backend.DocumentManager = documentManager
	backend.DocumentContentManager = documentContentManager

	cmd.Execute(backend)
}
