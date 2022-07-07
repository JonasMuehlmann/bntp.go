#!/usr/bin/env sh

rm model/repository/**/*_repository.go || true

go run ./tools/schema_converter
go run ./tools/generate_external_cli_documentation
go run ./tools/generate_domain_models
go run ./tools/generate_repository_interfaces
go run ./tools/generate_sql_repositories
go run ./tools/generate_repository_tests

git restore model/repository/document_content_repository.go model/repository/fs/document_content_repository.go
