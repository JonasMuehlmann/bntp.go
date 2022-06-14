#!/usr/bin/env sh

rm model/repository/**/*_converter.go || true
rm model/repository/**/*_updater.go || true
rm model/repository/**/*_repository.go || true
go run ./tools/generate_sql_repositories
go run ./tools/generate_sql_repository_model_converters/main.go
go run ./tools/generate_sql_repository_updater_converters/main.go
go run ./tools/generate_sql_repository_filter_converters/main.go
go run ./tools/generate_repository_interfaces && git restore model/repository/document_content_repository.go model/repository/fs/document_content_repository.go
