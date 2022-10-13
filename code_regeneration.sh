#!/usr/bin/env sh

rm model/repository/**/*_repository.go || true
rm model/repository/**/*_repository_test.go || true

go run ./tools/schema_converter
go run ./tools/generate_external_cli_documentation
go run ./tools/generate_domain_models
go run ./tools/generate_repository_interfaces
go run ./tools/generate_sql_repositories

cp templates/*_test.go model/repository/sqlite3/
# cp templates/*_test.go model/repository/mssql/
# cp templates/*_test.go model/repository/mysql/
# cp templates/*_test.go model/repository/psql/

# sed -i "s/sqlite3/psql/g" model/repository/psql/*_test.go
# sed -i "s/sqlite3/mssql/g" model/repository/mssql/*_test.go
# sed -i "s/sqlite3/mysql/g" model/repository/mysql/*_test.go

if ! grep "maybe.*Generated" model/repository/**/*.go >/dev/null; then
	sed -ri "s/(\S*GeneratedColumns\s*=).*/\1[]string{}\nmaybe\1 []string{\"id\"}/g" model/repository/**/*.go
fi

git restore model/repository/document_content_repository.go model/repository/fs/document_content_repository.go
